package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	"github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth"
	authUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authUsecase"
	"github.com/Aritiaya50217/MicroserviceWithGolang/pkg/request"
	"github.com/Aritiaya50217/MicroserviceWithGolang/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	AuthHttpHandlerService interface {
		Login(c echo.Context) error
		RefreshToken(c echo.Context) error
		Logout(c echo.Context) error
	}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}

func (h *authHttpHandler) Login(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(auth.PlayerLoginReq)
	fmt.Println("e : \np: \n", req.Email, req.Password)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.Login(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *authHttpHandler) RefreshToken(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)
	req := new(auth.RefreshTokenReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.RefreshToken(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *authHttpHandler) Logout(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(auth.LogoutReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	res, err := h.authUsecase.Logout(ctx, req.CredentialId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, &response.MsgResponse{
		Message: fmt.Sprintf("Deleted count : %d ", res),
	})
}
