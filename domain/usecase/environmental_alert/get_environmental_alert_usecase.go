package environmental_alert

import (
	"context"

	"monitoring_service/domain/entity"
	"monitoring_service/domain/repository"
)

type GetEnvironmentalAlertUsecaseI interface {
	Execute(ctx context.Context, id string) (*entity.EnvironmentalAlert, error)
}

type GetEnvironmentalAlertUsecase struct {
	repo repository.EnvironmentalAlertRepository
}

func NewGetEnvironmentalAlertUsecase(repo repository.EnvironmentalAlertRepository) *GetEnvironmentalAlertUsecase {
	return &GetEnvironmentalAlertUsecase{
		repo: repo,
	}
}

func (u *GetEnvironmentalAlertUsecase) Execute(ctx context.Context, id string) (*entity.EnvironmentalAlert, error) {
	alert, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrAlertNotFound
	}

	return alert, nil
}
