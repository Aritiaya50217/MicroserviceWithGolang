package inventoryhandler

import (
	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	inventoryUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}
	inventoryQueueHandler        struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{cfg, inventoryUsecase}
}
