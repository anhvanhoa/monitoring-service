# Domain Layer - Production Service

## Cấu trúc thư mục

```
domain/
├── entity/           # Các thực thể (entities)
│   ├── harvest_record.go # Entity và request/response cho HarvestRecord
│   └── pest_disease_record.go # Entity và request/response cho PestDiseaseRecord
├── repository/       # Interface repository
│   ├── harvest_record_repository.go
│   └── pest_disease_record_repository.go
└── usecase/          # Interface use case
    ├── harvest_record/   # Use cases cho HarvestRecord
    │   ├── create_harvest_record_usecase.go
    │   ├── get_harvest_record_usecase.go
    │   ├── update_harvest_record_usecase.go
    │   ├── delete_harvest_record_usecase.go
    │   ├── list_harvest_record_usecase.go
    │   ├── get_harvest_records_by_planting_cycle_usecase.go
    │   └── harvest_record_usecase.go
    └── pest_disease_record/ # Use cases cho PestDiseaseRecord
        ├── create_pest_disease_record_usecase.go
        ├── get_pest_disease_record_usecase.go
        ├── update_pest_disease_record_usecase.go
        ├── delete_pest_disease_record_usecase.go
        ├── list_pest_disease_record_usecase.go
        ├── get_pest_disease_records_by_planting_cycle_usecase.go
        └── pest_disease_record_usecase.go
```

## Mô tả các thành phần

### 2. Entity
#### HarvestRecord
- **HarvestRecord**: Entity chính cho bản ghi thu hoạch
- **FilterHarvestRecord**: Filter tìm kiếm bản ghi thu hoạch
- **QualityGrade**: Enum cho chất lượng sản phẩm (A+, A, B, C, Reject)
- **SizeClassification**: Enum cho phân loại kích thước (XL, L, M, S, XS)

#### PestDiseaseRecord
- **PestDiseaseRecord**: Entity chính cho bản ghi sâu bệnh
- **FilterPestDiseaseRecord**: Filter tìm kiếm bản ghi sâu bệnh
- **RecordType**: Enum cho loại bản ghi (pest, disease, nutrient_deficiency, environmental_stress)
- **Severity**: Enum cho mức độ nghiêm trọng (low, medium, high, critical)
- **DetectionMethod**: Enum cho phương pháp phát hiện (visual, trap, sensor, lab_test)
- **RecoveryStatus**: Enum cho trạng thái phục hồi (treating, recovering, recovered, failed, spreading)

### 3. Repository Interface
#### HarvestRecordRepository
- `Create()`: Tạo bản ghi thu hoạch mới
- `GetByID()`: Lấy bản ghi thu hoạch theo ID
- `Update()`: Cập nhật bản ghi thu hoạch
- `Delete()`: Xóa bản ghi thu hoạch
- `List()`: Lấy danh sách với filter và phân trang
- `GetByPlantingCycleID()`: Lấy theo chu kỳ trồng
- `Count()`: Đếm số lượng
- `CountByPlantingCycleID()`: Đếm số lượng theo chu kỳ trồng

#### PestDiseaseRecordRepository
- `Create()`: Tạo bản ghi sâu bệnh mới
- `GetByID()`: Lấy bản ghi sâu bệnh theo ID
- `Update()`: Cập nhật bản ghi sâu bệnh
- `Delete()`: Xóa bản ghi sâu bệnh
- `List()`: Lấy danh sách với filter và phân trang
- `GetByPlantingCycleID()`: Lấy theo chu kỳ trồng
- `Count()`: Đếm số lượng
- `CountByPlantingCycleID()`: Đếm số lượng theo chu kỳ trồng

### 4. Use Case
#### HarvestRecord Use Cases
- **CreateHarvestRecordUsecase**: Tạo bản ghi thu hoạch mới
- **GetHarvestRecordUsecase**: Lấy thông tin bản ghi thu hoạch
- **UpdateHarvestRecordUsecase**: Cập nhật bản ghi thu hoạch
- **DeleteHarvestRecordUsecase**: Xóa bản ghi thu hoạch
- **ListHarvestRecordUsecase**: Lấy danh sách bản ghi thu hoạch
- **GetHarvestRecordsByPlantingCycleUsecase**: Lấy tất cả bản ghi thu hoạch của một chu kỳ trồng

