package environmental_alert_service

import (
	"context"

	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
)

func (s *EnvironmentalAlertService) DeleteEnvironmentalAlert(ctx context.Context, req *proto_environmental_alert.DeleteEnvironmentalAlertRequest) (*proto_environmental_alert.DeleteEnvironmentalAlertResponse, error) {
	success, err := s.eaUsecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	message := "Cảnh báo đã được xóa thành công"
	if !success {
		message = "Không thể xóa cảnh báo"
	}

	return &proto_environmental_alert.DeleteEnvironmentalAlertResponse{
		Success: success,
		Message: message,
	}, nil
}
