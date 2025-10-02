package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"github.com/anhvanhoa/service-core/domain/oops"
)

// AlertType định nghĩa các loại cảnh báo
type AlertType string

const (
	AlertTypeTemperatureHigh  AlertType = "temperature_high"
	AlertTypeTemperatureLow   AlertType = "temperature_low"
	AlertTypeHumidityHigh     AlertType = "humidity_high"
	AlertTypeHumidityLow      AlertType = "humidity_low"
	AlertTypePHHigh           AlertType = "ph_high"
	AlertTypePHLow            AlertType = "ph_low"
	AlertTypeWaterShortage    AlertType = "water_shortage"
	AlertTypeEquipmentFailure AlertType = "equipment_failure"
)

// ThresholdType định nghĩa kiểu ngưỡng
type ThresholdType string

const (
	ThresholdTypeMin   ThresholdType = "min"
	ThresholdTypeMax   ThresholdType = "max"
	ThresholdTypeRange ThresholdType = "range"
)

// Severity định nghĩa mức độ nghiêm trọng
type Severity string

const (
	SeverityInfo      Severity = "info"
	SeverityWarning   Severity = "warning"
	SeverityCritical  Severity = "critical"
	SeverityEmergency Severity = "emergency"
)

// Status định nghĩa trạng thái cảnh báo
type Status string

const (
	StatusActive       Status = "active"
	StatusAcknowledged Status = "acknowledged"
	StatusResolved     Status = "resolved"
	StatusEscalated    Status = "escalated"
	StatusIgnored      Status = "ignored"
)

// NotificationSent lưu trữ thông tin các thông báo đã gửi
type NotificationSent struct {
	Email      bool
	SMS        bool
	Push       bool
	Webhook    bool
	SentAt     time.Time
	Recipients []string
}

// Value implements driver.Valuer interface for JSON storage
func (n NotificationSent) Value() (driver.Value, error) {
	return json.Marshal(n)
}

// Scan implements sql.Scanner interface for JSON retrieval
func (n *NotificationSent) Scan(value any) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		if str, ok := value.(string); ok {
			bytes = []byte(str)
		} else {
			return fmt.Errorf("cannot scan %T into NotificationSent", value)
		}
	}

	return json.Unmarshal(bytes, n)
}

type EnvironmentalAlert struct {
	tableName        struct{} `pg:"environmental_alerts"`
	ID               string
	DeviceID         string
	AlertType        AlertType
	CurrentValue     float64
	ThresholdValue   float64
	ThresholdType    ThresholdType
	Severity         Severity
	Priority         int
	TriggeredAt      time.Time
	AcknowledgedAt   *time.Time
	AcknowledgedBy   string
	ResolvedAt       *time.Time
	ResolvedBy       string
	AutoResolve      bool
	Status           Status
	AutoActionTaken  string
	EscalationLevel  int
	NotificationSent string
	ImpactAssessment string
	ResolutionNotes  string
	CreatedBy        string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}

// TableName trả về tên bảng trong database
func (e *EnvironmentalAlert) TableName() any {
	return e.tableName
}

// IsActive kiểm tra xem cảnh báo có đang hoạt động không
func (ea *EnvironmentalAlert) IsActive() bool {
	return ea.Status == StatusActive
}

// IsResolved kiểm tra xem cảnh báo đã được xử lý chưa
func (ea *EnvironmentalAlert) IsResolved() bool {
	return ea.Status == StatusResolved
}

// IsAcknowledged kiểm tra xem cảnh báo đã được xác nhận chưa
func (ea *EnvironmentalAlert) IsAcknowledged() bool {
	return ea.Status == StatusAcknowledged
}

// CanBeResolved kiểm tra xem cảnh báo có thể được xử lý không
func (ea *EnvironmentalAlert) CanBeResolved() bool {
	return ea.Status == StatusActive || ea.Status == StatusAcknowledged
}

// CanBeAcknowledged kiểm tra xem cảnh báo có thể được xác nhận không
func (ea *EnvironmentalAlert) CanBeAcknowledged() bool {
	return ea.Status == StatusActive
}

