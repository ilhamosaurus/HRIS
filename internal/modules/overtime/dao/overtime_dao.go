package dao

import (
	"context"
	"time"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"gorm.io/gorm"
)

type OvertimeDAO interface {
	Create(context.Context, *model.Overtime) error
	Update(context.Context, *model.Overtime) error
	UpdateStatus(context.Context, int64, types.Status) error
	GetByID(context.Context, int64) (*model.Overtime, error)
	GetOvertimeByDateUsername(context.Context, time.Time, string) (*model.Overtime, error)
	GetOvertimes(context.Context, map[string]any, int, int) ([]*model.Overtime, int64, error)
	Delete(context.Context, int64) error
}

func NewOvertimeDAO(db *gorm.DB) OvertimeDAO {
	return &overtimeDAO{db: db}
}

type overtimeDAO struct {
	db *gorm.DB
}

func (d *overtimeDAO) Create(ctx context.Context, overtime *model.Overtime) error {
	return d.db.WithContext(ctx).Create(overtime).Error
}

func (d *overtimeDAO) Update(ctx context.Context, overtime *model.Overtime) error {
	return d.db.WithContext(ctx).Where(&model.Overtime{ID: overtime.ID}).Updates(overtime).Error
}

func (d *overtimeDAO) UpdateStatus(ctx context.Context, id int64, status types.Status) error {
	return d.db.WithContext(ctx).Where(&model.Overtime{ID: id}).Update("status", status).Error
}

func (d *overtimeDAO) GetByID(ctx context.Context, id int64) (*model.Overtime, error) {
	var overtime model.Overtime
	if err := d.db.WithContext(ctx).Where(&model.Overtime{ID: id}).First(&overtime).Error; err != nil {
		return nil, err
	}
	return &overtime, nil
}

func (d *overtimeDAO) GetOvertimeByDateUsername(ctx context.Context, date time.Time, username string) (*model.Overtime, error) {
	var overtime model.Overtime
	if err := d.db.WithContext(ctx).Where(&model.Overtime{Date: date, Username: username}).First(&overtime).Error; err != nil {
		return nil, err
	}
	return &overtime, nil
}

func (d *overtimeDAO) GetOvertimes(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.Overtime, int64, error) {
	var (
		overtimes []*model.Overtime
		total     int64
	)
	db := d.db.WithContext(ctx).Model(&model.Overtime{})

	for k, v := range query {
		db = db.Where(k, v)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("date DESC").Limit(pageSize).Offset(offset).Find(&overtimes).Error; err != nil {
		return nil, 0, err
	}
	return overtimes, total, nil
}

func (d *overtimeDAO) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Where(&model.Overtime{ID: id}).Delete(&model.Overtime{}).Error
}
