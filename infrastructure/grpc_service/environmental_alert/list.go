package environmental_alert_service

import (
	"context"
	"monitoring_service/domain/entity"
	"monitoring_service/domain/usecase/environmental_alert"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
)

func (s *EnvironmentalAlertService) ListEnvironmentalAlerts(ctx context.Context, req *proto_environmental_alert.ListEnvironmentalAlertsRequest) (*proto_environmental_alert.ListEnvironmentalAlertsResponse, error) {
	usecaseReq := s.convertListEnvironmentalAlertsRequest(req)
	alerts, total, err := s.eaUsecase.List(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}

	protoAlerts := make([]*proto_environmental_alert.EnvironmentalAlert, len(alerts))
	for i, alert := range alerts {
		protoAlerts[i] = s.convertToProtoEnvironmentalAlert(alert)
	}

	totalPages := s.helper.CalculateTotalPages(total, int64(req.Pagination.PageSize))
	return &proto_environmental_alert.ListEnvironmentalAlertsResponse{
		Alerts: protoAlerts,
		Pagination: &proto_common.PaginationResponse{
			Page:       int32(req.Pagination.Page),
			PageSize:   int32(req.Pagination.PageSize),
			TotalPages: int32(totalPages),
			Total:      int32(total),
		},
	}, nil
}

func (s *EnvironmentalAlertService) convertListEnvironmentalAlertsRequest(req *proto_environmental_alert.ListEnvironmentalAlertsRequest) environmental_alert.ListEnvironmentalAlertRequest {
	return environmental_alert.ListEnvironmentalAlertRequest{
		Pagination: common.Pagination{
			Page:      int(req.Pagination.Page),
			PageSize:  int(req.Pagination.PageSize),
			SortBy:    req.Pagination.SortBy,
			SortOrder: req.Pagination.SortOrder,
			Search:    req.Pagination.Search,
		},
		Filters: s.convertToEntityAlertFilters(req.Filters),
	}
}

func (s *EnvironmentalAlertService) convertToEntityAlertFilters(filters *proto_environmental_alert.AlertFilters) entity.AlertFilters {
	var entityFilters entity.AlertFilters

	if filters == nil {
		return entityFilters
	}

	entityFilters = entity.AlertFilters{
		DeviceID:        filters.DeviceId,
		AlertType:       entity.AlertType(filters.AlertType),
		Status:          entity.Status(filters.Status),
		Severity:        entity.Severity(filters.Severity),
		Priority:        int(filters.Priority),
		IsOverdue:       filters.IsOverdue,
		EscalationLevel: int(filters.EscalationLevel),
		AutoResolve:     filters.AutoResolve,
		CreatedBy:       filters.CreatedBy,
		AcknowledgedBy:  filters.AcknowledgedBy,
		ResolvedBy:      filters.ResolvedBy,
	}

	if filters.StartDate != nil {
		startTime := filters.StartDate.AsTime()
		entityFilters.StartDate = &startTime
	}
	if filters.EndDate != nil {
		endTime := filters.EndDate.AsTime()
		entityFilters.EndDate = &endTime
	}

	return entityFilters
}
