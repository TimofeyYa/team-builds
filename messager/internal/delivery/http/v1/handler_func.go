package v1

import (
	"errors"
	"strings"
	"teamBuild/messages/internal/models"
	httpParcer "teamBuilds/libs/http_parcer"

	"github.com/gin-gonic/gin"
)

type loginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Status       bool   `json:"status"`
}

func (h *Handler) Login(c *gin.Context, cred models.Credentials) (*loginResponse, *httpParcer.ErrorHTTP) {
	token, refreshToken, err := h.service.LoginUser(c, cred)
	if err != nil {
		return nil, err
	}

	c.SetCookie("refresh_token", refreshToken, 3600, "/", "localhost", false, true)

	return &loginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		Status:       true,
	}, nil
}

type registrationResponse struct {
	User   models.User `json:"user"`
	Status bool        `json:"status"`
}

func (h *Handler) Registration(c *gin.Context, regData models.RegistrationUser) (*registrationResponse, *httpParcer.ErrorHTTP) {
	userData, err := h.service.CreateUser(c, regData)
	if err != nil {
		return nil, err
	}

	return &registrationResponse{
		User:   *userData,
		Status: true,
	}, err
}

// Интерфейс обращается на роут когда токен, который он содержит, истекает по времени
// Тогда он делает запрос на обновление токена
func (h *Handler) Authorization(c *gin.Context) {
	tokens := models.TokenPair{}
	jwt, err := h.getJWTFromHeader(c)
	if err != nil {
		h.errorResponse(c, 401, err.Error())
	}
	tokens.JWT = jwt
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		h.errorResponse(c, 401, err.Error())
		return
	}
	tokens.RefreshToken = refreshToken

	newTokens, httpErr := h.service.Authorization(c, &tokens)
	if httpErr != nil {
		h.errorResponse(c, httpErr.Code, httpErr.Msg)
		return
	}

	c.SetCookie("refresh_token", newTokens.RefreshToken, 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"status": true,
		"token":  newTokens.JWT,
	})
}

func (h *Handler) validateJWT(c *gin.Context) {
	jwt, err := h.getJWTFromHeader(c)
	if err != nil {
		h.errorResponse(c, 401, err.Error())
		return
	}

	httpErr := h.service.ValidateToken(c, jwt)
	if err != nil {
		h.errorResponse(c, httpErr.Code, httpErr.Msg)
		return
	}
}

func (h *Handler) getJWTFromHeader(c *gin.Context) (string, error) {
	tokenHeader := c.GetHeader("Authorization")
	if len(tokenHeader) == 0 {
		return "", errors.New("empty Authorization header")
	}
	tokenParts := strings.Split(tokenHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header")
	}

	return tokenParts[1], nil
}
