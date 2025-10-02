package environmental_alert_service

import (
	"monitoring_service/domain/usecase/environmental_alert"
	"monitoring_service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/utils"
	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
)

type EnvironmentalAlertService struct {
	proto_environmental_alert.UnsafeEnvironmentalAlertServiceServer
	eaUsecase environmental_alert.EnvironmentalAlertUsecase
	helper    utils.Helper
}

func NewEnvironmentalAlertService(repo repo.Repositories) proto_environmental_alert.EnvironmentalAlertServiceServer {
	eaUsecase := environmental_alert.NewEnvironmentalAlertUsecase(repo.EnvironmentalAlertRepository())
	return &EnvironmentalAlertService{
		eaUsecase: eaUsecase,
		helper:    utils.NewHelper(),
	}
}
