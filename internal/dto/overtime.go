package dto

type CreateOvertimeRequest struct {
	Date        string `json:"date" validate:"required"`
	Username    string `json:"username" validate:"required"`
	StartTime   string `json:"start_time" validate:"required,datetime"`
	EndTime     string `json:"end_time" validate:"required,datetime"`
	Description string `json:"description"`
}

type OvertimeResponse struct {
	ID          int64   `json:"id"`
	Username    string  `json:"username"`
	Date        string  `json:"date"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Hours       float64 `json:"hours"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Approval    string  `json:"approval"`
}

type UpdateOvertimeRequest struct {
	Date        string `json:"date" validate:"omitempty"`
	Username    string `json:"username" validate:"omitempty"`
	StartTime   string `json:"start_time" validate:"omitempty,datetime"`
	EndTime     string `json:"end_time" validate:"omitempty,datetime"`
	Description string `json:"description" validate:"omitempty"`
}

type ApprovalRequest struct {
	Status string `json:"status" validate:"required,oneof=approved rejected"`
}

type OvertimeQuery struct {
	Date     string `query:"date" validate:"omitempty"`
	Username string `query:"username" validate:"omitempty"`
	Status   string `query:"status" validate:"omitempty,oneof=draft submitted approved rejected done"`
	Approval string `query:"approval" validate:"omitempty"`
	Page     int    `query:"page" validate:"min=1"`
	PageSize int    `query:"page_size" validate:"min=1,max=100"`
}

type OvertimesResponse struct {
	List     []*OvertimeResponse `json:"list"`
	Total    int64               `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
}
