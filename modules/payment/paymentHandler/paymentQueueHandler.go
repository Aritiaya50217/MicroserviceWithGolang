package paymenthandler

import (
	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	paymentUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/payment/paymentUsecase"
	paymentusecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{cfg, paymentUsecase}
}
