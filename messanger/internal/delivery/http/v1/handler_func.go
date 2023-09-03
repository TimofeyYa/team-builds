package v1

import (
	"context"
	"teamBuild/messages/internal/models"
	httpParcer "teamBuilds/libs/http_parcer"
)

type loginResponse struct {
	Token  string `json:"token"`
	Status bool   `json:"status"`
}

func (h *Handler) Login(c context.Context, cred models.Credentials) (*loginResponse, *httpParcer.ErrorHTTP) {

	return nil, nil
}

type registrationResponse struct {
	Data   models.User
	Status bool `json:"status"`
}

func (h *Handler) Registration(c context.Context, regData models.RegistrationUser) (*registrationResponse, *httpParcer.ErrorHTTP) {
	userData, err := h.service.CreateUser(regData)
	return &registrationResponse{
		Data:   *userData,
		Status: true,
	}, err
}

func (h *Handler) Authorization() {

}
