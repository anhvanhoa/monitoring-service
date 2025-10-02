package environmental_alert

import (
	"context"
	"encoding/json"
	"time"

	"monitoring_service/domain/entity"
	"monitoring_service/domain/repository"
)

type CreateEnvironmentalAlertRequest struct {
	DeviceID         string
	AlertType        entity.AlertType
	CurrentValue     float64
	ThresholdValue   float64
	ThresholdType    entity.ThresholdType
	Severity         entity.Severity
	Priority         int
	AutoResolve      bool
	AutoActionTaken  string
	EscalationLevel  int
	ImpactAssessment string
	CreatedBy        string
}

type CreateEnvironmentalAlertUsecaseI interface {
	Execute(ctx context.Context, req *CreateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error)
}

type CreateEnvironmentalAlertUsecase struct {
	repo repository.EnvironmentalAlertRepository
}

func NewCreateEnvironmentalAlertUsecase(repo repository.EnvironmentalAlertRepository) CreateEnvironmentalAlertUsecaseI {
	return &CreateEnvironmentalAlertUsecase{
		repo: repo,
	}
}

func (u *CreateEnvironmentalAlertUsecase) Execute(ctx context.Context, req *CreateEnvironmentalAlertRequest) (*entity.EnvironmentalAlert, error) {
	alert := &entity.EnvironmentalAlert{
		DeviceID:         req.DeviceID,
		AlertType:        req.AlertType,
		CurrentValue:     req.CurrentValue,
		ThresholdValue:   req.ThresholdValue,
		ThresholdType:    req.ThresholdType,
		Severity:         req.Severity,
		Priority:         req.Priority,
		TriggeredAt:      time.Now(),
		AutoResolve:      req.AutoResolve,
		AutoActionTaken:  req.AutoActionTaken,
		EscalationLevel:  req.EscalationLevel,
		ImpactAssessment: req.ImpactAssessment,
		CreatedBy:        req.CreatedBy,
		Status:           entity.StatusActive,
		CreatedAt:        time.Now(),
	}

	notificationData := entity.NotificationSent{
		Email:      true,
		SMS:        false,
		Push:       false,
		Webhook:    false,
		SentAt:     time.Now(),
		Recipients: []string{},
	}

	notificationJSON, err := json.Marshal(notificationData)
	if err != nil {
		return nil, err
	}

	alert.NotificationSent = string(notificationJSON)

	if err := alert.Validate(); err != nil {
		return nil, err
	}

	if err := u.repo.Create(ctx, alert); err != nil {
		return nil, err
	}

	return alert, nil
}
