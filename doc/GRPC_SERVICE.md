# gRPC Service Implementation

## Production Service gRPC Implementation

### Service Structure

```
infrastructure/grpc_service/
├── monitoring_service.go          # Main Production Service implementation
├── harvest_record_service.go     # Harvest Record service implementation
├── pest_disease_service.go       # Pest Disease Record service implementation
└── base.go                       # Base service structure
```

### Base Service Structure

```go
// file: infrastructure/grpc_service/base.go
package grpc_service

import (
	"monitoring_service/domain/usecase/harvest_record"
	"monitoring_service/domain/usecase/pest_disease_record"
	pb "monitoring_service/proto_generated"
)

type ProductionService struct {
	pb.UnimplementedProductionServiceServer
	harvestRecordUsecase     harvest_record.HarvestRecordUsecase
	pestDiseaseRecordUsecase pest_disease_record.PestDiseaseRecordUsecase
}

func NewProductionService(
	harvestRecordUsecase harvest_record.HarvestRecordUsecase,
	pestDiseaseRecordUsecase pest_disease_record.PestDiseaseRecordUsecase,
) pb.ProductionServiceServer {
	return &ProductionService{
		harvestRecordUsecase:     harvestRecordUsecase,
		pestDiseaseRecordUsecase: pestDiseaseRecordUsecase,
	}
}
```

### Harvest Record Service Implementation

```go
// file: infrastructure/grpc_service/harvest_record_service.go
package grpc_service

import (
	"context"
	"monitoring_service/domain/entity"
	"monitoring_service/domain/usecase/harvest_record"
	pb "monitoring_service/proto_generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ProductionService) CreateHarvestRecord(
	ctx context.Context,
	req *pb.CreateHarvestRecordRequest,
) (*pb.CreateHarvestRecordResponse, error) {
	// Convert proto request to usecase request
	usecaseReq := &harvest_record.CreateHarvestRecordRequest{
		PlantingCycleID:    req.PlantingCycleId,
		HarvestDate:        req.HarvestDate.AsTime(),
		HarvestTime:        req.HarvestTime.AsTime(),
		QuantityKg:         req.QuantityKg,
		QualityGrade:       req.QualityGrade,
		SizeClassification: req.SizeClassification,
		MarketPricePerKg:   req.MarketPricePerKg,
		LaborHours:         req.LaborHours,
		LaborCost:          req.LaborCost,
		PackagingCost:      req.PackagingCost,
		StorageLocation:    req.StorageLocation,
		StorageTemperature: req.StorageTemperature,
		BuyerInformation:   req.BuyerInformation,
		DeliveryDate:       req.DeliveryDate.AsTime(),
		WeatherAtHarvest:   req.WeatherAtHarvest,
		PlantHealthRating:  int(req.PlantHealthRating),
		Notes:              req.Notes,
		Images:             req.Images,
		CreatedBy:          req.CreatedBy,
	}
	
	// Call usecase
	harvestRecord, err := s.harvestRecordUsecase.Create(ctx, usecaseReq)
	if err != nil {
		return nil, s.handleError(err)
	}
	
	// Convert entity to proto response
	protoRecord := s.convertToProtoHarvestRecord(harvestRecord)
	
	return &pb.CreateHarvestRecordResponse{
		HarvestRecord: protoRecord,
	}, nil
}

func (s *ProductionService) GetHarvestRecord(
	ctx context.Context,
	req *pb.GetHarvestRecordRequest,
) (*pb.GetHarvestRecordResponse, error) {
	harvestRecord, err := s.harvestRecordUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, s.handleError(err)
	}
	
	protoRecord := s.convertToProtoHarvestRecord(harvestRecord)
	
	return &pb.GetHarvestRecordResponse{
		HarvestRecord: protoRecord,
	}, nil
}

func (s *ProductionService) ListHarvestRecords(
	ctx context.Context,
	req *pb.ListHarvestRecordsRequest,
) (*pb.ListHarvestRecordsResponse, error) {
	// Convert proto pagination to common pagination
	pagination := common.Pagination{
		Page:     req.Pagination.Page,
		PageSize: req.Pagination.PageSize,
		SortBy:   req.Pagination.SortBy,
		SortOrder: req.Pagination.SortOrder,
	}
	
	// Convert proto filter to entity filter
	filter := entity.FilterHarvestRecord{
		HarvestDate:        req.Filter.HarvestDate.AsTime(),
		QualityGrade:       req.Filter.QualityGrade,
		SizeClassification: req.Filter.SizeClassification,
		MarketPricePerKg:   req.Filter.MarketPricePerKg,
		TotalRevenue:       req.Filter.TotalRevenue,
		PlantHealthRating:  int(req.Filter.PlantHealthRating),
		Notes:              req.Filter.Notes,
		Images:             req.Filter.Images,
		CreatedBy:          req.Filter.CreatedBy,
		CreatedAt:          req.Filter.CreatedAt.AsTime(),
	}
	
	harvestRecords, total, err := s.harvestRecordUsecase.List(ctx, pagination, filter)
		if err != nil {
		return nil, s.handleError(err)
	}
	
	// Convert entities to proto
	var protoRecords []*pb.HarvestRecord
	for _, record := range harvestRecords {
		protoRecords = append(protoRecords, s.convertToProtoHarvestRecord(record))
	}
	
	return &pb.ListHarvestRecordsResponse{
		HarvestRecords: protoRecords,
		Total:          total,
	}, nil
}

func (s *ProductionService) convertToProtoHarvestRecord(record *entity.HarvestRecord) *pb.HarvestRecord {
	protoRecord := &pb.HarvestRecord{
		Id:                 record.ID,
		PlantingCycleId:    record.PlantingCycleID,
		QuantityKg:         record.QuantityKg,
		QualityGrade:       record.QualityGrade,
		SizeClassification: record.SizeClassification,
		MarketPricePerKg:   record.MarketPricePerKg,
		TotalRevenue:       record.TotalRevenue,
		LaborHours:         record.LaborHours,
		LaborCost:          record.LaborCost,
		PackagingCost:      record.PackagingCost,
		StorageLocation:    record.StorageLocation,
		StorageTemperature: record.StorageTemperature,
		BuyerInformation:   record.BuyerInformation,
		WeatherAtHarvest:   record.WeatherAtHarvest,
		PlantHealthRating:  int32(record.PlantHealthRating),
		Notes:              record.Notes,
		Images:             record.Images,
		CreatedBy:          record.CreatedBy,
		CreatedAt:          timestamppb.New(record.CreatedAt),
	}
	
	if record.HarvestDate != nil {
		protoRecord.HarvestDate = timestamppb.New(*record.HarvestDate)
	}
	if record.HarvestTime != nil {
		protoRecord.HarvestTime = timestamppb.New(*record.HarvestTime)
	}
	if record.DeliveryDate != nil {
		protoRecord.DeliveryDate = timestamppb.New(*record.DeliveryDate)
	}
	if record.UpdatedAt != nil {
		protoRecord.UpdatedAt = timestamppb.New(*record.UpdatedAt)
	}
	
	return protoRecord
}

func (s *ProductionService) handleError(err error) error {
	switch err {
	case harvest_record.ErrHarvestRecordNotFound:
		return status.Error(codes.NotFound, "Harvest record not found")
	case harvest_record.ErrInvalidQualityGrade:
		return status.Error(codes.InvalidArgument, "Invalid quality grade")
	case harvest_record.ErrInvalidSizeClassification:
		return status.Error(codes.InvalidArgument, "Invalid size classification")
	case harvest_record.ErrInvalidPlantHealthRating:
		return status.Error(codes.InvalidArgument, "Invalid plant health rating")
	default:
		return status.Error(codes.Internal, "Internal server error")
	}
}
```

