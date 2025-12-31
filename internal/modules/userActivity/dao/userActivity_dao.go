package dao

import (
	"context"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"gorm.io/gorm"
)

type UserActivityDAO interface {
	Create(ctx context.Context, log *model.UserActivity) error
	GetUserActivities(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.UserActivity, int64, error)
}

func NewUserActivityDAO(db *gorm.DB) UserActivityDAO {
	return &userActivityDAO{
		db: db,
	}
}

type userActivityDAO struct {
	db *gorm.DB
}

func (d *userActivityDAO) Create(ctx context.Context, log *model.UserActivity) error {
	return d.db.WithContext(ctx).Create(log).Error
}

func (d *userActivityDAO) GetUserActivities(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.UserActivity, int64, error) {
	var (
		logs  []*model.UserActivity
		total int64
	)
	db := d.db.WithContext(ctx).Model(&model.UserActivity{})

	for key, value := range query {
		db = db.Where(key, value)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("time DESC").Limit(pageSize).Offset(offset).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
