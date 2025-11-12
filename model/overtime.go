package model

import "time"

type Overtime struct {
	ID          int32     `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Date        time.Time `json:"date" gorm:"column:Date;type:date;index:idx_date_overtime,unique;not null"`
	Username    string    `json:"username" gorm:"column:Username;type:varchar(48);not null"`
	StartTime   time.Time `json:"startTime" gorm:"column:StartTime;type:timestamp;not null"`
	EndTime     time.Time `json:"endTime" gorm:"column:EndTime;type:timestamp;not null"`
	Hours       float64   `json:"hours" gorm:"column:Hours;type:numeric(18,6);not null"`
	Description string    `json:"description" gorm:"column:Description"`
	Status      string    `json:"status" gorm:"column:Status;type:varchar(48);not null"`
	Approval    string    `json:"approval" gorm:"column:Approval;not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Name"`
}

func (Overtime) TableName() string {
	return "hris_overtime"
}

func (o Overtime) CalculateHours() float64 {
	if o.StartTime.IsZero() || o.EndTime.IsZero() {
		return 0
	}

	if o.EndTime.Before(o.StartTime) {
		return 0
	}

	return o.EndTime.Sub(o.StartTime).Hours()
}

func (m *Model) AddOvertime(overtime *Overtime) error {
	return m.db.Create(overtime).Error
}

func (m *Model) UpdateOvertime(overtime *Overtime) error {
	return m.db.Model(Overtime{}).Where(&Overtime{ID: overtime.ID}).Updates(overtime).Error
}

func (m *Model) GetOvertimeById(id int32) (*Overtime, error) {
	var overtime Overtime
	err := m.db.Model(Overtime{}).Where(&Overtime{ID: id}).First(&overtime).Error
	return &overtime, err
}

func (m *Model) GetOvertime(user, status string, date time.Time) *Overtime {
	var overtime Overtime
	m.db.Model(Overtime{}).Where(&Overtime{Username: user, Date: date, Status: status}).First(&overtime)
	return &overtime
}

func (m *Model) GetOvertimes(cond *Overtime) ([]Overtime, error) {
	var overtimes []Overtime
	err := m.db.Model(Overtime{}).Where(cond).Find(&overtimes).Error
	return overtimes, err
}

func (m *Model) DeleteOvertime(id int32) error {
	return m.db.Model(Overtime{}).Where(&Overtime{ID: id}).Delete(&Overtime{}).Error
}