// Acknowledge đánh dấu cảnh báo đã được xác nhận
func (ea *EnvironmentalAlert) Acknowledge(userID string) {
	now := time.Now()
	ea.AcknowledgedAt = &now
	ea.AcknowledgedBy = userID
	ea.Status = StatusAcknowledged
	ea.UpdatedAt = &now
}

// Resolve đánh dấu cảnh báo đã được xử lý
func (ea *EnvironmentalAlert) Resolve(userID string, notes string) {
	now := time.Now()
	ea.ResolvedAt = &now
	ea.ResolvedBy = userID
	ea.Status = StatusResolved
	ea.UpdatedAt = &now
	ea.ResolutionNotes = notes
}

// Escalate nâng cấp độ cảnh báo
func (ea *EnvironmentalAlert) Escalate() {
	ea.EscalationLevel++
	ea.Status = StatusEscalated
	now := time.Now()
	ea.UpdatedAt = &now
}

// Ignore bỏ qua cảnh báo
func (ea *EnvironmentalAlert) Ignore() {
	ea.Status = StatusIgnored
	now := time.Now()
	ea.UpdatedAt = &now
}

// SetAutoAction ghi lại hành động tự động đã thực hiện
func (ea *EnvironmentalAlert) SetAutoAction(action string) {
	ea.AutoActionTaken = action
	now := time.Now()
	ea.UpdatedAt = &now
}

// SetNotificationSent cập nhật thông tin thông báo đã gửi
func (ea *EnvironmentalAlert) SetNotificationSent(notification NotificationSent) {
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return
	}
	ea.NotificationSent = string(notificationJSON)
	now := time.Now()
	ea.UpdatedAt = &now
}

// GetDurationSinceTriggered tính thời gian từ khi cảnh báo được kích hoạt
func (ea *EnvironmentalAlert) GetDurationSinceTriggered() time.Duration {
	return time.Since(ea.TriggeredAt)
}

// IsOverdue kiểm tra xem cảnh báo có bị quá hạn không (dựa trên priority)
func (ea *EnvironmentalAlert) IsOverdue() bool {
	if ea.IsResolved() {
		return false
	}

	duration := ea.GetDurationSinceTriggered()

	// Thời gian quá hạn dựa trên priority (1=highest, 5=lowest)
	overdueThresholds := map[int]time.Duration{
		1: 15 * time.Minute, // Emergency: 15 phút
		2: 30 * time.Minute, // Critical: 30 phút
		3: 1 * time.Hour,    // Warning: 1 giờ
		4: 2 * time.Hour,    // Info: 2 giờ
		5: 4 * time.Hour,    // Lowest: 4 giờ
	}

	threshold, exists := overdueThresholds[ea.Priority]
	if !exists {
		threshold = 1 * time.Hour // Default
	}

	return duration > threshold
}

// GetSeverityLevel trả về số thứ tự mức độ nghiêm trọng (1-4)
func (ea *EnvironmentalAlert) GetSeverityLevel() int {
	severityLevels := map[Severity]int{
		SeverityInfo:      1,
		SeverityWarning:   2,
		SeverityCritical:  3,
		SeverityEmergency: 4,
	}

	level, exists := severityLevels[ea.Severity]
	if !exists {
		return 1
	}

	return level
}

func (ea *EnvironmentalAlert) Validate() error {
	if !ea.isValidAlertType(ea.AlertType) {
		return oops.Newf("Loại cảnh báo không hợp lệ: %s", ea.AlertType)
	}
	if !ea.isValidSeverity(ea.Severity) {
		return oops.Newf("Mức độ nghiêm trọng không hợp lệ: %s", ea.Severity)
	}

	if !ea.isValidStatus(ea.Status) {
		return oops.Newf("Trạng thái cảnh báo không hợp lệ: %s", ea.Status)
	}

	if ea.Priority < 1 || ea.Priority > 5 {
		return oops.Newf("Mức độ ưu tiên phải từ 1 đến 5, nhận được: %d", ea.Priority)
	}

	if ea.EscalationLevel < 1 {
		return oops.Newf("Cấp độ leo thang phải >= 1, nhận được: %d", ea.EscalationLevel)
	}

	if !ea.isValidThresholdType(ea.ThresholdType) {
		return oops.Newf("Kiểu ngưỡng không hợp lệ: %s", ea.ThresholdType)
	}

	// if err := ea.validateValueConsistency(); err != nil {
	// 	return err
	// }

	// if err := ea.validateTimestampConsistency(); err != nil {
	// 	return err
	// }

	// if err := ea.validateBusinessRules(); err != nil {
	// 	return err
	// }

	return nil
}

