package v1

import (
	"context"
	"fmt"
	"teamBuild/messages/internal/models"
	httpParcer "teamBuilds/libs/http_parcer"
)

type loginResponse struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

func (h *Handler) Login(c context.Context, cred models.Credentials) (*loginResponse, *httpParcer.ErrorHTTP) {
	fmt.Println(cred.Login)
	return nil, nil
}

func (h *Handler) Registration() {

}

func (h *Handler) Authorization() {

}
