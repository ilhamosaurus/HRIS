package service

import (
	"context"
	"errors"
	"time"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/model"
	"github.com/ilhamosaurus/HRIS/internal/modules/overtime/dao"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"gorm.io/gorm"
)

type OvertimeService interface {
	Create(context.Context, *dto.CreateOvertimeRequest) (*dto.OvertimeResponse, error)
	Update(context.Context, int64, *dto.UpdateOvertimeRequest) (*dto.OvertimeResponse, error)
	GetByID(context.Context, int64) (*dto.OvertimeResponse, error)
	GetOvertimes(context.Context, *dto.OvertimeQuery) (*dto.OvertimesResponse, error)
	Delete(context.Context, int64) error
	Submit(context.Context, int64) (*dto.OvertimeResponse, error)
	ProcessApproval(context.Context, int64, *dto.ApprovalRequest) (*dto.OvertimeResponse, error)
}

func NewOvertimeService(overtimeDAO dao.OvertimeDAO) OvertimeService {
	return &overtimeService{
		overtimeDAO: overtimeDAO,
	}
}

type overtimeService struct {
	overtimeDAO dao.OvertimeDAO
}

func (s *overtimeService) Create(ctx context.Context, req *dto.CreateOvertimeRequest) (*dto.OvertimeResponse, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, types.ErrDateFormat
	}

	startTime, err := time.Parse(time.RFC3339Nano, req.StartTime)
	if err != nil {
		return nil, types.ErrDateTimeFormat
	}

	endTime, err := time.Parse(time.RFC3339Nano, req.EndTime)
	if err != nil {
		return nil, types.ErrDateTimeFormat
	}

	if endTime.Before(startTime) {
		return nil, errors.New("end time must be later then start time")
	}

	overtime, err := s.overtimeDAO.GetOvertimeByDateUsername(ctx, date, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if overtime == nil {
		overtime = &model.Overtime{
			Date:        date,
			Username:    req.Username,
			Description: req.Description,
			Status:      types.Draft,
		}
	}

	if overtime.Status != types.Draft {
		return s.convertToResponse(overtime), nil
	}
	overtime.StartTime = startTime
	overtime.EndTime = endTime
	overtime.Hours = overtime.CalculateHours()

	userRole, _ := util.GetRoleFromCtx(ctx)
	if userRole != types.Admin {
		username, _ := util.GetUsernameFromCtx(ctx)
		overtime.Username = username
	}

	err = s.overtimeDAO.Create(ctx, overtime)
	if err != nil {
		return nil, err
	}

	return s.convertToResponse(overtime), nil
}

func (s *overtimeService) Update(ctx context.Context, id int64, req *dto.UpdateOvertimeRequest) (*dto.OvertimeResponse, error) {
	overtime, err := s.overtimeDAO.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Date != "" {
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return nil, types.ErrDateFormat
		}
		overtime.Date = date
	}

	if req.StartTime != "" {
		startTime, err := time.Parse(time.RFC3339Nano, req.StartTime)
		if err != nil {
			return nil, types.ErrDateTimeFormat
		}
		overtime.StartTime = startTime
	}

	if req.EndTime != "" {
		endTime, err := time.Parse(time.RFC3339Nano, req.EndTime)
		if err != nil {
			return nil, types.ErrDateTimeFormat
		}
		overtime.EndTime = endTime
	}

	if overtime.EndTime.Before(overtime.StartTime) {
		return nil, errors.New("end time must be later then start time")
	}

	overtime.Hours = overtime.CalculateHours()
	if req.Username != "" {
		userRole, _ := util.GetRoleFromCtx(ctx)
		if userRole != types.Admin {
			username, _ := util.GetUsernameFromCtx(ctx)
			overtime.Username = username
		} else {
			overtime.Username = req.Username
		}
	}

	if req.Description != "" {
		overtime.Description = req.Description
	}

	if err := s.overtimeDAO.Update(ctx, overtime); err != nil {
		return nil, err
	}
	return s.convertToResponse(overtime), nil
}

