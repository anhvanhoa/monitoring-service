# Crop Service

Microservice quản lý giống cây trồng và chu kỳ trồng trọt trong hệ thống nông nghiệp, được xây dựng bằng Go và tuân theo nguyên tắc Clean Architecture.

## 🏗️ Kiến trúc

Dự án này tuân theo **Clean Architecture** với sự phân tách rõ ràng các mối quan tâm:

```
├── domain/           # Tầng logic nghiệp vụ
│   ├── entity/       # Các thực thể nghiệp vụ cốt lõi
│   │   ├── plant_variety.go      # Entity giống cây trồng
│   │   └── planting_cycle.go     # Entity chu kỳ trồng
│   ├── repository/   # Giao diện truy cập dữ liệu
│   │   ├── plant_variety_repository.go
│   │   └── planting_cycle_repository.go
│   └── usecase/      # Các trường hợp sử dụng nghiệp vụ
│       ├── plant_variety/        # Use cases giống cây trồng
│       └── planting_cycle/       # Use cases chu kỳ trồng
├── infrastructure/   # Các mối quan tâm bên ngoài
│   ├── grpc_service/ # Triển khai API gRPC
│   │   ├── plant_variety/        # gRPC handlers giống cây trồng
│   │   └── planting_cycle/       # gRPC handlers chu kỳ trồng
│   └── repo/         # Triển khai repository cơ sở dữ liệu
├── bootstrap/        # Khởi tạo ứng dụng
└── cmd/             # Điểm vào ứng dụng
```

## 🚀 Tính năng

### Quản lý Giống cây trồng
- ✅ Tạo, đọc, cập nhật, xóa giống cây trồng
- ✅ Liệt kê giống cây với bộ lọc (loại, mùa vụ, trạng thái)
- ✅ Tìm kiếm giống cây theo điều kiện môi trường (nhiệt độ, độ ẩm, pH)
- ✅ Lọc theo yêu cầu ánh sáng và nước
- ✅ Lọc theo mùa vụ và loại cây
- ✅ Hỗ trợ phân trang và sắp xếp
- ✅ Xác thực dữ liệu đầu vào

### Quản lý Chu kỳ trồng
- ✅ Tạo, đọc, cập nhật, xóa chu kỳ trồng
- ✅ Liệt kê chu kỳ với bộ lọc (khu vực, giống cây, trạng thái, ngày tháng)
- ✅ Theo dõi tiến độ chu kỳ trồng (lập kế hoạch → gieo hạt → cấy ghép → phát triển → thu hoạch)
- ✅ Quản lý lịch gieo hạt và thu hoạch
- ✅ Báo cáo chu kỳ sắp thu hoạch và quá hạn
- ✅ Lấy chu kỳ theo giống cây và khu vực
- ✅ Cập nhật trạng thái và ngày thu hoạch

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.24.6
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC
- **Kiến trúc**: Clean Architecture
- **Dependencies**:
  - `github.com/go-pg/pg/v10` - PostgreSQL ORM
  - `google.golang.org/grpc` - Framework gRPC
  - `github.com/spf13/viper` - Quản lý cấu hình
  - `go.uber.org/zap` - Logging có cấu trúc

## 📋 Yêu cầu hệ thống

