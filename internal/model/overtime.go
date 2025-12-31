package model

import "time"

type Overtime struct {
	ID          int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Date        time.Time `json:"date" gorm:"column:date;type:date;index:idx_date_overtime;not null"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(48);not null"`
	StartTime   time.Time `json:"start_time" gorm:"column:start_time;type:timestamp;not null"`
	EndTime     time.Time `json:"end_time" gorm:"column:end_time;type:timestamp;not null"`
	Hours       float64   `json:"hours" gorm:"column:hours;type:numeric(18,6);not null"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status;type:varchar(48);not null"`
	Approval    string    `json:"approval" gorm:"column:approval;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Name"`
}

func (Overtime) TableName() string {
	return "hris_overtime"
}

func (o *Overtime) CalculateHours() float64 {
	if o.StartTime.IsZero() || o.EndTime.IsZero() {
		return 0
	}

	if o.EndTime.Before(o.StartTime) {
		return 0
	}

	return o.EndTime.Sub(o.StartTime).Hours()
}
