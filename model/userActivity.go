package model

import "time"

type UserActivity struct {
	ID            int32     `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Time          time.Time `json:"time" gorm:"column:Time;type:timestamp;autoCreateTime"`
	Username      string    `json:"username" gorm:"column:Username;type:varchar(48);not null"`
	Address       string    `json:"address" gorm:"column:Address;type:varchar(48);not null"`
	Feature       string    `json:"feature" gorm:"column:Feature;type:varchar(48);not null"`
	AccessType    string    `json:"accessType" gorm:"column:AccessType;type:varchar(48);not null"`
	AccessDetails *string   `json:"accessDetails" gorm:"column:AccessDetails"`

	User User `json:"user" gorm:"foreignKey:Username;references:Name"`
}

func (UserActivity) TableName() string {
	return "hris_user_activity"
}

func AddUserActivity(log UserActivity) error {
	return db.Create(&log).Error
}

func GetUserActivity(cond *UserActivity) ([]UserActivity, error) {
	var logs []UserActivity
	err := db.Model(UserActivity{}).Where(cond).Find(&logs).Error
	return logs, err
}
