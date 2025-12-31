package model

import "time"

type UserActivity struct {
	ID            int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Time          time.Time `json:"time" gorm:"column:time;type:timestamp;autoCreateTime"`
	Username      string    `json:"username" gorm:"column:username;type:varchar(48);not null"`
	Address       string    `json:"address" gorm:"column:address;type:varchar(48);not null"`
	Feature       string    `json:"feature" gorm:"column:feature;type:varchar(48);not null"`
	AccessType    string    `json:"access_type" gorm:"column:access_type;type:varchar(48);not null"`
	AccessDetails *string   `json:"access_details" gorm:"column:access_details"`

	User *User `json:"user,omitempty" gorm:"foreignKey:Username;references:Name"`
}

func (UserActivity) TableName() string {
	return "hris_user_activity"
}
