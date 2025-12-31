package model

import "time"

type Reimburse struct {
	ID         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(48);index:idx_name_reimburse,unique;not null"`
	Username   string    `json:"username" gorm:"column:username;type:varchar(48);not null"`
	Amount     float64   `json:"amount" gorm:"column:amount;type:numeric(18,6);not null"`
	Reason     string    `json:"reason" gorm:"column:reason;type:varchar(255)"`
	Attachment []byte    `json:"attachment" gorm:"column:attachment"`
	Status     string    `json:"status" gorm:"column:status;type:varchar(48);not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;type:timestamp;autoUpdateTime"`

	User User `json:"user" gorm:"foreignKey:Username;references:Name"`
}

func (Reimburse) TableName() string {
	return "hris_reimburse"
}