func (s *overtimeService) GetByID(ctx context.Context, id int64) (*dto.OvertimeResponse, error) {
	overtime, err := s.overtimeDAO.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.convertToResponse(overtime), nil
}

func (s *overtimeService) GetOvertimes(ctx context.Context, req *dto.OvertimeQuery) (*dto.OvertimesResponse, error) {
	query := make(map[string]any)

	userRole, _ := util.GetRoleFromCtx(ctx)
	if userRole != types.Admin {
		username, _ := util.GetUsernameFromCtx(ctx)
		query["username = ?"] = username
	} else if req.Username != "" {
		query["username LIKE ?"] = "%" + req.Username + "%"
	}

	if req.Date != "" {
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return nil, types.ErrDateFormat
		}
		query["date = ?"] = date
	}

	if req.Status != "" {
		query["status = ?"] = types.StringToStatus(req.Status)
	}

	if req.Approval != "" {
		query["approval LIKE ?"] = "%" + req.Approval + "%"
	}

	overtimes, total, err := s.overtimeDAO.GetOvertimes(ctx, query, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	result := dto.OvertimesResponse{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	result.List = make([]*dto.OvertimeResponse, len(overtimes))
	for i, overtime := range overtimes {
		result.List[i] = s.convertToResponse(overtime)
	}

	return &result, nil
}

func (s *overtimeService) Delete(ctx context.Context, id int64) error {
	overtime, err := s.overtimeDAO.GetByID(ctx, id)
	if err != nil {
		return err
	}

	userRole, _ := util.GetRoleFromCtx(ctx)
	username, _ := util.GetUsernameFromCtx(ctx)
	if userRole != types.Admin && username != overtime.Username {
		return types.ErrUnauthorized
	}

	return s.overtimeDAO.Delete(ctx, id)
}

func (s *overtimeService) Submit(ctx context.Context, id int64) (*dto.OvertimeResponse, error) {
	overtime, err := s.overtimeDAO.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userRole, _ := util.GetRoleFromCtx(ctx)
	username, _ := util.GetUsernameFromCtx(ctx)
	if (userRole != types.Admin && username != overtime.Username) || overtime.Status != types.Draft {
		return nil, types.ErrUnauthorized
	}

	if err := s.overtimeDAO.UpdateStatus(ctx, id, types.Submitted); err != nil {
		return nil, err
	}

	overtime.Status = types.Submitted
	return s.convertToResponse(overtime), nil
}

func (s *overtimeService) ProcessApproval(ctx context.Context, id int64, req *dto.ApprovalRequest) (*dto.OvertimeResponse, error) {
	overtime, err := s.overtimeDAO.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userRole, _ := util.GetRoleFromCtx(ctx)
	if userRole != types.Admin || overtime.Status != types.Submitted {
		return nil, types.ErrUnauthorized
	}

	username, _ := util.GetUsernameFromCtx(ctx)
	overtime.Approval = username
	overtime.Status = types.StringToStatus(req.Status)
	if err := s.overtimeDAO.Update(ctx, overtime); err != nil {
		return nil, err
	}

	return s.convertToResponse(overtime), nil
}

func (s *overtimeService) convertToResponse(overtime *model.Overtime) *dto.OvertimeResponse {
	if overtime == nil {
		return nil
	}

	return &dto.OvertimeResponse{
		ID:          overtime.ID,
		Username:    overtime.Username,
		Date:        overtime.Date.Format("2006-01-02"),
		StartTime:   overtime.StartTime.Format(time.RFC3339Nano),
		EndTime:     overtime.EndTime.Format(time.RFC3339Nano),
		Hours:       overtime.Hours,
		Description: overtime.Description,
		Status:      overtime.Status.String(),
		Approval:    overtime.Approval,
	}
}
