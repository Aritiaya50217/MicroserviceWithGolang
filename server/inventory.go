package server

import (
	inventoryGrpcHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/inventory/inventoryHandler"
	inventoryHttpHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/inventory/inventoryHandler"
	inventoryRepository "github.com/Aritiaya50217/MicroserviceWithGolang/modules/inventory/inventoryRepository"
	inventoryUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHttpHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryGrpcHandler.NewInventoryGrpcHttpHandler(usecase)
	queueHandler := inventoryHttpHandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	// set api group
	inventory := s.app.Group("/inventory_v1")
	inventory.GET("/", s.healthCheckService)
}
