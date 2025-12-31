package dao

import (
	"context"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"gorm.io/gorm"
)

type UserDAO interface {
	Create(context.Context, *model.User) error
	Update(context.Context, *model.User) error
	Delete(context.Context, int64) error
	GetByUsername(context.Context, string) (*model.User, error)
	GetById(context.Context, int64) (*model.User, error)
	List(context.Context, map[string]any, int, int) ([]*model.User, int64, error)
}

type userDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAO{db: db}
}

func (d *userDAO) Create(ctx context.Context, user *model.User) error {
	return d.db.WithContext(ctx).Create(user).Error
}

func (d *userDAO) Update(ctx context.Context, user *model.User) error {
	return d.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (d *userDAO) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error
}

func (d *userDAO) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	return &user, d.db.WithContext(ctx).Model(&model.User{}).Where("name = ?", username).First(&user).Error
}

func (d *userDAO) GetById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	return &user, d.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error
}

func (d *userDAO) List(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User

	db := d.db.WithContext(ctx).Model(&model.User{})

	for key, value := range query {
		db = db.Where(key, value)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, total, err
	}
	return users, total, nil
}
