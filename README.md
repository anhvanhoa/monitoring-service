# Monitoring Service

Microservice quản lý cảnh báo môi trường từ hệ thống IoT trong nông nghiệp, được xây dựng bằng Go và tuân theo nguyên tắc Clean Architecture. Service này giám sát các thông số môi trường như nhiệt độ, độ ẩm, pH, mức nước và trạng thái thiết bị để đảm bảo điều kiện tối ưu cho cây trồng.

## 🏗️ Kiến trúc

Dự án này tuân theo **Clean Architecture** với sự phân tách rõ ràng các mối quan tâm:

```
├── domain/           # Tầng logic nghiệp vụ
│   ├── entity/       # Các thực thể nghiệp vụ cốt lõi
│   │   └── environmental_alert.go # Entity cảnh báo môi trường
│   ├── repository/   # Giao diện truy cập dữ liệu
│   │   └── environmental_alert_repository.go
│   └── usecase/      # Các trường hợp sử dụng nghiệp vụ
│       └── environmental_alert/ # Use cases cảnh báo môi trường
├── infrastructure/   # Các mối quan tâm bên ngoài
│   ├── grpc_service/ # Triển khai API gRPC
│   │   ├── environmental_alert/ # gRPC handlers cảnh báo môi trường
│   │   └── server.go            # Thiết lập gRPC server
│   └── repo/         # Triển khai repository cơ sở dữ liệu
├── bootstrap/        # Khởi tạo ứng dụng
└── cmd/             # Điểm vào ứng dụng
```

## 🚀 Tính năng

### Quản lý Cảnh báo Môi trường
- ✅ Tạo, đọc, cập nhật, xóa cảnh báo môi trường
- ✅ Liệt kê cảnh báo với bộ lọc (thiết bị, loại cảnh báo, trạng thái, mức độ nghiêm trọng)
- ✅ Theo dõi các loại cảnh báo: nhiệt độ, độ ẩm, pH, mức nước, hỏng thiết bị
- ✅ Quản lý trạng thái cảnh báo (hoạt động, đã xác nhận, đã xử lý, leo thang, bỏ qua)
- ✅ Hệ thống ưu tiên và leo thang cảnh báo
- ✅ Tự động xử lý và hành động khi điều kiện được phục hồi
- ✅ Thông báo đa kênh (email, SMS, push notification, webhook)
- ✅ Đánh giá tác động và ghi chú xử lý
- ✅ Thống kê và báo cáo cảnh báo
- ✅ Xác thực dữ liệu đầu vào và quy tắc nghiệp vụ

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.25.0
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC
- **Kiến trúc**: Clean Architecture
- **Dependencies**:
  - `github.com/go-pg/pg/v10` - PostgreSQL ORM
  - `google.golang.org/grpc` - Framework gRPC
  - `github.com/spf13/viper` - Quản lý cấu hình
  - `go.uber.org/zap` - Logging có cấu trúc
  - `github.com/anhvanhoa/service-core` - Core service utilities
  - `github.com/anhvanhoa/sf-proto` - Protocol buffer definitions

## 📋 Yêu cầu hệ thống

