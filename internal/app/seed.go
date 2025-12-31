package app

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"github.com/ilhamosaurus/HRIS/internal/modules/user/dao"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"gorm.io/gorm"
)

type MockUser struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Salary   float64 `json:"salary"`
}

func Seed(db *gorm.DB, hasher *util.Hasher, userDAO dao.UserDAO) error {
	ctx := context.Background()
	hashedAdminPassword := hasher.GenerateSHAHash(setting.AdminUser.Password)
	adminUser := model.User{
		Name:     setting.AdminUser.Username,
		Password: hashedAdminPassword,
		UserRole: types.Admin.String(),
		Salary:   0,
		Active:   true,
	}
	existingAdmin, err := userDAO.GetByUsername(ctx, adminUser.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingAdmin.Name != adminUser.Name {
		if err := userDAO.Create(ctx, &adminUser); err != nil {
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

	tx := db.Begin()
	for _, mockUser := range mockUsers {
		user := model.User{
			Name:     mockUser.Username,
			Password: mockUser.Password,
			Email:    mockUser.Username + "@example.com",
			UserRole: types.Employee.String(),
			Salary:   mockUser.Salary,
			Active:   true,
		}

		existingUser, err := userDAO.GetByUsername(ctx, user.Name)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}
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