### Pest Disease Record Service Implementation

```go
// file: infrastructure/grpc_service/pest_disease_service.go
package grpc_service

import (
	"context"
	"monitoring_service/domain/entity"
	"monitoring_service/domain/usecase/pest_disease_record"
	pb "monitoring_service/proto_generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ProductionService) CreatePestDiseaseRecord(
	ctx context.Context,
	req *pb.CreatePestDiseaseRecordRequest,
) (*pb.CreatePestDiseaseRecordResponse, error) {
	// Convert proto request to usecase request
	usecaseReq := &pest_disease_record.CreatePestDiseaseRecordRequest{
		PlantingCycleID:        req.PlantingCycleId,
		Type:                   req.Type,
		Name:                   req.Name,
		ScientificName:         req.ScientificName,
		Severity:               req.Severity,
		AffectedAreaPercentage: req.AffectedAreaPercentage,
		AffectedPlantCount:     int(req.AffectedPlantCount),
		DetectionDate:          req.DetectionDate.AsTime(),
		DetectionMethod:        req.DetectionMethod,
		Symptoms:               req.Symptoms,
		TreatmentApplied:       req.TreatmentApplied,
		TreatmentDate:          req.TreatmentDate.AsTime(),
		TreatmentCost:          req.TreatmentCost,
		TreatmentDurationDays:  int(req.TreatmentDurationDays),
		RecoveryStatus:         req.RecoveryStatus,
		EffectivenessRating:    int(req.EffectivenessRating),
		FollowUpDate:           req.FollowUpDate.AsTime(),
		PreventionMeasures:     req.PreventionMeasures,
		EnvironmentalFactors:   req.EnvironmentalFactors,
		Images:                 req.Images,
		Notes:                  req.Notes,
		CreatedBy:              req.CreatedBy,
	}
	
	// Call usecase
	pestRecord, err := s.pestDiseaseRecordUsecase.Create(ctx, usecaseReq)
	if err != nil {
		return nil, s.handlePestDiseaseError(err)
	}
	
	// Convert entity to proto response
	protoRecord := s.convertToProtoPestDiseaseRecord(pestRecord)
	
	return &pb.CreatePestDiseaseRecordResponse{
		PestDiseaseRecord: protoRecord,
	}, nil
}

func (s *ProductionService) convertToProtoPestDiseaseRecord(record *entity.PestDiseaseRecord) *pb.PestDiseaseRecord {
	protoRecord := &pb.PestDiseaseRecord{
		Id:                     record.ID,
		PlantingCycleId:        record.PlantingCycleID,
		Type:                   record.Type,
		Name:                   record.Name,
		ScientificName:         record.ScientificName,
		Severity:               record.Severity,
		AffectedAreaPercentage: record.AffectedAreaPercentage,
		AffectedPlantCount:     int32(record.AffectedPlantCount),
		DetectionMethod:        record.DetectionMethod,
		Symptoms:               record.Symptoms,
		TreatmentApplied:       record.TreatmentApplied,
		TreatmentCost:          record.TreatmentCost,
		TreatmentDurationDays:  int32(record.TreatmentDurationDays),
		RecoveryStatus:         record.RecoveryStatus,
		EffectivenessRating:    int32(record.EffectivenessRating),
		PreventionMeasures:     record.PreventionMeasures,
		EnvironmentalFactors:   record.EnvironmentalFactors,
		Images:                 record.Images,
		Notes:                  record.Notes,
		CreatedBy:              record.CreatedBy,
		CreatedAt:              timestamppb.New(record.CreatedAt),
	}
	
	if record.DetectionDate != nil {
		protoRecord.DetectionDate = timestamppb.New(*record.DetectionDate)
	}
	if record.TreatmentDate != nil {
		protoRecord.TreatmentDate = timestamppb.New(*record.TreatmentDate)
	}
	if record.FollowUpDate != nil {
		protoRecord.FollowUpDate = timestamppb.New(*record.FollowUpDate)
	}
	if record.UpdatedAt != nil {
		protoRecord.UpdatedAt = timestamppb.New(*record.UpdatedAt)
	}
	
	return protoRecord
}

func (s *ProductionService) handlePestDiseaseError(err error) error {
	switch err {
	case pest_disease_record.ErrPestDiseaseRecordNotFound:
		return status.Error(codes.NotFound, "Pest disease record not found")
	case pest_disease_record.ErrInvalidRecordType:
		return status.Error(codes.InvalidArgument, "Invalid record type")
	case pest_disease_record.ErrInvalidSeverity:
		return status.Error(codes.InvalidArgument, "Invalid severity level")
	case pest_disease_record.ErrInvalidDetectionMethod:
		return status.Error(codes.InvalidArgument, "Invalid detection method")
	case pest_disease_record.ErrInvalidRecoveryStatus:
		return status.Error(codes.InvalidArgument, "Invalid recovery status")
	case pest_disease_record.ErrInvalidEffectivenessRating:
		return status.Error(codes.InvalidArgument, "Invalid effectiveness rating")
	default:
		return status.Error(codes.Internal, "Internal server error")
	}
}
```

### Server Initialization

```go
// file: infrastructure/grpc_service/server.go
package grpc_service

import (
	"log"
	"net"
	
	"google.golang.org/grpc"
	"monitoring_service/domain/usecase/harvest_record"
	"monitoring_service/domain/usecase/pest_disease_record"
	"monitoring_service/infrastructure/repo"
	pb "monitoring_service/proto_generated"
)

func StartGRPCServer(port string, db *pg.DB) {
	// Initialize repositories
	repos := repo.InitRepositories(db)
	
	// Initialize usecases
	harvestRecordUsecase := harvest_record.NewHarvestRecordUsecase(repos.HarvestRecordRepository)
	pestDiseaseRecordUsecase := pest_disease_record.NewPestDiseaseRecordUsecase(repos.PestDiseaseRecordRepository)
	
	// Create gRPC server
	grpcServer := grpc.NewServer()
	
	// Register services
	productionService := NewProductionService(harvestRecordUsecase, pestDiseaseRecordUsecase)
	pb.RegisterProductionServiceServer(grpcServer, productionService)
	
	// Start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	log.Printf("gRPC server listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

