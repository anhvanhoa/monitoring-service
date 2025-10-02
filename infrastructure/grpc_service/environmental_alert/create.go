package environmental_alert_service

import (
	"context"
	"monitoring_service/domain/entity"
	"monitoring_service/domain/usecase/environmental_alert"

	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *EnvironmentalAlertService) CreateEnvironmentalAlert(ctx context.Context, req *proto_environmental_alert.CreateEnvironmentalAlertRequest) (*proto_environmental_alert.CreateEnvironmentalAlertResponse, error) {
	usecaseReq := s.convertCreateEnvironmentalAlertRequest(req)
	ea, err := s.eaUsecase.Create(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	return &proto_environmental_alert.CreateEnvironmentalAlertResponse{
		Alert: s.convertToProtoEnvironmentalAlert(ea),
	}, nil
}

func (s *EnvironmentalAlertService) convertCreateEnvironmentalAlertRequest(req *proto_environmental_alert.CreateEnvironmentalAlertRequest) *environmental_alert.CreateEnvironmentalAlertRequest {
	return &environmental_alert.CreateEnvironmentalAlertRequest{
		DeviceID:         req.DeviceId,
		AlertType:        entity.AlertType(req.AlertType),
		CurrentValue:     req.CurrentValue,
		ThresholdValue:   req.ThresholdValue,
		ThresholdType:    entity.ThresholdType(req.ThresholdType),
		Severity:         entity.Severity(req.Severity),
		Priority:         int(req.Priority),
		AutoResolve:      req.AutoResolve,
		AutoActionTaken:  req.AutoActionTaken,
		EscalationLevel:  int(req.EscalationLevel),
		ImpactAssessment: req.ImpactAssessment,
		CreatedBy:        req.CreatedBy,
	}
}

func (s *EnvironmentalAlertService) convertToProtoEnvironmentalAlert(ea *entity.EnvironmentalAlert) *proto_environmental_alert.EnvironmentalAlert {
	eaProto := &proto_environmental_alert.EnvironmentalAlert{
		Id:               ea.ID,
		DeviceId:         ea.DeviceID,
		CurrentValue:     ea.CurrentValue,
		ThresholdValue:   ea.ThresholdValue,
		AlertType:        string(ea.AlertType),
		ThresholdType:    string(ea.ThresholdType),
		Severity:         string(ea.Severity),
		Priority:         int32(ea.Priority),
		AutoResolve:      ea.AutoResolve,
		AutoActionTaken:  ea.AutoActionTaken,
		EscalationLevel:  int32(ea.EscalationLevel),
		ImpactAssessment: ea.ImpactAssessment,
		CreatedBy:        ea.CreatedBy,
		CreatedAt:        timestamppb.New(ea.CreatedAt),
		Status:           string(ea.Status),
		AcknowledgedBy:   ea.AcknowledgedBy,
		ResolvedBy:       ea.ResolvedBy,
		ResolutionNotes:  ea.ResolutionNotes,
	}
	if ea.AcknowledgedAt != nil {
		eaProto.AcknowledgedAt = timestamppb.New(*ea.AcknowledgedAt)
	}
	if ea.ResolvedAt != nil {
		eaProto.ResolvedAt = timestamppb.New(*ea.ResolvedAt)
	}
	if ea.UpdatedAt != nil {
		eaProto.UpdatedAt = timestamppb.New(*ea.UpdatedAt)
	}
	return eaProto
}