#### PestDiseaseRecord Use Cases
- **CreatePestDiseaseRecordUsecase**: Tạo bản ghi sâu bệnh mới
- **GetPestDiseaseRecordUsecase**: Lấy thông tin bản ghi sâu bệnh
- **UpdatePestDiseaseRecordUsecase**: Cập nhật bản ghi sâu bệnh
- **DeletePestDiseaseRecordUsecase**: Xóa bản ghi sâu bệnh
- **ListPestDiseaseRecordUsecase**: Lấy danh sách bản ghi sâu bệnh
- **GetPestDiseaseRecordsByPlantingCycleUsecase**: Lấy tất cả bản ghi sâu bệnh của một chu kỳ trồng

## Cách sử dụng

### 1. Tạo Use Case
```go
// Tạo repository (implement từ infrastructure layer)
harvestRecordRepo := infrastructure.NewHarvestRecordRepository(db)
pestDiseaseRecordRepo := infrastructure.NewPestDiseaseRecordRepository(db)

// Tạo use case
harvestRecordUsecase := harvest_record.NewHarvestRecordUsecase(harvestRecordRepo)
pestDiseaseRecordUsecase := pest_disease_record.NewPestDiseaseRecordUsecase(pestDiseaseRecordRepo)
```

### 2. Sử dụng Use Case
```go
// Tạo request cho HarvestRecord
harvestReq := &harvest_record.CreateHarvestRecordRequest{
    PlantingCycleID:    "cycle123",
    QuantityKg:         100.5,
    QualityGrade:       "A",
    SizeClassification: "L",
    MarketPricePerKg:   50000,
    PlantHealthRating: 4,
    CreatedBy:          "user123",
}

// Thực thi use case
harvestRecord, err := harvestRecordUsecase.Create(ctx, harvestReq)

// Tạo request cho PestDiseaseRecord
pestReq := &pest_disease_record.CreatePestDiseaseRecordRequest{
    PlantingCycleID:        "cycle123",
    Type:                   "pest",
    Name:                   "Aphids",
    Severity:               "medium",
    AffectedAreaPercentage: 15.5,
    DetectionMethod:        "visual",
    CreatedBy:              "user123",
}

// Thực thi use case
pestRecord, err := pestDiseaseRecordUsecase.Create(ctx, pestReq)
```

## 5. Repository Implementation

### Infrastructure Layer
Các repository interface được implement trong `infrastructure/repo/`:

#### HarvestRecordRepository Implementation
- **File**: `infrastructure/repo/harvest_record_repository.go`
- **Database**: PostgreSQL với go-pg/pg/v10 ORM
- **Features**: CRUD operations, filtering, pagination, counting

#### PestDiseaseRecordRepository Implementation  
- **File**: `infrastructure/repo/pest_disease_record_repository.go`
- **Database**: PostgreSQL với go-pg/pg/v10 ORM
- **Features**: CRUD operations, filtering, pagination, counting

#### Repository Factory
- **File**: `infrastructure/repo/repository_factory.go`
- **Purpose**: Tạo repository instances một cách tập trung
- **Usage**: `factory.NewHarvestRecordRepository()`, `factory.NewPestDiseaseRecordRepository()`

#### Repository Initialization
- **File**: `infrastructure/repo/init.go`
- **Purpose**: Khởi tạo tất cả repositories
- **Usage**: `repos := repo.InitRepositories(db)`

### Cách sử dụng Repository Implementation

```go
import (
    "monitoring_service/infrastructure/repo"
    "github.com/go-pg/pg/v10"
)

// Khởi tạo database connection
db := pg.Connect(&pg.Options{
    Addr:     "localhost:5432",
    User:     "postgres", 
    Password: "password",
    Database: "monitoring_service",
})

// Khởi tạo repositories
repos := repo.InitRepositories(db)

// Sử dụng với usecases
harvestRecordUsecase := harvest_record.NewHarvestRecordUsecase(repos.HarvestRecordRepository)
pestDiseaseRecordUsecase := pest_disease_record.NewPestDiseaseRecordUsecase(repos.PestDiseaseRecordRepository)
```

## Lưu ý
- Tất cả use case đều nhận `context.Context` làm tham số đầu tiên
- Các repository interface được implement trong infrastructure layer với PostgreSQL
- Các use case interface được implement trong infrastructure layer (gRPC service)
- Sử dụng Clean Architecture pattern để tách biệt business logic và infrastructure
- Repository implementations hỗ trợ filtering, pagination và counting
- Sử dụng go-pg/pg/v10 ORM cho database operations