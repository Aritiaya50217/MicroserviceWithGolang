package server

import (
	"log"

	authGrpcHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authHandler"
	authHttpHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authHandler"
	authPb "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authPb"
	authRepository "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authRepository"
	authUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/auth/authUsecase"
	grpccon "github.com/Aritiaya50217/MicroserviceWithGolang/pkg/database/grpcCon"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHttpHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authGrpcHandler.NewAuthGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, nil)

		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()
	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("/", s.healthCheckService)

	auth.GET("/test/:player_id", s.healthCheckService)

}
