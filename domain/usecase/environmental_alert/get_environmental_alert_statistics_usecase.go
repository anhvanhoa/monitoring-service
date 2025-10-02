package environmental_alert

import (
	"context"

	"monitoring_service/domain/entity"
	"monitoring_service/domain/repository"
)

type GetEnvironmentalAlertStatisticsRequest struct {
}

type GetEnvironmentalAlertStatisticsUsecaseI interface {
	Execute(ctx context.Context, req *GetEnvironmentalAlertStatisticsRequest) (*entity.AlertStatistics, error)
}

type GetEnvironmentalAlertStatisticsUsecase struct {
	repo repository.EnvironmentalAlertRepository
}

func NewGetEnvironmentalAlertStatisticsUsecase(repo repository.EnvironmentalAlertRepository) GetEnvironmentalAlertStatisticsUsecaseI {
	return &GetEnvironmentalAlertStatisticsUsecase{
		repo: repo,
	}
}

func (u *GetEnvironmentalAlertStatisticsUsecase) Execute(ctx context.Context, req *GetEnvironmentalAlertStatisticsRequest) (*entity.AlertStatistics, error) {
	stats, err := u.repo.GetAlertStatistics(ctx)
	if err != nil {
		return nil, ErrGetAlertStatistics
	}

	return stats, nil
}
