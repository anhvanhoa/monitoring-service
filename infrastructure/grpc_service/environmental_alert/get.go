package environmental_alert_service

import (
	"context"

	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
)

func (s *EnvironmentalAlertService) GetEnvironmentalAlert(ctx context.Context, req *proto_environmental_alert.GetEnvironmentalAlertRequest) (*proto_environmental_alert.GetEnvironmentalAlertResponse, error) {
	ea, err := s.eaUsecase.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_environmental_alert.GetEnvironmentalAlertResponse{
		Alert: s.convertToProtoEnvironmentalAlert(ea),
	}, nil
}
