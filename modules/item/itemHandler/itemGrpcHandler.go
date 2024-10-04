package itemhandler

import (
	"context"

	itemPb "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemPb"
	itemUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemPb.UnimplementedItemGrpcServiceServer
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase: itemUsecase}
}
func (g *itemGrpcHandler) FindItemsInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	// return g.itemUsecase.FindItemInIds(ctx, req)
	return nil, nil
}