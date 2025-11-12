package model

import (
	"encoding/json"
	"os"

	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
)

type MockUser struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Salary   float64 `json:"salary"`
}

func Seed(model *Model) error {
	hasher := util.NewHasher(setting.Server.Secret)

	hashedAdminPassword := hasher.GenerateSHAHash(setting.AdminUser.Password)
	adminUser := User{
		Name:     setting.AdminUser.Username,
		Password: hashedAdminPassword,
		UserRole: types.Admin.String(),
		Salary:   0,
		Active:   true,
	}
	existingAdmin := model.GetUserByUsername(adminUser.Name)
	if existingAdmin.Name != adminUser.Name {
		if err := model.AddUser(adminUser); err != nil {
			return err
		}
	}
	b, err := os.ReadFile("./data/mock_users.json")
	if err != nil {
		return err
	}

	var mockUsers []MockUser
	err = json.Unmarshal(b, &mockUsers)
	if err != nil {
		return err
	}

	tx := model.db.Begin()
	for i := range mockUsers {
		user := User{
			Name:     mockUsers[i].Username,
			Password: mockUsers[i].Password,
			UserRole: types.Employee.String(),
			Salary:   mockUsers[i].Salary,
			Active:   true,
		}

		existingUser := model.GetUserByUsername(user.Name)
		if existingUser.Name == user.Name {
			continue
		}

		hashedPassword := hasher.GenerateSHAHash(user.Password)
		user.Password = hashedPassword
		err = tx.Create(&user).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
