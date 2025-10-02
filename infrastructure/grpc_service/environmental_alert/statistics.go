package environmental_alert_service

import (
	"context"
	"monitoring_service/domain/usecase/environmental_alert"

	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *EnvironmentalAlertService) GetEnvironmentalAlertStatistics(ctx context.Context, req *emptypb.Empty) (*proto_environmental_alert.GetEnvironmentalAlertStatisticsResponse, error) {
	usecaseReq := &environmental_alert.GetEnvironmentalAlertStatisticsRequest{}
	stats, err := s.eaUsecase.GetStatistics(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}

	return &proto_environmental_alert.GetEnvironmentalAlertStatisticsResponse{
		Statistics: &proto_environmental_alert.AlertStatistics{
			TotalAlerts:    stats.TotalAlerts,
			ActiveAlerts:   stats.ActiveAlerts,
			ResolvedAlerts: stats.ResolvedAlerts,
			OverdueAlerts:  stats.OverdueAlerts,
		},
	}, nil
}
