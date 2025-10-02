package environmental_alert

import (
	"context"

	"monitoring_service/domain/entity"
	"monitoring_service/domain/repository"
)

type UpdateEnvironmentalAlertRequest struct {
	ID               string
	DeviceID         string
	AlertType        entity.AlertType
	CurrentValue     float64
	ThresholdValue   float64
	ThresholdType    entity.ThresholdType
	Severity         entity.Severity
	Priority         int
	Status           entity.Status
	AutoResolve      bool
	AutoActionTaken  string
	EscalationLevel  int
	ImpactAssessment string
	ResolutionNotes  string
	UpdatedBy        string
}

type UpdateEnvironmentalAlertUsecaseI interface {
	Execute(ctx context.Context, req *UpdateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error)
}

type UpdateEnvironmentalAlertUsecase struct {
	repo repository.EnvironmentalAlertRepository
}

func NewUpdateEnvironmentalAlertUsecase(repo repository.EnvironmentalAlertRepository) *UpdateEnvironmentalAlertUsecase {
	return &UpdateEnvironmentalAlertUsecase{
		repo: repo,
	}
}

func (u *UpdateEnvironmentalAlertUsecase) Execute(ctx context.Context, req *UpdateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error) {
	alert, err := u.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrAlertNotFound
	}

	if err := alert.Validate(); err != nil {
		return nil, err
	}
	if err := u.repo.Update(ctx, alert); err != nil {
		return nil, ErrUpdateAlert
	}

	return alert, nil
}
