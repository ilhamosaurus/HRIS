package model

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(48);index:idx_username,unique;not null"`
	Password  string    `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"column:email;type:varchar(48);index:idx_email,unique;not null"`
	UserRole  string    `json:"userRole" gorm:"column:user_role;type:varchar(48);not null"`
	Salary    float64   `json:"rate" gorm:"column:salary;type:numeric(18,6);not null"`
	Active    bool      `json:"active" gorm:"column:active;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}

func (User) TableName() string {
	return "hris_user"
}
