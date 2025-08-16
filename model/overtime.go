package model

import "time"

type Overtime struct {
	ID          int32     `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Date        time.Time `json:"date" gorm:"column:Date;type:date;index:idx_date_overtime,unique;not null"`
	Username    string    `json:"username" gorm:"column:Username;type:varchar(48);not null"`
	StartTime   time.Time `json:"startTime" gorm:"column:StartTime;type:timestamp;not null"`
	EndTime     time.Time `json:"endTime" gorm:"column:EndTime;type:timestamp;not null"`
	Hours       float64   `json:"hours" gorm:"column:Hours;type:numeric(18,6);not null"`
	Description string    `json:"description" gorm:"column:Description;type:varchar(255);not null"`
	Status      string    `json:"status" gorm:"column:Status;type:varchar(48);not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Username"`
}

func (Overtime) TableName() string {
	return "hris_overtime"
}

func AddOvertime(overtime Overtime) error {
	return db.Create(&overtime).Error
}

func UpdateOvertime(overtime Overtime) error {
	return db.Model(Overtime{}).Where(&Overtime{ID: overtime.ID}).Updates(overtime).Error
}

func GetOvertimeById(id int32) (Overtime, error) {
	var overtime Overtime
	err := db.Model(Overtime{}).Where(&Overtime{ID: id}).First(&overtime).Error
	return overtime, err
}

func GetOvertime(user string, date time.Time) (Overtime, error) {
	var overtime Overtime
	err := db.Model(Overtime{}).Where(&Overtime{Username: user, Date: date}).First(&overtime).Error
	return overtime, err
}

func GetOvertimes(cond *Overtime) ([]Overtime, error) {
	var overtimes []Overtime
	err := db.Model(Overtime{}).Where(cond).Find(&overtimes).Error
	return overtimes, err
}

func DeleteOvertime(id int32) error {
	return db.Model(Overtime{}).Where(&Overtime{ID: id}).Delete(&Overtime{}).Error
}
