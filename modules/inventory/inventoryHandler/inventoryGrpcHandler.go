package inventoryhandler

import (
	inventoryUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
		// inventoryPb.UnimplementedInventoryGrpcServiceServer
	}
)

func NewInventoryGrpcHttpHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	// return &inventoryGrpcHandler{inventoryUsecase}
	return nil
}

// func (g *inventoryGrpcHandler) IsAvaliableToSell(ctx context.Context, req *inventoryPb.IsAvaliableToSellReq) (*inventoryPb.IsAvaliableToSellRes, error) {
// 	return nil, nil
// }
