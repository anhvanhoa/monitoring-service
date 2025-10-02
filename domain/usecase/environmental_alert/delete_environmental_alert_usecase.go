package environmental_alert

import (
	"context"

	"monitoring_service/domain/repository"
)

type DeleteEnvironmentalAlertUsecaseI interface {
	Execute(ctx context.Context, id string) (bool, error)
}

type DeleteEnvironmentalAlertUsecase struct {
	repo repository.EnvironmentalAlertRepository
}

func NewDeleteEnvironmentalAlertUsecase(repo repository.EnvironmentalAlertRepository) DeleteEnvironmentalAlertUsecaseI {
	return &DeleteEnvironmentalAlertUsecase{
		repo: repo,
	}
}

func (u *DeleteEnvironmentalAlertUsecase) Execute(ctx context.Context, id string) (bool, error) {
	_, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return false, ErrAlertNotFound
	}

	if err := u.repo.Delete(ctx, id); err != nil {
		return false, ErrDeleteAlert
	}

	return true, nil
}