- Go 1.25.0 trở lên
- PostgreSQL 12 trở lên
- [golang-migrate](https://github.com/golang-migrate/migrate) để quản lý migration cơ sở dữ liệu

## 🚀 Hướng dẫn nhanh

### 1. Clone repository
```bash
git clone <repository-url>
cd monitoring_service
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
url_db: "postgres://pg:123456@localhost:5432/monitoring_service_db?sslmode=disable"
name_service: "MonitoringService"
port_grpc: 50057
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

**10 cảnh báo môi trường với các trạng thái đa dạng:**

**Các loại cảnh báo:**
- **Nhiệt độ**: Cảnh báo nhiệt độ cao/thấp
- **Độ ẩm**: Cảnh báo độ ẩm cao/thấp  
- **pH**: Cảnh báo pH cao/thấp
- **Nước**: Cảnh báo thiếu nước
- **Thiết bị**: Cảnh báo hỏng thiết bị

**Trạng thái cảnh báo:**
- **Active**: Cảnh báo đang hoạt động
- **Acknowledged**: Đã được xác nhận
- **Resolved**: Đã được xử lý
- **Escalated**: Đã leo thang
- **Ignored**: Đã bỏ qua

**Mức độ nghiêm trọng:**
- **Info**: Thông tin
- **Warning**: Cảnh báo
- **Critical**: Nghiêm trọng
- **Emergency**: Khẩn cấp

**Dữ liệu bao gồm:**
- Giá trị hiện tại và ngưỡng cảnh báo
- Thời gian kích hoạt, xác nhận và xử lý
- Hành động tự động đã thực hiện
- Đánh giá tác động và ghi chú xử lý
- Thông tin thông báo đã gửi

## 📁 Cấu trúc Dự án

```
monitoring_service/
├── bootstrap/                 # Khởi tạo ứng dụng
│   ├── app.go               # Khởi tạo app
│   └── env.go               # Cấu hình môi trường
├── cmd/                     # Điểm vào ứng dụng
│   ├── main.go             # Điểm vào service chính
│   └── client/             # gRPC client để test
├── domain/                  # Logic nghiệp vụ (Clean Architecture)
│   ├── entity/             # Các thực thể nghiệp vụ cốt lõi
│   │   └── environmental_alert.go # Entity cảnh báo môi trường và DTOs
│   ├── repository/         # Giao diện truy cập dữ liệu
│   │   └── environmental_alert_repository.go
│   └── usecase/            # Các trường hợp sử dụng nghiệp vụ
│       └── environmental_alert/ # Use cases cảnh báo môi trường
│           ├── create_environmental_alert_usecase.go
│           ├── get_environmental_alert_usecase.go
│           ├── list_environmental_alert_usecase.go
│           ├── update_environmental_alert_usecase.go
│           ├── delete_environmental_alert_usecase.go
│           ├── get_environmental_alert_statistics_usecase.go
│           └── environmental_alert_usecase.go
├── infrastructure/          # Các mối quan tâm bên ngoài
│   ├── grpc_service/       # Triển khai API gRPC
│   │   ├── environmental_alert/ # gRPC handlers cảnh báo môi trường
│   │   │   ├── create.go
│   │   │   ├── get.go
│   │   │   ├── list.go
│   │   │   ├── update.go
│   │   │   ├── delete.go
│   │   │   └── statistics.go
│   │   └── server.go             # Thiết lập gRPC server
│   └── repo/               # Triển khai cơ sở dữ liệu
│       ├── environmental_alert_repository.go
│       └── init.go
├── migrations/              # Database migrations
│   ├── 000000_common.up.sql
│   ├── 000002_create_environmental_alerts.up.sql
│   └── seed/                     # Dữ liệu mẫu
│       └── 000006_seed_environmental_alerts.up.sql
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

### Cảnh báo Môi trường (Environmental Alert)
- **ID**: Định danh duy nhất
- **DeviceID**: ID thiết bị IoT phát hiện cảnh báo
- **AlertType**: Loại cảnh báo (temperature_high, temperature_low, humidity_high, humidity_low, ph_high, ph_low, water_shortage, equipment_failure)
- **CurrentValue**: Giá trị đo được tại thời điểm cảnh báo
- **ThresholdValue**: Ngưỡng quy định gây ra cảnh báo
- **ThresholdType**: Kiểu ngưỡng (min, max, range)
- **Severity**: Mức độ nghiêm trọng (info, warning, critical, emergency)
- **Priority**: Mức độ ưu tiên (1 cao nhất, 5 thấp nhất)
- **TriggeredAt**: Thời điểm kích hoạt cảnh báo
- **AcknowledgedAt**: Thời điểm xác nhận cảnh báo
- **AcknowledgedBy**: Người xác nhận cảnh báo
- **ResolvedAt**: Thời điểm xử lý cảnh báo
- **ResolvedBy**: Người xử lý cảnh báo
- **AutoResolve**: Tự động xử lý khi điều kiện an toàn
- **Status**: Trạng thái (active, acknowledged, resolved, escalated, ignored)
- **AutoActionTaken**: Hành động tự động đã thực hiện
- **EscalationLevel**: Cấp độ leo thang
- **NotificationSent**: Thông tin thông báo đã gửi (JSON)
- **ImpactAssessment**: Đánh giá tác động
- **ResolutionNotes**: Ghi chú xử lý
- **CreatedBy**: Người tạo cảnh báo
- **Timestamps**: Thời gian tạo/cập nhật

## 🔌 API Endpoints

Service cung cấp các endpoint gRPC:

### Environmental Alert Service
- `CreateEnvironmentalAlert` - Tạo cảnh báo môi trường mới
- `GetEnvironmentalAlert` - Lấy thông tin cảnh báo theo ID
- `UpdateEnvironmentalAlert` - Cập nhật thông tin cảnh báo
- `DeleteEnvironmentalAlert` - Xóa cảnh báo môi trường
- `ListEnvironmentalAlerts` - Liệt kê cảnh báo với bộ lọc
- `GetEnvironmentalAlertStatistics` - Lấy thống kê cảnh báo
- `GetActiveAlerts` - Lấy danh sách cảnh báo đang hoạt động
- `GetByDevice` - Lấy cảnh báo theo thiết bị
- `GetByAlertType` - Lấy cảnh báo theo loại
- `GetByStatus` - Lấy cảnh báo theo trạng thái
- `GetBySeverity` - Lấy cảnh báo theo mức độ nghiêm trọng
- `GetByPriority` - Lấy cảnh báo theo mức độ ưu tiên
- `GetByDateRange` - Lấy cảnh báo theo khoảng thời gian
- `GetOverdueAlerts` - Lấy cảnh báo quá hạn
- `GetEscalatedAlerts` - Lấy cảnh báo đã leo thang
- `AcknowledgeAlert` - Xác nhận cảnh báo
- `ResolveAlert` - Xử lý cảnh báo
- `EscalateAlert` - Leo thang cảnh báo
- `IgnoreAlert` - Bỏ qua cảnh báo

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

**Lưu ý**: Service này được thiết kế để quản lý cảnh báo môi trường từ hệ thống IoT trong nông nghiệp, tuân theo các nguyên tắc kiến trúc microservice để có thể mở rộng và bảo trì dễ dàng. Service giúp giám sát và phản ứng kịp thời với các thay đổi môi trường để đảm bảo điều kiện tối ưu cho cây trồng.
