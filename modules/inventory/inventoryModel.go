package inventory

import (
	"github.com/Aritiaya50217/MicroserviceWithGolang/modules/item"
	"github.com/Aritiaya50217/MicroserviceWithGolang/modules/models"
)

type (
	UpdateInventoryReq struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}
	ItemInInventory struct {
		InventoryId string `json:"inventory_id"`
		PlayerId    string `json:"player_id"`
		*item.ItemShowCase
	}
	InventorySearchReq struct {
		models.PaginateRep
	}
	RollbackPlayerInventoryReq struct {
		InventoryId string `json:"inventory_id"`
		PlayerId    string `json:"player_id"`
		ItemId      string `json:"item_id"`
	}
)
