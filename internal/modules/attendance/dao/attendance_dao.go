package dao

import (
	"context"
	"time"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"gorm.io/gorm"
)

type AttendanceDAO interface {
	Create(context.Context, *model.Attendance) error
	Update(context.Context, *model.Attendance) error
	GetByID(context.Context, int64) (*model.Attendance, error)
	GetByDateUsername(context.Context, string, time.Time) (*model.Attendance, error)
	GetAttendaces(context.Context, map[string]any, int, int) ([]*model.Attendance, int64, error)
	Delete(context.Context, int64) error
}

func NewAttendanceDAO(db *gorm.DB) AttendanceDAO {
	return &attendanceDAO{db: db}
}

type attendanceDAO struct {
	db *gorm.DB
}

func (d *attendanceDAO) Create(ctx context.Context, attendance *model.Attendance) error {
	return d.db.WithContext(ctx).Create(attendance).Error
}

func (d *attendanceDAO) Update(ctx context.Context, attendance *model.Attendance) error {
	return d.db.WithContext(ctx).Where(&model.Attendance{ID: attendance.ID}).Updates(attendance).Error
}

func (d *attendanceDAO) GetByID(ctx context.Context, id int64) (*model.Attendance, error) {
	var attendance model.Attendance
	if err := d.db.WithContext(ctx).Where(&model.Attendance{ID: id}).First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (d *attendanceDAO) GetByDateUsername(ctx context.Context, username string, date time.Time) (*model.Attendance, error) {
	var attendance model.Attendance
	if err := d.db.WithContext(ctx).Where(&model.Attendance{Username: username, Date: date}).First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (d *attendanceDAO) GetAttendaces(ctx context.Context, query map[string]any, page, pageSize int) ([]*model.Attendance, int64, error) {
	var (
		attendances []*model.Attendance
		total       int64
	)

	db := d.db.WithContext(ctx).Model(&model.Attendance{})

	for k, v := range query {
		db = db.Where(k, v)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("date DESC").Limit(pageSize).Offset(offset).Find(&attendances).Error; err != nil {
		return nil, 0, err
	}

	return attendances, total, nil
}

func (d *attendanceDAO) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Where(&model.Attendance{ID: id}).Delete(&model.Attendance{}).Error
}
