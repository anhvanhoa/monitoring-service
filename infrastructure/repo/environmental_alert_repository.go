package repo

import (
	"context"

	"monitoring_service/domain/entity"
	"monitoring_service/domain/repository"

	"github.com/go-pg/pg/v10"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type EnvironmentalAlertRepositoryImpl struct {
	db     *pg.DB
	helper utils.Helper
}

func NewEnvironmentalAlertRepository(db *pg.DB, helper utils.Helper) repository.EnvironmentalAlertRepository {
	return &EnvironmentalAlertRepositoryImpl{
		db:     db,
		helper: helper,
	}
}

func (r *EnvironmentalAlertRepositoryImpl) Create(ctx context.Context, alert *entity.EnvironmentalAlert) error {
	_, err := r.db.Model(alert).Context(ctx).Insert()
	return err
}

func (r *EnvironmentalAlertRepositoryImpl) GetByID(ctx context.Context, id string) (*entity.EnvironmentalAlert, error) {
	alert := &entity.EnvironmentalAlert{}
	err := r.db.Model(alert).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return alert, nil
}

func (r *EnvironmentalAlertRepositoryImpl) Update(ctx context.Context, alert *entity.EnvironmentalAlert) error {
	_, err := r.db.Model(alert).Context(ctx).Where("id = ?", alert.ID).UpdateNotZero()
	return err
}

func (r *EnvironmentalAlertRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model((*entity.EnvironmentalAlert)(nil)).Context(ctx).Where("id = ?", id).Delete()
	return err
}

func (r *EnvironmentalAlertRepositoryImpl) GetAll(ctx context.Context, pagination common.Pagination, filters *entity.AlertFilters) ([]*entity.EnvironmentalAlert, int64, error) {
	var alerts []*entity.EnvironmentalAlert
	query := r.db.Model(&alerts).Context(ctx)

	if filters != nil {
		if filters.DeviceID != "" {
			query = query.Where("device_id = ?", filters.DeviceID)
		}
		if filters.AlertType != "" {
			query = query.Where("alert_type = ?", filters.AlertType)
		}
		if filters.Status != "" {
			query = query.Where("status = ?", filters.Status)
		}
		if filters.Severity != "" {
			query = query.Where("severity = ?", filters.Severity)
		}
		if filters.Priority > 0 {
			query = query.Where("priority = ?", filters.Priority)
		}
		if filters.StartDate != nil {
			query = query.Where("triggered_at >= ?", *filters.StartDate)
		}
		if filters.EndDate != nil {
			query = query.Where("triggered_at <= ?", *filters.EndDate)
		}
		if filters.IsOverdue {
			query = query.Where("status IN (?, ?)", entity.StatusActive, entity.StatusAcknowledged)
		}
		if filters.EscalationLevel > 0 {
			query = query.Where("escalation_level = ?", filters.EscalationLevel)
		}
		if filters.AutoResolve {
			query = query.Where("auto_resolve = ?", true)
		}
		if filters.CreatedBy != "" {
			query = query.Where("created_by = ?", filters.CreatedBy)
		}
		if filters.AcknowledgedBy != "" {
			query = query.Where("acknowledged_by = ?", filters.AcknowledgedBy)
		}
		if filters.ResolvedBy != "" {
			query = query.Where("resolved_by = ?", filters.ResolvedBy)
		}
	}

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	if pagination.Page > 0 && pagination.PageSize > 0 {
		offset := r.helper.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Offset(offset)
	}
	query = query.Limit(pagination.PageSize)

	query = query.Order("created_at DESC")
	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	if filters != nil && filters.IsOverdue {
		var overdueAlerts []*entity.EnvironmentalAlert
		for _, alert := range alerts {
			if alert.IsOverdue() {
				overdueAlerts = append(overdueAlerts, alert)
			}
		}
		alerts = overdueAlerts
		total = len(overdueAlerts)
	}

	return alerts, int64(total), nil
}

func (r *EnvironmentalAlertRepositoryImpl) GetAlertStatistics(ctx context.Context) (*entity.AlertStatistics, error) {
	stats := &entity.AlertStatistics{}

	totalCount, err := r.db.Model((*entity.EnvironmentalAlert)(nil)).Context(ctx).Count()
	if err != nil {
		return nil, err
	}
	stats.TotalAlerts = int64(totalCount)

	activeCount, err := r.db.Model((*entity.EnvironmentalAlert)(nil)).Context(ctx).Where("status = ?", entity.StatusActive).Count()
	if err != nil {
		return nil, err
	}
	stats.ActiveAlerts = int64(activeCount)

	resolvedCount, err := r.db.Model((*entity.EnvironmentalAlert)(nil)).Context(ctx).Where("status = ?", entity.StatusResolved).Count()
	if err != nil {
		return nil, err
	}
	stats.ResolvedAlerts = int64(resolvedCount)

	var alerts []*entity.EnvironmentalAlert
	err = r.db.Model(&alerts).Context(ctx).
		Where("status IN (?, ?)", entity.StatusActive, entity.StatusAcknowledged).
		Select()
	if err != nil {
		return nil, err
	}

	var overdueCount int64
	for _, alert := range alerts {
		if alert.IsOverdue() {
			overdueCount++
		}
	}
	stats.OverdueAlerts = overdueCount

	return stats, nil
}
