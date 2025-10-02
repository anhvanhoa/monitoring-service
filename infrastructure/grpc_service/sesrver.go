package grpc_service

import (
	"monitoring_service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	proto_environmentalAlertService proto_environmental_alert.EnvironmentalAlertServiceServer,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_environmental_alert.RegisterEnvironmentalAlertServiceServer(server, proto_environmentalAlertService)
		},
	)
}
