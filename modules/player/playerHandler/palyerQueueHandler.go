package playerhandler

import (
	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	playerusecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQuqueHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

// คอยรับ message จาก kafka
func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerusecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQuqueHandler{cfg: cfg, playerUsecase: playerUsecase}
}
