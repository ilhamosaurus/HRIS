package dto

import "github.com/ilhamosaurus/HRIS/internal/model"

type UpdateAttendanceRequest struct {
	Date     string `json:"date" validate:"omitempty"`
	CheckIn  string `json:"check_in" validate:"omitempty,datetime"`
	CheckOut string `json:"check_out" validate:"omitempty,datetime"`
}

type AttendancesQuery struct {
	Username string `query:"username" validate:"omitempty,gte=3"`
	Date     string `query:"date"`
	Page     int    `query:"page" validate:"min=1"`
	PageSize int    `query:"page_size" validate:"min=1,max=100"`
}

type AttendanceResponse struct {
	List     []*model.Attendance `json:"list"`
	Total    int64               `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}
