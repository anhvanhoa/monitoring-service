package environmental_alert

import "github.com/anhvanhoa/service-core/domain/oops"

var (
	ErrAlertNotFound      = oops.New("Không tìm thấy cảnh báo")
	ErrUpdateAlert        = oops.New("Lỗi khi cập nhật cảnh báo")
	ErrDeleteAlert        = oops.New("Lỗi khi xóa cảnh báo")
	ErrListAlert          = oops.New("Lỗi khi lấy danh sách cảnh báo")
	ErrGetAlertStatistics = oops.New("Lỗi khi lấy thống kê cảnh báo")
	ErrCreateAlert        = oops.New("Lỗi khi tạo cảnh báo")
)
