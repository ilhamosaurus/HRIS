package model

import "time"

type Attendance struct {
	ID        int32      `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Username  string     `json:"username" gorm:"column:Username;type:varchar(48);not null"`
	Date      time.Time  `json:"date" gorm:"column:Date;type:date;index:idx_date_attendance,unique;not null"`
	CheckIn   time.Time  `json:"checkIn" gorm:"column:CheckIn;type:timestamp;not null"`
	CheckOut  *time.Time `json:"checkOut" gorm:"column:CheckOut;type:timestamp"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Name"`
}

func (Attendance) TableName() string {
	return "hris_attendance"
}

func AddAttendance(attendance Attendance) error {
	return db.Create(&attendance).Error
}

func UpdateAttendance(attendance Attendance) error {
	return db.Model(Attendance{}).Where(&Attendance{ID: attendance.ID}).Updates(attendance).Error
}

func GetAttendace(user string, date time.Time) Attendance {
	var attendance Attendance
	db.Model(Attendance{}).Where(&Attendance{Username: user, Date: date}).First(&attendance)
	return attendance
}

func GetAttendaces(cond *Attendance) []Attendance {
	var attendances []Attendance
	db.Model(Attendance{}).Where(cond).Find(&attendances)
	return attendances
}
