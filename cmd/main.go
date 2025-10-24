package main

import (
	"context"
	"monitoring_service/bootstrap"
	"monitoring_service/infrastructure/grpc_client"
	"monitoring_service/infrastructure/grpc_service"
	environmental_alert_service "monitoring_service/infrastructure/grpc_service/environmental_alert"

	gc "github.com/anhvanhoa/service-core/domain/grpc_client"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	permissionClient := grpc_client.NewPermissionClient(clientFactory.GetClient(env.PermissionServiceAddr))

	proto_environmentalAlertService := environmental_alert_service.NewEnvironmentalAlertService(app.Repo)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		app.Cache,
		proto_environmentalAlertService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionClient.PermissionServiceClient.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to register permission: " + err.Error())
	}
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