- Go 1.24.6 trở lên
- PostgreSQL 12 trở lên
- [golang-migrate](https://github.com/golang-migrate/migrate) để quản lý migration cơ sở dữ liệu

## 🚀 Hướng dẫn nhanh

### 1. Clone repository
```bash
git clone <repository-url>
cd farm-service
```

### 2. Cài đặt dependencies
```bash
go mod download
```

### 3. Thiết lập cơ sở dữ liệu
```bash
# Tạo cơ sở dữ liệu
make create-db

# Chạy migrations
make up
```

### 4. Cấu hình ứng dụng
Sao chép và chỉnh sửa file cấu hình:
```bash
cp dev.config.yml config.yml
```

Cập nhật chuỗi kết nối cơ sở dữ liệu trong `config.yml`:
```yaml
node_env: "development"
url_db: "postgres://postgres:123456@localhost:5432/crop_service_db?sslmode=disable"
name_service: "CropService"
port_grpc: 50054
host_grpc: "localhost"
interval_check: "20s"
timeout_check: "15s"
```

### 5. Chạy ứng dụng
```bash
# Build và chạy service chính
make run

# Hoặc chạy client để test
make client
```

## 🗄️ Quản lý Cơ sở dữ liệu

Dự án sử dụng `golang-migrate` để quản lý schema cơ sở dữ liệu:

```bash
# Chạy tất cả migrations đang chờ
make up

# Rollback migration cuối cùng
make down

# Reset cơ sở dữ liệu hoàn toàn
make reset

# Tạo migration mới
make create name=migration_name

# Force migration đến phiên bản cụ thể
make force version=1
```

## 🌱 Dữ liệu mẫu

Dự án bao gồm dữ liệu mẫu để phát triển và kiểm thử:

```bash
# Chèn dữ liệu mẫu vào cơ sở dữ liệu
make seed

# Reset cơ sở dữ liệu và chèn dữ liệu mẫu
make seed-reset

# Chèn dữ liệu mẫu vào cơ sở dữ liệu Docker
make docker-seed
```

### Dữ liệu mẫu bao gồm:

**15 giống cây trồng với thông tin chi tiết:**
- **Rau cải**: Cải bắp, Cải ngọt, Cải xoong, Cải xoăn
- **Rau củ**: Cà rốt, Khoai tây, Củ cải trắng  
- **Rau quả**: Cà chua, Ớt chuông, Dưa chuột
- **Rau thơm**: Rau mùi, Húng quế, Bạc hà
- **Rau lá xanh**: Rau muống, Rau dền

Mỗi giống cây bao gồm:
- Thông tin cơ bản (tên, tên khoa học, loại, mùa vụ)
- Điều kiện môi trường tối ưu (nhiệt độ, độ ẩm, pH)
- Yêu cầu chăm sóc (nước, ánh sáng)
- Thời gian phát triển và mô tả chi tiết

**15 chu kỳ trồng với trạng thái đa dạng:**
- Các chu kỳ với trạng thái khác nhau (lập kế hoạch, gieo hạt, cấy ghép, phát triển, ra hoa, thu hoạch, hoàn thành, thất bại)
- Dữ liệu thực tế về ngày gieo hạt, cấy ghép, thu hoạch dự kiến và thực tế
- Thông tin về số lượng cây, lô hạt giống và ghi chú
- Liên kết với giống cây trồng và khu vực trồng

## 📁 Cấu trúc Dự án

```
crop-service/
├── bootstrap/                 # Khởi tạo ứng dụng
│   ├── app.go               # Khởi tạo app
│   └── env.go               # Cấu hình môi trường
├── cmd/                     # Điểm vào ứng dụng
│   ├── main.go             # Điểm vào service chính
│   └── client/             # gRPC client để test
├── domain/                  # Logic nghiệp vụ (Clean Architecture)
│   ├── entity/             # Các thực thể nghiệp vụ cốt lõi
│   │   ├── plant_variety.go      # Entity giống cây trồng và DTOs
│   │   └── planting_cycle.go     # Entity chu kỳ trồng và DTOs
│   ├── repository/         # Giao diện truy cập dữ liệu
│   │   ├── plant_variety_repository.go
│   │   └── planting_cycle_repository.go
│   └── usecase/            # Các trường hợp sử dụng nghiệp vụ
│       ├── plant_variety/        # Use cases giống cây trồng
│       │   ├── create_plant_variety_usecase.go
│       │   ├── get_plant_variety_usecase.go
│       │   ├── list_plant_variety_usecase.go
│       │   ├── search_plant_varieties_usecase.go
│       │   └── ... (các use case khác)
│       └── planting_cycle/       # Use cases chu kỳ trồng
│           ├── create_planting_cycle_usecase.go
│           ├── get_planting_cycle_usecase.go
│           ├── list_planting_cycle_usecase.go
│           └── ... (các use case khác)
├── infrastructure/          # Các mối quan tâm bên ngoài
│   ├── grpc_service/       # Triển khai API gRPC
│   │   ├── plant_variety/        # gRPC handlers giống cây trồng
│   │   ├── planting_cycle/       # gRPC handlers chu kỳ trồng
│   │   └── server.go             # Thiết lập gRPC server
│   └── repo/               # Triển khai cơ sở dữ liệu
│       ├── plant_variety_repository.go
│       ├── planting_cycle_repository.go
│       └── base.go
├── migrations/              # Database migrations
│   ├── 000000_common.up.sql
│   ├── 000002_create_plant_varieties_table.up.sql
│   ├── 000003_create_planting_cycles_table.up.sql
│   └── seed/                     # Dữ liệu mẫu
├── script/seed/             # Script chèn dữ liệu mẫu
├── doc/                     # Tài liệu
└── logs/                    # Log ứng dụng
```

## 🔧 Các lệnh có sẵn

```bash
# Thao tác cơ sở dữ liệu
make up              # Chạy migrations
make down            # Rollback migration
make reset           # Reset cơ sở dữ liệu
make create-db       # Tạo cơ sở dữ liệu
make drop-db         # Xóa cơ sở dữ liệu

# Ứng dụng
make build           # Build ứng dụng
make run             # Chạy service chính
make client          # Chạy client test
make test            # Chạy tests

# Trợ giúp
make help            # Hiển thị tất cả lệnh có sẵn
```

## 📊 Mô hình Dữ liệu

### Giống cây trồng (Plant Variety)
- **ID**: Định danh duy nhất
- **Name**: Tên giống cây trồng
- **ScientificName**: Tên khoa học
- **Category**: Loại cây (rau cải, rau củ, rau quả, rau thơm, rau lá xanh)
- **GrowingSeason**: Mùa vụ phù hợp
- **GrowthDurationDays**: Thời gian phát triển (ngày)
- **OptimalTempMin/Max**: Nhiệt độ tối ưu (min/max)
- **OptimalHumidityMin/Max**: Độ ẩm tối ưu (min/max)
- **PHMin/Max**: Độ pH tối ưu (min/max)
- **WaterRequirement**: Yêu cầu nước (thấp, trung bình, cao)
- **LightRequirement**: Yêu cầu ánh sáng (ít, trung bình, nhiều)
- **Description**: Mô tả chi tiết
- **MediaID**: ID phương tiện truyền thông
- **Status**: Trạng thái (active, inactive)
- **CreatedBy**: Định danh người tạo
- **Timestamps**: Thời gian tạo/cập nhật

### Chu kỳ trồng (Planting Cycle)
- **ID**: Định danh duy nhất
- **CycleName**: Tên chu kỳ trồng
- **GrowingZoneID**: ID khu vực trồng
- **PlantVarietyID**: ID giống cây trồng
- **SeedDate**: Ngày gieo hạt
- **TransplantDate**: Ngày cấy ghép
- **ExpectedHarvestDate**: Ngày thu hoạch dự kiến
- **ActualHarvestDate**: Ngày thu hoạch thực tế
- **PlantQuantity**: Số lượng cây
- **SeedBatch**: Lô hạt giống
- **Status**: Trạng thái (planning, seeding, transplanting, growing, flowering, harvesting, completed, failed)
- **Notes**: Ghi chú
- **CreatedBy**: Định danh người tạo
- **Timestamps**: Thời gian tạo/cập nhật

## 🔌 API Endpoints

Service cung cấp các endpoint gRPC:

### Plant Variety Service
- `CreatePlantVariety` - Tạo giống cây trồng mới
- `GetPlantVariety` - Lấy thông tin giống cây trồng theo ID
- `UpdatePlantVariety` - Cập nhật thông tin giống cây trồng
- `DeletePlantVariety` - Xóa giống cây trồng
- `ListPlantVarieties` - Liệt kê giống cây trồng với bộ lọc
- `SearchPlantVarieties` - Tìm kiếm giống cây trồng
- `GetActivePlantVarieties` - Lấy danh sách giống cây trồng đang hoạt động
- `GetByCategory` - Lấy giống cây theo loại
- `GetBySeason` - Lấy giống cây theo mùa vụ
- `GetByStatus` - Lấy giống cây theo trạng thái
- `GetByTemperatureRange` - Lấy giống cây theo khoảng nhiệt độ
- `GetByHumidityRange` - Lấy giống cây theo khoảng độ ẩm
- `GetByWaterRequirement` - Lấy giống cây theo yêu cầu nước
- `GetByLightRequirement` - Lấy giống cây theo yêu cầu ánh sáng

### Planting Cycle Service
- `CreatePlantingCycle` - Tạo chu kỳ trồng mới
- `GetPlantingCycle` - Lấy thông tin chu kỳ trồng theo ID
- `UpdatePlantingCycle` - Cập nhật thông tin chu kỳ trồng
- `DeletePlantingCycle` - Xóa chu kỳ trồng
- `ListPlantingCycles` - Liệt kê chu kỳ trồng với bộ lọc
- `GetActivePlantingCycles` - Lấy danh sách chu kỳ trồng đang hoạt động
- `GetByVariety` - Lấy chu kỳ trồng theo giống cây
- `GetByZone` - Lấy chu kỳ trồng theo khu vực
- `GetByStatus` - Lấy chu kỳ trồng theo trạng thái
- `GetByDateRange` - Lấy chu kỳ trồng theo khoảng ngày
- `GetBySeedDateRange` - Lấy chu kỳ trồng theo khoảng ngày gieo hạt
- `GetByHarvestDateRange` - Lấy chu kỳ trồng theo khoảng ngày thu hoạch
- `GetUpcomingHarvests` - Lấy chu kỳ sắp thu hoạch
- `GetOverdueHarvests` - Lấy chu kỳ thu hoạch quá hạn
- `GetCycleWithDetails` - Lấy chu kỳ trồng với thông tin chi tiết
- `GetCyclesWithDetails` - Lấy danh sách chu kỳ trồng với thông tin chi tiết
- `UpdateStatus` - Cập nhật trạng thái chu kỳ trồng
- `UpdateHarvestDate` - Cập nhật ngày thu hoạch

## 🧪 Testing

Chạy client test để tương tác với service:

```bash
make client
```

Điều này sẽ khởi động một client tương tác nơi bạn có thể test tất cả các endpoint gRPC.

## 📝 Cấu hình

Ứng dụng sử dụng Viper để quản lý cấu hình. Các tùy chọn cấu hình chính:

- `node_env`: Môi trường (development, production)
- `url_db`: Chuỗi kết nối PostgreSQL
- `name_service`: Tên service cho discovery
- `port_grpc`: Cổng gRPC server
- `host_grpc`: Host gRPC server
- `interval_check`: Khoảng thời gian kiểm tra sức khỏe
- `timeout_check`: Timeout kiểm tra sức khỏe

## 🚀 Triển khai

1. **Build ứng dụng**:
   ```bash
   make build
   ```

2. **Thiết lập cơ sở dữ liệu production**:
   ```bash
   make create-db
   make up
   ```

3. **Chạy service**:
   ```bash
   ./bin/app
   ```

## 🤝 Đóng góp

1. Fork repository
2. Tạo feature branch
3. Thực hiện thay đổi
4. Thêm tests nếu cần thiết
5. Submit pull request

## 📄 Giấy phép

Dự án này được cấp phép theo MIT License.

## 🆘 Hỗ trợ

Để được hỗ trợ và đặt câu hỏi, vui lòng tạo issue trong repository.

---

**Lưu ý**: Service này được thiết kế để quản lý giống cây trồng và chu kỳ trồng trọt trong hệ thống nông nghiệp, tuân theo các nguyên tắc kiến trúc microservice để có thể mở rộng và bảo trì dễ dàng.
