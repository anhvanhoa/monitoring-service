CREATE TABLE environmental_alerts (
    id VARCHAR(36) PRIMARY KEY DEFAULT (gen_random_uuid()), -- ID duy nhất của cảnh báo
    device_id VARCHAR(36),                       -- Liên kết tới thiết bị IoT phát hiện cảnh báo
    alert_type VARCHAR(100),                     -- Loại cảnh báo (nhiệt độ, độ ẩm, pH, thiếu nước, hỏng thiết bị...)
    current_value DECIMAL(12,4),                 -- Giá trị đo được tại thời điểm cảnh báo
    threshold_value DECIMAL(12,4),               -- Ngưỡng quy định (giới hạn gây ra cảnh báo)
    threshold_type VARCHAR(20),                  -- Kiểu ngưỡng: nhỏ nhất (min), lớn nhất (max), hoặc trong khoảng (range)
    severity VARCHAR(50),                        -- Mức độ nghiêm trọng: thông tin, cảnh báo, nghiêm trọng, khẩn cấp
    priority INTEGER DEFAULT 3,                  -- Mức độ ưu tiên (1 cao nhất, 5 thấp nhất)
    triggered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm hệ thống kích hoạt cảnh báo
    acknowledged_at TIMESTAMP,                   -- Thời điểm cảnh báo được người dùng xác nhận đã nhận
    acknowledged_by VARCHAR(36),                 -- Người dùng đã xác nhận cảnh báo
    resolved_at TIMESTAMP,                       -- Thời điểm cảnh báo được xử lý xong
    resolved_by VARCHAR(36),                     -- Người dùng đã xử lý cảnh báo
    auto_resolve BOOLEAN DEFAULT FALSE,          -- Cho biết cảnh báo có tự động đóng khi điều kiện an toàn được phục hồi không
    status VARCHAR(50) DEFAULT 'active',         -- Trạng thái cảnh báo: đang hoạt động, đã xác nhận, đã xử lý, leo thang, bỏ qua
    auto_action_taken VARCHAR(500),              -- Hành động tự động đã thực hiện (nếu có, ví dụ: bật quạt, bơm nước...)
    escalation_level INTEGER DEFAULT 1,          -- Cấp độ leo thang (nếu cảnh báo không được xử lý thì sẽ nâng mức độ)
    notification_sent JSON,                      -- Ghi lại các thông báo đã gửi (email, SMS, app notification...)
    impact_assessment TEXT,                      -- Đánh giá tác động (ảnh hưởng tới cây trồng, thiết bị...)
    resolution_notes TEXT,                       -- Ghi chú về cách xử lý cảnh báo
    created_by VARCHAR(36),                      -- Người tạo bản ghi cảnh báo
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Thời điểm cập nhật
);

-- Tạo các index
CREATE INDEX idx_environmental_alerts_device ON environmental_alerts (device_id);
CREATE INDEX idx_environmental_alerts_status_severity ON environmental_alerts (status, severity);
CREATE INDEX idx_environmental_alerts_triggered ON environmental_alerts (triggered_at);
CREATE INDEX idx_environmental_alerts_type ON environmental_alerts (alert_type);

-- Tạo trigger để tự động cập nhật updated_at
CREATE TRIGGER update_environmental_alerts_updated_at
    BEFORE UPDATE ON environmental_alerts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Thêm comment cho bảng và các cột
COMMENT ON TABLE environmental_alerts IS 'Bảng lưu trữ các cảnh báo môi trường từ hệ thống IoT';
COMMENT ON COLUMN environmental_alerts.alert_type IS 'temperature_high, temperature_low, humidity_high, humidity_low, ph_high, ph_low, water_shortage, equipment_failure';
COMMENT ON COLUMN environmental_alerts.threshold_type IS 'min, max, range';
COMMENT ON COLUMN environmental_alerts.severity IS 'info, warning, critical, emergency';
COMMENT ON COLUMN environmental_alerts.priority IS '1=highest, 5=lowest';
COMMENT ON COLUMN environmental_alerts.status IS 'active, acknowledged, resolved, escalated, ignored';
COMMENT ON COLUMN environmental_alerts.notification_sent IS 'Tracking notifications sent';
