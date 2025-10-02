CREATE TABLE environmental_alerts (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()), -- ID duy nhất của cảnh báo
    device_id VARCHAR(36),                       -- Liên kết tới thiết bị IoT phát hiện cảnh báo
    alert_type VARCHAR(100) COMMENT 'temperature_high, temperature_low, humidity_high, humidity_low, ph_high, ph_low, water_shortage, equipment_failure', 
        -- Loại cảnh báo (nhiệt độ, độ ẩm, pH, thiếu nước, hỏng thiết bị...)
    current_value DECIMAL(12,4),                 -- Giá trị đo được tại thời điểm cảnh báo
    threshold_value DECIMAL(12,4),               -- Ngưỡng quy định (giới hạn gây ra cảnh báo)
    threshold_type VARCHAR(20) COMMENT 'min, max, range', 
        -- Kiểu ngưỡng: nhỏ nhất (min), lớn nhất (max), hoặc trong khoảng (range)
    severity VARCHAR(50) COMMENT 'info, warning, critical, emergency', 
        -- Mức độ nghiêm trọng: thông tin, cảnh báo, nghiêm trọng, khẩn cấp
    priority INTEGER DEFAULT 3 COMMENT '1=highest, 5=lowest', 
        -- Mức độ ưu tiên (1 cao nhất, 5 thấp nhất)
    triggered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
        -- Thời điểm hệ thống kích hoạt cảnh báo
    acknowledged_at TIMESTAMP,                   -- Thời điểm cảnh báo được người dùng xác nhận đã nhận
    acknowledged_by VARCHAR(36),                 -- Người dùng đã xác nhận cảnh báo
    resolved_at TIMESTAMP,                       -- Thời điểm cảnh báo được xử lý xong
    resolved_by VARCHAR(36),                     -- Người dùng đã xử lý cảnh báo
    auto_resolve BOOLEAN DEFAULT FALSE,          -- Cho biết cảnh báo có tự động đóng khi điều kiện an toàn được phục hồi không
    status VARCHAR(50) DEFAULT 'active' COMMENT 'active, acknowledged, resolved, escalated, ignored', 
        -- Trạng thái cảnh báo: đang hoạt động, đã xác nhận, đã xử lý, leo thang, bỏ qua
    auto_action_taken VARCHAR(500),              -- Hành động tự động đã thực hiện (nếu có, ví dụ: bật quạt, bơm nước...)
    escalation_level INTEGER DEFAULT 1,          -- Cấp độ leo thang (nếu cảnh báo không được xử lý thì sẽ nâng mức độ)
    notification_sent JSON COMMENT 'Tracking notifications sent', 
        -- Ghi lại các thông báo đã gửi (email, SMS, app notification...)
    impact_assessment TEXT,                      -- Đánh giá tác động (ảnh hưởng tới cây trồng, thiết bị...)
    resolution_notes TEXT,                       -- Ghi chú về cách xử lý cảnh báo
    created_by VARCHAR(36),                      -- Người tạo bản ghi cảnh báo
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    
    INDEX idx_environmental_alerts_device (device_id), -- Index tìm kiếm theo thiết bị
    INDEX idx_environmental_alerts_status_severity (status, severity), -- Index theo trạng thái & mức độ
    INDEX idx_environmental_alerts_triggered (triggered_at), -- Index theo thời gian cảnh báo
    INDEX idx_environmental_alerts_type (alert_type)   -- Index theo loại cảnh báo
);
