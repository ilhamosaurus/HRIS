package dao

import (
	"context"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"gorm.io/gorm"
)

type PayslipDAO interface {
	Create(ctx context.Context, payslip *model.Payslip) error
	Update(ctx context.Context, payslip *model.Payslip) error
	GetByCode(ctx context.Context, code string) (*model.Payslip, error)
	GetPayslips(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.Payslip, int64, error)
	Delete(ctx context.Context, code string) error
}

func NewPayslipDAO(db *gorm.DB) PayslipDAO {
	return &payslipDAO{
		db: db,
	}
}

type payslipDAO struct {
	db *gorm.DB
}

func (d *payslipDAO) Create(ctx context.Context, payslip *model.Payslip) error {
	return d.db.WithContext(ctx).Create(&payslip).Error
}

func (d *payslipDAO) Update(ctx context.Context, payslip *model.Payslip) error {
	return d.db.WithContext(ctx).Where(&model.Payslip{Code: payslip.Code}).Updates(payslip).Error
}

func (d *payslipDAO) GetByCode(ctx context.Context, code string) (*model.Payslip, error) {
	var payslip model.Payslip
	if err := d.db.WithContext(ctx).Where(&model.Payslip{Code: code}).First(&payslip).Error; err != nil {
		return nil, err
	}
	return &payslip, nil
}

func (d *payslipDAO) GetPayslips(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.Payslip, int64, error) {
	var (
		payslips []*model.Payslip
		total    int64
	)
	db := d.db.WithContext(ctx).Model(&model.Payslip{})

	for key, value := range query {
		db = db.Where(key, value)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("date DESC").Limit(pageSize).Offset(offset).Find(&payslips).Error; err != nil {
		return nil, 0, err
	}
	return payslips, total, nil
}

func (d *payslipDAO) Delete(ctx context.Context, code string) error {
	return d.db.WithContext(ctx).Where(&model.Payslip{Code: code}).Delete(&model.Payslip{}).Error
}
