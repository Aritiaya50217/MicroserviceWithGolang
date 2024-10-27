package playerhandler

import (
	"context"

	playerPb "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerPb"
	playerUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerUsecase playerUsecase.PlayerUsecaseService
		playerPb.UnimplementedPlayerGrpcServiceServer
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{
		playerUsecase: playerUsecase,
	}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return g.playerUsecase.FindOnePlayerCredential(ctx, req.Password, req.Email)
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return g.playerUsecase.FindOnePlayerProfileToRefresh(ctx, req.PlayerId)

}

func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}
