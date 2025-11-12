package model

import "time"

type Reimburse struct {
	ID         int32     `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"column:Name;type:varchar(48);index:idx_name_reimburse,unique;not null"`
	Username   string    `json:"username" gorm:"column:Username;type:varchar(48);not null"`
	Amount     float64   `json:"amount" gorm:"column:Amount;type:numeric(18,6);not null"`
	Reason     *string   `json:"reason" gorm:"column:Reason;type:varchar(255)"`
	Attachment []byte    `json:"attachment" gorm:"column:Attachment"`
	Status     string    `json:"status" gorm:"column:Status;type:varchar(48);not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Name"`
}

func (Reimburse) TableName() string {
	return "hris_reimburse"
}

func (m *Model) AddReimburse(reimburse Reimburse) error {
	return m.db.Create(&reimburse).Error
}

func (m *Model) UpdateReimburse(reimburse Reimburse) error {
	return m.db.Model(Reimburse{}).Where(&Reimburse{ID: reimburse.ID}).Updates(reimburse).Error
}

func (m *Model) GetReimburseById(id int32) (Reimburse, error) {
	var reimburse Reimburse
	err := m.db.Model(Reimburse{}).Where(&Reimburse{ID: id}).First(&reimburse).Error
	return reimburse, err
}

func (m *Model) GetReimburseByName(name string) (Reimburse, error) {
	var reimburse Reimburse
	err := m.db.Model(Reimburse{}).Where(&Reimburse{Name: name}).First(&reimburse).Error
	return reimburse, err
}

func (m *Model) GetReimburses(cond *Reimburse) ([]Reimburse, error) {
	var reimburses []Reimburse
	err := m.db.Model(Reimburse{}).Where(cond).Find(&reimburses).Error
	return reimburses, err
}

func (m *Model) DeleteReimburse(id int32) error {
	return m.db.Model(Reimburse{}).Where(&Reimburse{ID: id}).Delete(&Reimburse{}).Error
}