func (ea *EnvironmentalAlert) isValidAlertType(alertType AlertType) bool {
	validTypes := []AlertType{
		AlertTypeTemperatureHigh,
		AlertTypeTemperatureLow,
		AlertTypeHumidityHigh,
		AlertTypeHumidityLow,
		AlertTypePHHigh,
		AlertTypePHLow,
		AlertTypeWaterShortage,
		AlertTypeEquipmentFailure,
	}

	return slices.Contains(validTypes, alertType)
}

func (ea *EnvironmentalAlert) isValidSeverity(severity Severity) bool {
	validSeverities := []Severity{
		SeverityInfo,
		SeverityWarning,
		SeverityCritical,
		SeverityEmergency,
	}

	return slices.Contains(validSeverities, severity)
}

func (ea *EnvironmentalAlert) isValidStatus(status Status) bool {
	validStatuses := []Status{
		StatusActive,
		StatusAcknowledged,
		StatusResolved,
		StatusEscalated,
		StatusIgnored,
	}

	return slices.Contains(validStatuses, status)
}

func (ea *EnvironmentalAlert) isValidThresholdType(thresholdType ThresholdType) bool {
	validTypes := []ThresholdType{
		ThresholdTypeMin,
		ThresholdTypeMax,
		ThresholdTypeRange,
	}

	return slices.Contains(validTypes, thresholdType)
}

func (ea *EnvironmentalAlert) validateValueConsistency() error {
	if ea.CurrentValue != 0 && ea.ThresholdValue != 0 && ea.ThresholdType != "" {
		current := ea.CurrentValue
		threshold := ea.ThresholdValue
		thresholdType := ea.ThresholdType

		switch ea.AlertType {
		case AlertTypeTemperatureHigh, AlertTypeHumidityHigh, AlertTypePHHigh:
			if thresholdType == ThresholdTypeMax && current <= threshold {
				return oops.Newf("Đối với %s với ngưỡng tối đa, giá trị hiện tại (%.2f) phải > giá trị ngưỡng (%.2f)",
					ea.AlertType, current, threshold)
			}
		case AlertTypeTemperatureLow, AlertTypeHumidityLow, AlertTypePHLow:
			if thresholdType == ThresholdTypeMin && current >= threshold {
				return oops.Newf("Đối với %s với ngưỡng tối thiểu, giá trị hiện tại (%.2f) phải < giá trị ngưỡng (%.2f)",
					ea.AlertType, current, threshold)
			}
		case AlertTypeWaterShortage:
			if thresholdType == ThresholdTypeMin && current >= threshold {
				return oops.Newf("Đối với thiếu nước với ngưỡng tối thiểu, giá trị hiện tại (%.2f) phải < giá trị ngưỡng (%.2f)",
					current, threshold)
			}
		}
	}

	if ea.CurrentValue != 0 && ea.CurrentValue < 0 {
		switch ea.AlertType {
		case AlertTypeTemperatureHigh, AlertTypeTemperatureLow, AlertTypeHumidityHigh, AlertTypeHumidityLow:
			return fmt.Errorf("không cho phép giá trị âm đối với %s", ea.AlertType)
		}
	}
	return nil
}

