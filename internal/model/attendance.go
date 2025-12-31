package model

import "time"

type Attendance struct {
	ID        int64      `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username  string     `json:"username" gorm:"column:username;type:varchar(48);not null"`
	Date      time.Time  `json:"date" gorm:"column:date;type:date;not null"`
	CheckIn   time.Time  `json:"check_in" gorm:"column:check_in;type:timestamp;not null"`
	CheckOut  *time.Time `json:"check_out,omitempty" gorm:"column:check_out;type:timestamp"`
	Longitude string     `json:"longitude,omitempty" gorm:"column:longitude"`
	Latitude  string     `json:"latitude,omitempty" gorm:"column:latitude"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;type:timestamp;autoUpdateTime"`

	User *User `json:"user,omitempty" gorm:"foreignKey:Username;references:Name"`
}

func (Attendance) TableName() string {
	return "hris_attendance"
}
