package service

import (
	"context"
	"strconv"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/model"
	"github.com/ilhamosaurus/HRIS/internal/modules/user/dao"
	"github.com/ilhamosaurus/HRIS/pkg/util"
)

type UserService interface {
	Create(context.Context, *dto.CreateUserRequest) error
	Update(context.Context, *dto.UpdateUserRequest) error
	Delete(context.Context, int64) error
	GetByUsername(context.Context, string) (*model.User, error)
	GetByID(context.Context, int64) (*dto.UserResponse, error)
	List(context.Context, *dto.UserQuery) (*dto.UserListResponse, error)
}

func NewUserService(userDAO dao.UserDAO, hasher *util.Hasher) UserService {
	return &userService{
		userDAO: userDAO,
		hasher:  hasher,
	}
}

type userService struct {
	userDAO dao.UserDAO
	hasher  *util.Hasher
}

func (s *userService) Create(ctx context.Context, req *dto.CreateUserRequest) error {
	user := s.toModelUser(req)
	return s.userDAO.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, req *dto.UpdateUserRequest) error {
	user := s.toModelUser(req)
	return s.userDAO.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	return s.userDAO.Delete(ctx, id)
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.userDAO.GetByUsername(ctx, username)
}

func (s *userService) GetByID(ctx context.Context, id int64) (*dto.UserResponse, error) {
	user, err := s.userDAO.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toUserResponse(user), nil
}

func (s *userService) List(ctx context.Context, reqQuery *dto.UserQuery) (*dto.UserListResponse, error) {
	query := make(map[string]any)
	if reqQuery.Role != "" {
		query["user_role = ?"] = reqQuery.Role
	}
	if reqQuery.Active != "" {
		active, err := strconv.ParseBool(reqQuery.Active)
		if err != nil {
			return nil, err
		}
		query["active = ?"] = active
	}
	users, total, err := s.userDAO.List(ctx, query, reqQuery.Page, reqQuery.PageSize)
	if err != nil {
		return nil, err
	}

	userResponses := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = s.toUserResponse(user)
	}
	return &dto.UserListResponse{
		Users:      userResponses,
		TotalCount: total,
		PageSize:   reqQuery.PageSize,
		Page:       reqQuery.Page,
	}, nil
}

func (s *userService) toModelUser(data any) *model.User {
	switch v := data.(type) {
	case *dto.CreateUserRequest:
		hashedPassword := s.hasher.GenerateSHAHash(v.Password)
		return &model.User{
			Name:     v.Username,
			Password: hashedPassword,
			Email:    v.Email,
			UserRole: v.Role,
			Salary:   v.Salary,
			Active:   v.Active,
		}
	case *dto.UpdateUserRequest:
		user := &model.User{
			ID:       v.ID,
			Name:     v.Username,
			Email:    v.Email,
			UserRole: v.Role,
			Salary:   v.Salary,
			Active:   v.Active,
		}
		if v.Password != "" {
			hashedPassword := s.hasher.GenerateSHAHash(v.Password)
			user.Password = hashedPassword
		}
		return user
	default:
		return nil
	}
}

func (s *userService) toUserResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Name,
		Email:    user.Email,
		Role:     user.UserRole,
		Salary:   user.Salary,
		Active:   user.Active,
	}
}
