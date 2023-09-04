package v1

import (
	"context"
	"teamBuild/messages/internal/models"
	httpParcer "teamBuilds/libs/http_parcer"
)

type loginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Status       bool   `json:"status"`
}

func (h *Handler) Login(c context.Context, cred models.Credentials) (*loginResponse, *httpParcer.ErrorHTTP) {
	token, refreshToken, err := h.service.LoginUser(c, cred)
	if err != nil {
		return nil, err
	}
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

func (h *Handler) Registration(c context.Context, regData models.RegistrationUser) (*registrationResponse, *httpParcer.ErrorHTTP) {
	userData, err := h.service.CreateUser(c, regData)
	if err != nil {
		return nil, err
	}

	return &registrationResponse{
		User:   *userData,
		Status: true,
	}, err
}

func (h *Handler) Authorization() {

}
