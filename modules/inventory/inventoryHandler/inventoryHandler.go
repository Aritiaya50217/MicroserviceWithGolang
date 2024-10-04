package inventoryhandler

import (
	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	inventoryUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}
	inventoryHttpHandler        struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{cfg, inventoryUsecase}
}
