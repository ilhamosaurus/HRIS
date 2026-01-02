package service

import (
	"context"
	"time"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/model"
	"github.com/ilhamosaurus/HRIS/internal/modules/attendance/dao"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
)

type AttendanceService interface {
	CheckIn(context.Context, time.Time, time.Time) error
	CheckOut(context.Context, time.Time, time.Time) error
	GetAttendances(context.Context, *dto.AttendancesQuery) ([]*model.Attendance, int64, error)
	Delete(context.Context, int64) error
}

func NewAttendanceService(attendanceDAO dao.AttendanceDAO) AttendanceService {
	return &attendanceService{attendanceDAO: attendanceDAO}
}

type attendanceService struct {
	attendanceDAO dao.AttendanceDAO
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

	return s.attendanceDAO.Create(ctx, &attendance)
}

func (s *attendanceService) CheckOut(ctx context.Context, date, checkOut time.Time) error {
	username, ok := util.GetUsernameFromCtx(ctx)
	if !ok {
		return types.ErrUsernameNotExist
	}

	attendance, err := s.attendanceDAO.GetByDateUsername(ctx, username, date)
	if err != nil {
		return err
	}

	attendance.CheckOut = &checkOut
	return s.attendanceDAO.Update(ctx, attendance)
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
