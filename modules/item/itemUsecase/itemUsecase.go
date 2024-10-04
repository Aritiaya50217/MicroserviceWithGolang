package itemusecase

import itemRepository "github.com/Aritiaya50217/MicroserviceWithGolang/modules/item/itemRepository"

type (
	ItemUsecaseService interface{}
	itemUsecase        struct {
		itemReposiroty itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemReposiroty itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemReposiroty}
}
