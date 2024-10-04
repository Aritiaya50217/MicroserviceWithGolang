package server

import (
	"log"

	itemGrpcHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemHandler"
	itemHttpHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemHandler"
	itemPb "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemPb"
	itemRepository "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemRepository"
	itemUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemUsecase"
	grpccon "github.com/Aritiaya50217/MicroserviceWithGolang/pkg/database/grpcCon"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHttpHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemGrpcHandler.NewItemGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)
		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listening on %s ", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	// set api group
	item := s.app.Group("/item_v1")
	item.GET("/", s.healthCheckService)

}
