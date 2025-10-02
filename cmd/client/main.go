package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Nhập trang (mặc định 1): ")
	offsetStr, _ := reader.ReadString('\n')
	offsetStr = cleanInput(offsetStr)
	offset := int32(1)
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(o)
		}
	}

	fmt.Print("Nhập số bản ghi mỗi trang (mặc định 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	return offset, limit
}

type EnvironmentalAlertClient struct {
	environmentalAlertClient proto_environmental_alert.EnvironmentalAlertServiceClient
	conn                     *grpc.ClientConn
}

func NewEnvironmentalAlertClient(address string) (*EnvironmentalAlertClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &EnvironmentalAlertClient{
		environmentalAlertClient: proto_environmental_alert.NewEnvironmentalAlertServiceClient(conn),
		conn:                     conn,
	}, nil
}

func (c *EnvironmentalAlertClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Environmental Alert Service Tests ==================

func (c *EnvironmentalAlertClient) TestCreateEnvironmentalAlert() {
	fmt.Println("\n=== Kiểm thử Tạo Cảnh báo Môi trường ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị: ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập loại cảnh báo (temperature_high/temperature_low/humidity_high/humidity_low/ph_high/ph_low/water_shortage/equipment_failure): ")
	alertType, _ := reader.ReadString('\n')
	alertType = cleanInput(alertType)

	fmt.Print("Nhập giá trị hiện tại: ")
	currentValueStr, _ := reader.ReadString('\n')
	currentValueStr = cleanInput(currentValueStr)
	currentValue := float64(25.0)
	if currentValueStr != "" {
		if v, err := strconv.ParseFloat(currentValueStr, 64); err == nil {
			currentValue = v
		}
	}

	fmt.Print("Nhập giá trị ngưỡng: ")
	thresholdValueStr, _ := reader.ReadString('\n')
	thresholdValueStr = cleanInput(thresholdValueStr)
	thresholdValue := float64(30.0)
	if thresholdValueStr != "" {
		if v, err := strconv.ParseFloat(thresholdValueStr, 64); err == nil {
			thresholdValue = v
		}
	}

	fmt.Print("Nhập kiểu ngưỡng (min/max/range): ")
	thresholdType, _ := reader.ReadString('\n')
	thresholdType = cleanInput(thresholdType)

	fmt.Print("Nhập mức độ nghiêm trọng (info/warning/critical/emergency): ")
	severity, _ := reader.ReadString('\n')
	severity = cleanInput(severity)

	fmt.Print("Nhập mức độ ưu tiên (1-10): ")
	priorityStr, _ := reader.ReadString('\n')
	priorityStr = cleanInput(priorityStr)
	priority := int32(5)
	if priorityStr != "" {
		if p, err := strconv.Atoi(priorityStr); err == nil {
			priority = int32(p)
		}
	}

	fmt.Print("Tự động xử lý (true/false): ")
	autoResolveStr, _ := reader.ReadString('\n')
	autoResolveStr = cleanInput(autoResolveStr)
	autoResolve := false
	if autoResolveStr == "true" {
		autoResolve = true
	}

	fmt.Print("Nhập hành động tự động: ")
	autoActionTaken, _ := reader.ReadString('\n')
	autoActionTaken = cleanInput(autoActionTaken)

	fmt.Print("Nhập cấp độ leo thang (1-5): ")
	escalationLevelStr, _ := reader.ReadString('\n')
	escalationLevelStr = cleanInput(escalationLevelStr)
	escalationLevel := int32(1)
	if escalationLevelStr != "" {
		if e, err := strconv.Atoi(escalationLevelStr); err == nil {
			escalationLevel = int32(e)
		}
	}

	fmt.Print("Nhập đánh giá tác động: ")
	impactAssessment, _ := reader.ReadString('\n')
	impactAssessment = cleanInput(impactAssessment)

	fmt.Print("Nhập người tạo: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.environmentalAlertClient.CreateEnvironmentalAlert(ctx, &proto_environmental_alert.CreateEnvironmentalAlertRequest{
		DeviceId:         deviceId,
		AlertType:        alertType,
		CurrentValue:     currentValue,
		ThresholdValue:   thresholdValue,
		ThresholdType:    thresholdType,
		Severity:         severity,
		Priority:         priority,
		AutoResolve:      autoResolve,
		AutoActionTaken:  autoActionTaken,
		EscalationLevel:  escalationLevel,
		ImpactAssessment: impactAssessment,
		CreatedBy:        createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateEnvironmentalAlert: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo cảnh báo môi trường:\n")
	if resp.Alert != nil {
		fmt.Printf("ID: %s\n", resp.Alert.Id)
		fmt.Printf("Device ID: %s\n", resp.Alert.DeviceId)
		fmt.Printf("Alert Type: %s\n", resp.Alert.AlertType)
		fmt.Printf("Current Value: %.2f\n", resp.Alert.CurrentValue)
		fmt.Printf("Threshold Value: %.2f\n", resp.Alert.ThresholdValue)
		fmt.Printf("Severity: %s\n", resp.Alert.Severity)
		fmt.Printf("Priority: %d\n", resp.Alert.Priority)
		fmt.Printf("Status: %s\n", resp.Alert.Status)
	}
}

func (c *EnvironmentalAlertClient) TestGetEnvironmentalAlert() {
	fmt.Println("\n=== Kiểm thử Lấy Cảnh báo Môi trường ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID cảnh báo môi trường: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.environmentalAlertClient.GetEnvironmentalAlert(ctx, &proto_environmental_alert.GetEnvironmentalAlertRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetEnvironmentalAlert: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy cảnh báo môi trường:\n")
	if resp.Alert != nil {
		fmt.Printf("ID: %s\n", resp.Alert.Id)
		fmt.Printf("Device ID: %s\n", resp.Alert.DeviceId)
		fmt.Printf("Loại cảnh báo: %s\n", resp.Alert.AlertType)
		fmt.Printf("Giá trị hiện tại: %.2f\n", resp.Alert.CurrentValue)
		fmt.Printf("Giá trị ngưỡng: %.2f\n", resp.Alert.ThresholdValue)
		fmt.Printf("Kiểu ngưỡng: %s\n", resp.Alert.ThresholdType)
		fmt.Printf("Mức độ nghiêm trọng: %s\n", resp.Alert.Severity)
		fmt.Printf("Mức độ ưu tiên: %d\n", resp.Alert.Priority)
		fmt.Printf("Trạng thái: %s\n", resp.Alert.Status)
		fmt.Printf("Tự động xử lý: %t\n", resp.Alert.AutoResolve)
		fmt.Printf("Hành động tự động: %s\n", resp.Alert.AutoActionTaken)
		fmt.Printf("Cấp độ leo thang: %d\n", resp.Alert.EscalationLevel)
		fmt.Printf("Đánh giá tác động: %s\n", resp.Alert.ImpactAssessment)
		fmt.Printf("Người tạo: %s\n", resp.Alert.CreatedBy)
	}
}

func (c *EnvironmentalAlertClient) TestListEnvironmentalAlerts() {
	fmt.Println("\n=== Kiểm thử Liệt kê Cảnh báo Môi trường ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	// Input filter options
	fmt.Println("\n--- Tùy chọn bộ lọc (để trống để bỏ qua) ---")

	fmt.Print("Nhập Device ID (để trống để bỏ qua): ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập loại cảnh báo (temperature_high/temperature_low/humidity_high/humidity_low/ph_high/ph_low/water_shortage/equipment_failure): ")
	alertType, _ := reader.ReadString('\n')
	alertType = cleanInput(alertType)

	fmt.Print("Nhập trạng thái (active/acknowledged/resolved/escalated/ignored): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	fmt.Print("Nhập mức độ nghiêm trọng (info/warning/critical/emergency): ")
	severity, _ := reader.ReadString('\n')
	severity = cleanInput(severity)

	fmt.Print("Nhập mức độ ưu tiên (1-10, để trống để bỏ qua): ")
	priorityStr, _ := reader.ReadString('\n')
	priorityStr = cleanInput(priorityStr)
	var priority int32
	if priorityStr != "" {
		if p, err := strconv.Atoi(priorityStr); err == nil {
			priority = int32(p)
		}
	}

	fmt.Print("Nhập cấp độ leo thang (0-5, để trống để bỏ qua): ")
	escalationLevelStr, _ := reader.ReadString('\n')
	escalationLevelStr = cleanInput(escalationLevelStr)
	var escalationLevel int32
	if escalationLevelStr != "" {
		if e, err := strconv.Atoi(escalationLevelStr); err == nil {
			escalationLevel = int32(e)
		}
	}

	fmt.Print("Cảnh báo quá hạn (true/false, để trống để bỏ qua): ")
	isOverdueStr, _ := reader.ReadString('\n')
	isOverdueStr = cleanInput(isOverdueStr)
	var isOverdue *bool
	if isOverdueStr != "" {
		overdue := isOverdueStr == "true"
		isOverdue = &overdue
	}

	fmt.Print("Tự động xử lý (true/false, để trống để bỏ qua): ")
	autoResolveStr, _ := reader.ReadString('\n')
	autoResolveStr = cleanInput(autoResolveStr)
	var autoResolve *bool
	if autoResolveStr != "" {
		resolve := autoResolveStr == "true"
		autoResolve = &resolve
	}

	fmt.Print("Nhập người tạo (để trống để bỏ qua): ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	fmt.Print("Nhập người xác nhận (để trống để bỏ qua): ")
	acknowledgedBy, _ := reader.ReadString('\n')
	acknowledgedBy = cleanInput(acknowledgedBy)

	fmt.Print("Nhập người xử lý (để trống để bỏ qua): ")
	resolvedBy, _ := reader.ReadString('\n')
	resolvedBy = cleanInput(resolvedBy)

	filters := &proto_environmental_alert.AlertFilters{}

	if deviceId != "" {
		filters.DeviceId = deviceId
	}
	if alertType != "" {
		filters.AlertType = alertType
	}
	if status != "" {
		filters.Status = status
	}
	if severity != "" {
		filters.Severity = severity
	}
	if priority > 0 {
		filters.Priority = priority
	}
	if escalationLevel > 0 {
		filters.EscalationLevel = escalationLevel
	}
	if isOverdue != nil {
		filters.IsOverdue = *isOverdue
	}
	if autoResolve != nil {
		filters.AutoResolve = *autoResolve
	}
	if createdBy != "" {
		filters.CreatedBy = createdBy
	}
	if acknowledgedBy != "" {
		filters.AcknowledgedBy = acknowledgedBy
	}
	if resolvedBy != "" {
		filters.ResolvedBy = resolvedBy
	}

	// Display current filters
	fmt.Println("\n--- Bộ lọc hiện tại ---")
	fmt.Printf("Filters: %+v\n", filters)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.environmentalAlertClient.ListEnvironmentalAlerts(ctx, &proto_environmental_alert.ListEnvironmentalAlertsRequest{
		Pagination: &proto_common.PaginationRequest{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "triggered_at",
			SortOrder: "desc",
		},
		Filters: filters,
	})
	if err != nil {
		fmt.Printf("Error calling ListEnvironmentalAlerts: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê cảnh báo môi trường:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách cảnh báo môi trường:\n")
	for i, alert := range resp.Alerts {
		fmt.Printf("  [%d] ID: %s, Device: %s, Loại: %s, Mức độ: %s, Trạng thái: %s\n",
			i+1, alert.Id, alert.DeviceId, alert.AlertType, alert.Severity, alert.Status)
	}
}

func (c *EnvironmentalAlertClient) TestGetEnvironmentalAlertStatistics() {
	fmt.Println("\n=== Kiểm thử Thống kê Cảnh báo Môi trường ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.environmentalAlertClient.GetEnvironmentalAlertStatistics(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Printf("Error calling GetEnvironmentalAlertStatistics: %v\n", err)
		return
	}

	fmt.Printf("Kết quả thống kê cảnh báo môi trường:\n")
	if resp.Statistics != nil {
		fmt.Printf("Tổng số cảnh báo: %d\n", resp.Statistics.TotalAlerts)
		fmt.Printf("Cảnh báo đang hoạt động: %d\n", resp.Statistics.ActiveAlerts)
		fmt.Printf("Cảnh báo đã xử lý: %d\n", resp.Statistics.ResolvedAlerts)
		fmt.Printf("Cảnh báo quá hạn: %d\n", resp.Statistics.OverdueAlerts)
	}
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== Ứng dụng kiểm thử gRPC Environmental Alert Service ===")
	fmt.Println("1. Dịch vụ Cảnh báo Môi trường")
	fmt.Println("0. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printEnvironmentalAlertMenu() {
	fmt.Println("\n=== Dịch vụ Cảnh báo Môi trường ===")
	fmt.Println("1. Tạo cảnh báo môi trường")
	fmt.Println("2. Lấy cảnh báo môi trường")
	fmt.Println("3. Liệt kê cảnh báo môi trường")
	fmt.Println("4. Thống kê cảnh báo môi trường")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Đang kết nối tới máy chủ gRPC tại %s...\n", address)
	client, err := NewEnvironmentalAlertClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Kết nối thành công!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Dịch vụ Cảnh báo Môi trường
			for {
				printEnvironmentalAlertMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateEnvironmentalAlert()
				case "2":
					client.TestGetEnvironmentalAlert()
				case "3":
					client.TestListEnvironmentalAlerts()
				case "4":
					client.TestGetEnvironmentalAlertStatistics()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
	}
}
