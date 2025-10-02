package repository

import (
	"context"

	"monitoring_service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type EnvironmentalAlertRepository interface {
	Create(ctx context.Context, alert *entity.EnvironmentalAlert) error
	GetByID(ctx context.Context, id string) (*entity.EnvironmentalAlert, error)
	Update(ctx context.Context, alert *entity.EnvironmentalAlert) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, pagination common.Pagination, filters *entity.AlertFilters) ([]*entity.EnvironmentalAlert, int64, error)
	GetAlertStatistics(ctx context.Context) (*entity.AlertStatistics, error)
}
