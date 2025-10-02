package environmental_alert

import (
	"context"

	"monitoring_service/domain/entity"
	"monitoring_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListEnvironmentalAlertRequest struct {
	Pagination common.Pagination
	Filters    entity.AlertFilters
}

type ListEnvironmentalAlertUsecaseI interface {
	Execute(ctx context.Context, req ListEnvironmentalAlertRequest) ([]*entity.EnvironmentalAlert, int64, error)
}

type ListEnvironmentalAlertUsecase struct {
	repo repository.EnvironmentalAlertRepository
}

func NewListEnvironmentalAlertUsecase(repo repository.EnvironmentalAlertRepository) *ListEnvironmentalAlertUsecase {
	return &ListEnvironmentalAlertUsecase{
		repo: repo,
	}
}

func (u *ListEnvironmentalAlertUsecase) Execute(ctx context.Context, req ListEnvironmentalAlertRequest) ([]*entity.EnvironmentalAlert, int64, error) {
	alerts, total, err := u.repo.GetAll(ctx, req.Pagination, &req.Filters)
	if err != nil {
		return nil, 0, ErrListAlert
	}
	return alerts, total, nil
}
