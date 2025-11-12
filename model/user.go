package model

import "time"

type User struct {
	ID        int32     `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"column:Name;type:varchar(48);index:idx_username,unique;not null"`
	Password  string    `json:"password" gorm:"column:Password;type:varchar(255);not null"`
	UserRole  string    `json:"userRole" gorm:"column:UserRole;type:varchar(48);not null"`
	Salary    float64   `json:"rate" gorm:"column:Salary;type:numeric(18,6);not null"`
	Active    bool      `json:"active" gorm:"column:Active;type:bool;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`
}

func (User) TableName() string {
	return "hris_user"
}

func (m *Model) AddUser(user User) error {
	return m.db.Create(&user).Error
}

func (m *Model) UpdateUser(user User) error {
	return m.db.Model(User{}).Where(&User{ID: user.ID}).Updates(user).Error
}

func (m *Model) DeleteUser(username string) error {
	return m.db.Model(User{}).Where(&User{Name: username}).Delete(&User{}).Error
}

func (m *Model) GetUserByUsername(username string) User {
	var user User
	m.db.Model(User{}).Where(&User{Name: username}).First(&user)
	return user
}

func (m *Model) GetUserById(id int32) User {
	var user User
	m.db.Model(User{}).Where(&User{ID: id}).First(&user)
	return user
}

func (m *Model) GetUsers(cond *User) []User {
	var users []User
	m.db.Model(User{}).Where(cond).Find(&users)
	return users
}