func (ea *EnvironmentalAlert) validateTimestampConsistency() error {
	if ea.TriggeredAt.IsZero() {
		return oops.Newf("Thời điểm kích hoạt cảnh báo là bắt buộc")
	}

	if ea.AcknowledgedAt != nil && ea.AcknowledgedAt.Before(ea.TriggeredAt) {
		return oops.Newf("Thời điểm xác nhận (%v) không thể trước thời điểm kích hoạt (%v)",
			*ea.AcknowledgedAt, ea.TriggeredAt)
	}

	if ea.ResolvedAt != nil && ea.ResolvedAt.Before(ea.TriggeredAt) {
		return oops.Newf("Thời điểm xử lý (%v) không thể trước thời điểm kích hoạt (%v)",
			*ea.ResolvedAt, ea.TriggeredAt)
	}

	if ea.AcknowledgedAt != nil && ea.ResolvedAt != nil && ea.ResolvedAt.Before(*ea.AcknowledgedAt) {
		return oops.Newf("Thời điểm xử lý (%v) không thể trước thời điểm xác nhận (%v)",
			*ea.ResolvedAt, *ea.AcknowledgedAt)
	}

	return nil
}

func (ea *EnvironmentalAlert) validateBusinessRules() error {
	if ea.Status == StatusAcknowledged {
		if ea.AcknowledgedAt == nil {
			return oops.Newf("Thời điểm xác nhận là bắt buộc khi trạng thái là đã xác nhận")
		}
		if ea.AcknowledgedBy == "" {
			return oops.Newf("Người xác nhận là bắt buộc khi trạng thái là đã xác nhận")
		}
	}

	if ea.Status == StatusResolved {
		if ea.ResolvedAt == nil {
			return oops.Newf("Thời điểm xử lý là bắt buộc khi trạng thái là đã xử lý")
		}
		if ea.ResolvedBy == "" {
			return oops.Newf("Người xử lý là bắt buộc khi trạng thái là đã xử lý")
		}
	}

	if ea.AutoActionTaken != "" && !ea.AutoResolve {
		return oops.Newf("Tự động xử lý phải là true khi có hành động tự động")
	}

	var notificationSent NotificationSent
	if len(ea.NotificationSent) > 0 {
		if err := json.Unmarshal([]byte(ea.NotificationSent), &notificationSent); err != nil {
			return oops.Newf("Invalid JSON in notification_sent: %v", err)
		}
		hasNotification := notificationSent.Email || notificationSent.SMS ||
			notificationSent.Push || notificationSent.Webhook
		if !hasNotification {
			return oops.Newf("Ít nhất một loại thông báo phải được gửi trong notification_sent")
		}
	}

	if ea.EscalationLevel > 5 {
		return oops.Newf("Cấp độ leo thang không được vượt quá 5, nhận được: %d", ea.EscalationLevel)
	}

	if err := ea.validatePrioritySeverityConsistency(); err != nil {
		return err
	}

	return nil
}

func (ea *EnvironmentalAlert) validatePrioritySeverityConsistency() error {
	if ea.Severity == SeverityEmergency && ea.Priority > 2 {
		return oops.Newf("Mức độ khẩn cấp nên có ưu tiên 1 hoặc 2, nhận được: %d", ea.Priority)
	}

	if ea.Severity == SeverityCritical && ea.Priority > 3 {
		return oops.Newf("Mức độ nghiêm trọng nên có ưu tiên 1-3, nhận được: %d", ea.Priority)
	}

	if ea.Severity == SeverityInfo && ea.Priority < 3 {
		return oops.Newf("Mức độ thông tin nên có ưu tiên 3-5, nhận được: %d", ea.Priority)
	}

	return nil
}

type AlertFilters struct {
	DeviceID        string
	AlertType       AlertType
	Status          Status
	Severity        Severity
	Priority        int
	StartDate       *time.Time
	EndDate         *time.Time
	IsOverdue       bool
	EscalationLevel int
	AutoResolve     bool
	CreatedBy       string
	AcknowledgedBy  string
	ResolvedBy      string
}

type AlertStatistics struct {
	TotalAlerts    int64
	ActiveAlerts   int64
	ResolvedAlerts int64
	OverdueAlerts  int64
}
