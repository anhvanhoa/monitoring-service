package environmental_alert_service

import (
	"context"
	"monitoring_service/domain/entity"
	"monitoring_service/domain/usecase/environmental_alert"

	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
)

func (s *EnvironmentalAlertService) UpdateEnvironmentalAlert(ctx context.Context, req *proto_environmental_alert.UpdateEnvironmentalAlertRequest) (*proto_environmental_alert.UpdateEnvironmentalAlertResponse, error) {
	usecaseReq := s.convertUpdateEnvironmentalAlertRequest(req)
	ea, err := s.eaUsecase.Update(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	return &proto_environmental_alert.UpdateEnvironmentalAlertResponse{
		Alert: s.convertToProtoEnvironmentalAlert(ea),
	}, nil
}

func (s *EnvironmentalAlertService) convertUpdateEnvironmentalAlertRequest(req *proto_environmental_alert.UpdateEnvironmentalAlertRequest) *environmental_alert.UpdateEnvironmentalAlertRequest {
	return &environmental_alert.UpdateEnvironmentalAlertRequest{
		ID:               req.Id,
		DeviceID:         req.DeviceId,
		AlertType:        entity.AlertType(req.AlertType),
		CurrentValue:     req.CurrentValue,
		ThresholdValue:   req.ThresholdValue,
		ThresholdType:    entity.ThresholdType(req.ThresholdType),
		Severity:         entity.Severity(req.Severity),
		Priority:         int(req.Priority),
		Status:           entity.Status(req.Status),
		AutoResolve:      req.AutoResolve,
		AutoActionTaken:  req.AutoActionTaken,
		EscalationLevel:  int(req.EscalationLevel),
		ImpactAssessment: req.ImpactAssessment,
		ResolutionNotes:  req.ResolutionNotes,
		UpdatedBy:        req.UpdatedBy,
	}
}
