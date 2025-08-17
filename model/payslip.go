package model

import "time"

type Payslip struct {
	Date          time.Time  `json:"date" gorm:"column:Date;type:date;primaryKey;not null"`
	Username      string     `json:"username" gorm:"column:Username;primaryKey;not null"`
	Attendance    int32      `json:"attendate" gorm:"column:Attendate;not null"`
	BaseSalary    float64    `json:"baseSalary" gorm:"column:BaseSalary;type:numeric(18,6);not null"`
	OvertimeHours *float64   `json:"overtimeHours" gorm:"column:OvertimeHours;type:numeric(18,6)"`
	OvertimePay   *float64   `json:"overtimePay" gorm:"column:OvertimePay;type:numeric(18,6)"`
	Reimburse     *float64   `json:"reimburse" gorm:"column:Reimburse;type:numeric(18,6)"`
	TakeHomePay   float64    `json:"takeHomePay" gorm:"column:TakeHomePay;type:numeric(18,6);not null"`
	Processed     bool       `json:"processed" gorm:"column:Processed;type:bool;not null"`
	ProcessedAt   *time.Time `json:"processedAt" gorm:"column:ProcessedAt;type:timestamp"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Name"`
}

func (Payslip) TableName() string {
	return "hris_payslip"
}

func AddPayslip(payslip Payslip) error {
	return db.Create(&payslip).Error
}

func UpdatePayslip(payslip Payslip) error {
	return db.Model(Payslip{}).Where(&Payslip{Date: payslip.Date, Username: payslip.Username}).Updates(payslip).Error
}

func GetPayslipById(username string, date time.Time) (Payslip, error) {
	var payslip Payslip
	err := db.Model(Payslip{}).Where(&Payslip{Username: username, Date: date}).First(&payslip).Error
	return payslip, err
}

func GetPayslips(cond *Payslip) ([]Payslip, error) {
	var payslips []Payslip
	err := db.Model(Payslip{}).Where(cond).Find(&payslips).Error
	return payslips, err
}

func DeletePayslip(username string, date time.Time) error {
	return db.Model(Payslip{}).Where(&Payslip{Username: username, Date: date}).Delete(&Payslip{}).Error
}
