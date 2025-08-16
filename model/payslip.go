package model

import "time"

type Payslip struct {
	ID            int32     `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Code          string    `json:"code" gorm:"column:Code;type:varchar(48);index:idx_code_payslip,unique;not null"`
	Username      string    `json:"username" gorm:"column:Username;type:varchar(48);not null"`
	Attendance    int32     `json:"attendate" gorm:"column:Attendate;not null"`
	BaseSalary    float64   `json:"baseSalary" gorm:"column:BaseSalary;type:numeric(18,6);not null"`
	OvertimeHours *float64  `json:"overtimeHours" gorm:"column:OvertimeHours;type:numeric(18,6)"`
	OvertimePay   *float64  `json:"overtimePay" gorm:"column:OvertimePay;type:numeric(18,6)"`
	Reimburse     *float64  `json:"reimburse" gorm:"column:Reimburse;type:numeric(18,6)"`
	TakeHomePay   float64   `json:"takeHomePay" gorm:"column:TakeHomePay;type:numeric(18,6);not null"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Username"`
}

func (Payslip) TableName() string {
	return "hris_payslip"
}

func AddPayslip(payslip Payslip) error {
	return db.Create(&payslip).Error
}

func UpdatePayslip(payslip Payslip) error {
	return db.Model(Payslip{}).Where(&Payslip{ID: payslip.ID}).Updates(payslip).Error
}

func GetPayslipById(id int32) (Payslip, error) {
	var payslip Payslip
	err := db.Model(Payslip{}).Where(&Payslip{ID: id}).First(&payslip).Error
	return payslip, err
}

func GetPayslip(code string) (Payslip, error) {
	var payslip Payslip
	err := db.Model(Payslip{}).Where(&Payslip{Code: code}).First(&payslip).Error
	return payslip, err
}

func GetPayslips(cond *Payslip) ([]Payslip, error) {
	var payslips []Payslip
	err := db.Model(Payslip{}).Where(cond).Find(&payslips).Error
	return payslips, err
}

func DeletePayslip(id int32) error {
	return db.Model(Payslip{}).Where(&Payslip{ID: id}).Delete(&Payslip{}).Error
}
