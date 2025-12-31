package model

import "time"

type Payslip struct {
	Code          string     `json:"code" gorm:"column:code;primaryKey;type:varchar(255);not null"`
	Date          time.Time  `json:"date" gorm:"column:date;type:date;not null"`
	Username      string     `json:"username" gorm:"column:username;not null"`
	Attendance    int32      `json:"attendance" gorm:"column:attendance;not null"`
	BaseSalary    float64    `json:"base_salary" gorm:"column:base_salary;type:numeric(18,6);not null"`
	OvertimeHours float64    `json:"overtime_hours" gorm:"column:overtime_hours;type:numeric(18,6)"`
	OvertimePay   float64    `json:"overtime_pay" gorm:"column:overtime_pay;type:numeric(18,6)"`
	Reimburse     float64    `json:"reimburse" gorm:"column:reimburse;type:numeric(18,6)"`
	TakeHomePay   float64    `json:"take_home_pay" gorm:"column:take_home_pay;type:numeric(18,6);not null"`
	Processed     bool       `json:"processed" gorm:"column:processed;type:bool;not null"`
	ProcessedAt   *time.Time `json:"processed_at" gorm:"column:processed_at;type:timestamp"`
	CreatedAt     time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"column:updated_at;type:timestamp;autoUpdateTime"`

	User User `json:"user,omitempty" gorm:"foreignKey:Username;references:Name"`
}

func (Payslip) TableName() string {
	return "hris_payslip"
}
