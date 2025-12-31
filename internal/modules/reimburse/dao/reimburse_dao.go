package dao

import (
	"context"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"gorm.io/gorm"
)

type ReimburseDAO interface {
	Create(ctx context.Context, reimburse *model.Reimburse) error
	Update(ctx context.Context, reimburse *model.Reimburse) error
	GetByID(ctx context.Context, id int64) (*model.Reimburse, error)
	GetReimburses(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.Reimburse, int64, error)
	Delete(ctx context.Context, id int64) error
}

func NewReimburseDAO(db *gorm.DB) ReimburseDAO {
	return &reimburseDAO{
		db: db,
	}
}

type reimburseDAO struct {
	db *gorm.DB
}

func (d *reimburseDAO) Create(ctx context.Context, reimburse *model.Reimburse) error {
	return d.db.WithContext(ctx).Create(&reimburse).Error
}

func (d *reimburseDAO) Update(ctx context.Context, reimburse *model.Reimburse) error {
	return d.db.WithContext(ctx).Where(&model.Reimburse{ID: reimburse.ID}).Updates(reimburse).Error
}

func (d *reimburseDAO) GetByID(ctx context.Context, id int64) (*model.Reimburse, error) {
	var reimburse model.Reimburse
	if err := d.db.WithContext(ctx).Where(&model.Reimburse{ID: id}).First(&reimburse).Error; err != nil {
		return nil, err
	}
	return &reimburse, nil
}

func (d *reimburseDAO) GetReimburses(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.Reimburse, int64, error) {
	var (
		reimburses []*model.Reimburse
		total      int64
	)
	db := d.db.WithContext(ctx).Model(&model.Reimburse{})

	for key, value := range query {
		db = db.Where(key, value)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&reimburses).Error; err != nil {
		return nil, 0, err
	}
	return reimburses, total, nil
}

func (d *reimburseDAO) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Where(&model.Reimburse{ID: id}).Delete(&model.Reimburse{}).Error
}
