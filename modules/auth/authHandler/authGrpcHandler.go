package auth

import (
	"context"

	authPb "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authPb"
	authUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{authUsecase: authUsecase}
}

func (g *authGrpcHandler) CredentialSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	// return g.authUsecase.AccessTokenSearch(ctx, req.AccessToken)
	return nil, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	// return g.authUsecase.RolesCount(ctx)
	return nil, nil
}
