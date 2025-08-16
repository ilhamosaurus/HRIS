package model

import "time"

type User struct {
	ID        int32     `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"column:Username;type:varchar(48);index:idx_username,unique;not null"`
	Password  string    `json:"password" gorm:"column:Password;type:varchar(255);not null"`
	UserRole  string    `json:"userRole" gorm:"column:UserRole;type:varchar(48);not null"`
	Salary    float64   `json:"rate" gorm:"column:Rate;type:numeric(18,6);not null"`
	Active    bool      `json:"active" gorm:"column:Active;type:bool;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:CreatedAt;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:UpdatedAt;type:timestamp;autoUpdateTime"`
}

func (User) TableName() string {
	return "hris_user"
}

func AddUser(user User) error {
	return db.Create(&user).Error
}

func UpdateUser(user User) error {
	return db.Model(User{}).Where(&User{ID: user.ID}).Updates(user).Error
}

func DeleteUser(username string) error {
	return db.Model(User{}).Where(&User{Username: username}).Delete(&User{}).Error
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Model(User{}).Where(&User{Username: username}).First(&user).Error
	return user, err
}

func GetUserById(id int32) (User, error) {
	var user User
	err := db.Model(User{}).Where(&User{ID: id}).First(&user).Error
	return user, err
}

func GetUsers(cond *User) ([]User, error) {
	var users []User
	err := db.Model(User{}).Where(cond).Find(&users).Error
	return users, err
}
