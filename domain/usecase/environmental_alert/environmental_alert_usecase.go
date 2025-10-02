package environmental_alert

import (
	"context"

	"monitoring_service/domain/entity"
	"monitoring_service/domain/repository"
)

type EnvironmentalAlertUsecase interface {
	Create(ctx context.Context, req *CreateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error)
	Get(ctx context.Context, id string) (*entity.EnvironmentalAlert, error)
	Update(ctx context.Context, req *UpdateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error)
	Delete(ctx context.Context, id string) (bool, error)
	List(ctx context.Context, req ListEnvironmentalAlertRequest) ([]*entity.EnvironmentalAlert, int64, error)
	GetStatistics(ctx context.Context, req *GetEnvironmentalAlertStatisticsRequest) (*entity.AlertStatistics, error)
}

type EnvironmentalAlertUsecaseImpl struct {
	createUsecase CreateEnvironmentalAlertUsecaseI
	getUsecase    GetEnvironmentalAlertUsecaseI
	updateUsecase UpdateEnvironmentalAlertUsecaseI
	deleteUsecase DeleteEnvironmentalAlertUsecaseI
	listUsecase   ListEnvironmentalAlertUsecaseI
	statsUsecase  GetEnvironmentalAlertStatisticsUsecaseI
}

func NewEnvironmentalAlertUsecase(repo repository.EnvironmentalAlertRepository) EnvironmentalAlertUsecase {
	return &EnvironmentalAlertUsecaseImpl{
		createUsecase: NewCreateEnvironmentalAlertUsecase(repo),
		getUsecase:    NewGetEnvironmentalAlertUsecase(repo),
		updateUsecase: NewUpdateEnvironmentalAlertUsecase(repo),
		deleteUsecase: NewDeleteEnvironmentalAlertUsecase(repo),
		listUsecase:   NewListEnvironmentalAlertUsecase(repo),
		statsUsecase:  NewGetEnvironmentalAlertStatisticsUsecase(repo),
	}
}

func (u *EnvironmentalAlertUsecaseImpl) Create(ctx context.Context, req *CreateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *EnvironmentalAlertUsecaseImpl) Get(ctx context.Context, id string) (*entity.EnvironmentalAlert, error) {
	return u.getUsecase.Execute(ctx, id)
}

func (u *EnvironmentalAlertUsecaseImpl) Update(ctx context.Context, req *UpdateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error) {
	return u.updateUsecase.Execute(ctx, req)
}

func (u *EnvironmentalAlertUsecaseImpl) Delete(ctx context.Context, id string) (bool, error) {
	return u.deleteUsecase.Execute(ctx, id)
}

func (u *EnvironmentalAlertUsecaseImpl) List(ctx context.Context, req ListEnvironmentalAlertRequest) ([]*entity.EnvironmentalAlert, int64, error) {
	return u.listUsecase.Execute(ctx, req)
}

func (u *EnvironmentalAlertUsecaseImpl) GetStatistics(ctx context.Context, req *GetEnvironmentalAlertStatisticsRequest) (*entity.AlertStatistics, error) {
	return u.statsUsecase.Execute(ctx, req)
}
