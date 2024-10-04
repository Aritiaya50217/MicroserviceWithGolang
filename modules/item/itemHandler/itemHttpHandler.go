package itemhandler

import (
	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	itemUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}
	itemHttpHandler        struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{cfg, itemUsecase}
}
