-- Seed data cho bảng environmental_alerts
INSERT INTO
    environmental_alerts (
        id,
        device_id,
        alert_type,
        current_value,
        threshold_value,
        threshold_type,
        severity,
        priority,
        triggered_at,
        status,
        auto_resolve,
        escalation_level,
        impact_assessment,
        created_by
    )
VALUES
    -- Cảnh báo nhiệt độ cao
    (
        '550e8400-e29b-41d4-a716-446655440001',
        'device-001',
        'temperature_high',
        35.5,
        30.0,
        'max',
        'critical',
        2,
        '2024-01-15 10:30:00',
        'active',
        false,
        1,
        'Nhiệt độ vượt ngưỡng có thể ảnh hưởng đến sự phát triển của cây trồng',
        'system'
    ),
    -- Cảnh báo độ ẩm thấp
    (
        '550e8400-e29b-41d4-a716-446655440002',
        'device-002',
        'humidity_low',
        25.0,
        40.0,
        'min',
        'warning',
        3,
        '2024-01-15 11:15:00',
        'acknowledged',
        true,
        1,
        'Độ ẩm thấp có thể gây khô hạn cho cây trồng',
        'system'
    ),
    -- Cảnh báo pH cao
    (
        '550e8400-e29b-41d4-a716-446655440003',
        'device-003',
        'ph_high',
        8.5,
        7.5,
        'max',
        'warning',
        3,
        '2024-01-15 12:00:00',
        'resolved',
        true,
        1,
        'pH cao có thể ảnh hưởng đến khả năng hấp thụ dinh dưỡng của cây',
        'system'
    ),
    -- Cảnh báo thiếu nước
    (
        '550e8400-e29b-41d4-a716-446655440004',
        'device-004',
        'water_shortage',
        10.0,
        20.0,
        'min',
        'critical',
        1,
        '2024-01-15 13:45:00',
        'escalated',
        false,
        2,
        'Mức nước thấp nghiêm trọng, cần bổ sung nước ngay lập tức',
        'system'
    ),
    -- Cảnh báo hỏng thiết bị
    (
        '550e8400-e29b-41d4-a716-446655440005',
        'device-005',
        'equipment_failure',
        NULL,
        NULL,
        NULL,
        'emergency',
        1,
        '2024-01-15 14:20:00',
        'active',
        false,
        1,
        'Hệ thống tưới nước tự động bị hỏng, cần kiểm tra và sửa chữa',
        'system'
    ),
    -- Cảnh báo nhiệt độ thấp
    (
        '550e8400-e29b-41d4-a716-446655440006',
        'device-001',
        'temperature_low',
        5.0,
        10.0,
        'min',
        'warning',
        3,
        '2024-01-15 15:30:00',
        'resolved',
        true,
        1,
        'Nhiệt độ thấp có thể làm chậm quá trình phát triển của cây',
        'system'
    ),
    -- Cảnh báo độ ẩm cao
    (
        '550e8400-e29b-41d4-a716-446655440007',
        'device-002',
        'humidity_high',
        85.0,
        80.0,
        'max',
        'info',
        4,
        '2024-01-15 16:00:00',
        'ignored',
        true,
        1,
        'Độ ẩm cao có thể tạo điều kiện cho nấm mốc phát triển',
        'system'
    ),
    -- Cảnh báo pH thấp
    (
        '550e8400-e29b-41d4-a716-446655440008',
        'device-003',
        'ph_low',
        5.5,
        6.0,
        'min',
        'warning',
        3,
        '2024-01-15 17:15:00',
        'acknowledged',
        false,
        1,
        'pH thấp có thể gây độc cho rễ cây',
        'system'
    ),
    -- Cảnh báo đã được xử lý
    (
        '550e8400-e29b-41d4-a716-446655440009',
        'device-006',
        'temperature_high',
        32.0,
        30.0,
        'max',
        'critical',
        2,
        '2024-01-14 09:00:00',
        'resolved',
        true,
        1,
        'Nhiệt độ đã được điều chỉnh về mức bình thường',
        'system'
    ),
    -- Cảnh báo với hành động tự động
    (
        '550e8400-e29b-41d4-a716-446655440010',
        'device-007',
        'humidity_low',
        30.0,
        40.0,
        'min',
        'warning',
        3,
        '2024-01-15 18:30:00',
        'resolved',
        true,
        1,
        'Độ ẩm đã được tăng lên bằng hệ thống phun sương',
        'system'
    );

-- Cập nhật thông tin xử lý cho các cảnh báo đã được xử lý
UPDATE environmental_alerts
SET
    resolved_at = '2024-01-15 16:30:00',
    resolved_by = 'admin-001',
    resolution_notes = 'Đã điều chỉnh nhiệt độ bằng cách bật quạt thông gió',
    auto_action_taken = 'Bật quạt thông gió tự động'
WHERE
    id = '550e8400-e29b-41d4-a716-446655440009';

UPDATE environmental_alerts
SET
    resolved_at = '2024-01-15 19:00:00',
    resolved_by = 'admin-002',
    resolution_notes = 'Đã bật hệ thống phun sương để tăng độ ẩm',
    auto_action_taken = 'Bật hệ thống phun sương tự động'
WHERE
    id = '550e8400-e29b-41d4-a716-446655440010';

-- Cập nhật thông tin xác nhận cho các cảnh báo đã được xác nhận
UPDATE environmental_alerts
SET
    acknowledged_at = '2024-01-15 11:30:00',
    acknowledged_by = 'user-001'
WHERE
    id = '550e8400-e29b-41d4-a716-446655440002';

UPDATE environmental_alerts
SET
    acknowledged_at = '2024-01-15 17:30:00',
    acknowledged_by = 'user-002'
WHERE
    id = '550e8400-e29b-41d4-a716-446655440008';