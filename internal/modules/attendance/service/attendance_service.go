package service

import (
	"context"
	"errors"
	"time"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/model"
	"github.com/ilhamosaurus/HRIS/internal/modules/attendance/dao"
	overtimedao "github.com/ilhamosaurus/HRIS/internal/modules/overtime/dao"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"gorm.io/gorm"
)

type AttendanceService interface {
	CheckIn(context.Context, time.Time, time.Time) error
	CheckOut(context.Context, time.Time, time.Time) error
	GetAttendances(context.Context, *dto.AttendancesQuery) ([]*model.Attendance, int64, error)
	Delete(context.Context, int64) error
}

func NewAttendanceService(attendanceDAO dao.AttendanceDAO, overtimeDAO overtimedao.OvertimeDAO) AttendanceService {
	return &attendanceService{
		attendanceDAO: attendanceDAO,
		overtimeDAO:   overtimeDAO,
	}
}

type attendanceService struct {
	attendanceDAO dao.AttendanceDAO
	overtimeDAO   overtimedao.OvertimeDAO
}

func (s *attendanceService) CheckIn(ctx context.Context, date, checkIn time.Time) error {
	attendance := model.Attendance{
		Date:    date,
		CheckIn: checkIn,
	}
	username, ok := util.GetUsernameFromCtx(ctx)
	if ok {
		attendance.Username = username
	}

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		_, err := s.overtimeDAO.GetOvertimeByDateUsername(ctx, date, username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return types.ErrNoOvertimeApproved
			}
			return err
		}
	}

	existingAttendance, err := s.attendanceDAO.GetByDateUsername(ctx, attendance.Username, date)
	if err == nil && existingAttendance != nil {
		existingAttendance = &attendance
		return s.attendanceDAO.Update(ctx, existingAttendance)
	}

	return s.attendanceDAO.Create(ctx, &attendance)
}

func (s *attendanceService) CheckOut(ctx context.Context, date, checkOut time.Time) error {
	var (
		overtime *model.Overtime
		err      error
	)
	username, ok := util.GetUsernameFromCtx(ctx)
	if !ok {
		return types.ErrUsernameNotExist
	}

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		overtime, err = s.overtimeDAO.GetOvertimeByDateUsername(ctx, date, username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return types.ErrNoOvertimeApproved
			}
			return err
		}
	}

	attendance, err := s.attendanceDAO.GetByDateUsername(ctx, username, date)
	if err != nil {
		return err
	}

	attendance.CheckOut = &checkOut
	if err := s.attendanceDAO.Update(ctx, attendance); err != nil {
		return err
	}

	if overtime != nil && overtime.Date.Equal(attendance.Date) && overtime.Username == attendance.Username {
		overtime.StartTime = attendance.CheckIn
		overtime.EndTime = checkOut
		overtime.Hours = checkOut.Sub(attendance.CheckIn).Hours()
		overtime.Status = types.Done
		s.overtimeDAO.Update(ctx, overtime)
	}

	return nil
}

func (s *attendanceService) GetAttendances(ctx context.Context, req *dto.AttendancesQuery) ([]*model.Attendance, int64, error) {
	query := make(map[string]any)
	if req.Username != "" {
		query["username = ?"] = req.Username
	}
	if req.Date != "" {
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return nil, 0, err
		}
		query["date = ?"] = date
	}

	userRole, ok := util.GetRoleFromCtx(ctx)
	if ok && userRole != types.Admin {
		username, ok := util.GetUsernameFromCtx(ctx)
		if ok {
			query["username = ?"] = username
		}
	}

	return s.attendanceDAO.GetAttendaces(ctx, query, req.Page, req.PageSize)
}

func (s *attendanceService) Delete(ctx context.Context, id int64) error {
	userRole, ok := util.GetRoleFromCtx(ctx)
	if !ok || userRole != types.Admin {
		return types.ErrUnauthorized
	}

	return s.attendanceDAO.Delete(ctx, id)
}
