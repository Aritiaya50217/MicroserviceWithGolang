package server

import (
	playerGrpcHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerHandler"
	playerHttpHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerHandler"
	playerQueueHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerHandler"
	playerRepository "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerRepository"
	playerUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHttpHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerGrpcHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerQueueHandler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = grpcHandler
	_ = queueHandler

	// set api group
	player := s.app.Group("/player_v1")
	player.GET("/", s.healthCheckService)
	player.POST("/palyer/register", httpHandler.CreatePlayer)
	player.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
	player.POST("/player/add-money", httpHandler.AddPlayerMoney)
	player.GET("/player/account/:player_id", httpHandler.GetPlayerSavingAccount)
}
