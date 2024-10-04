package server

import (
	paymentHttpHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/payment/paymentHandler"
	paymentQueueHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/payment/paymentHandler"
	paymentRepository "github.com/Aritiaya50217/MicroserviceWithGolang/modules/payment/paymentRepository"
	paymentUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHttpHandler.NewPaymentHttpHandler(s.cfg, usecase)
	grpcHandler := paymentQueueHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler

	// set api group
	payment := s.app.Group("/payment_v1")
	payment.GET("/", s.healthCheckService)

}
