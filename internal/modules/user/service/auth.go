package service

import (
	"context"
	"errors"

	"github.com/ilhamosaurus/HRIS/internal/dto"
	"github.com/ilhamosaurus/HRIS/internal/modules/user/dao"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	Login(context.Context, *dto.LoginRequest) (string, error)
}

func NewAuthService(userDAO dao.UserDAO, hasher *util.Hasher) AuthService {
	return &authService{
		userDAO: userDAO,
		hasher:  hasher,
	}
}

type authService struct {
	userDAO dao.UserDAO
	hasher  *util.Hasher
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (string, error) {
	user, err := s.userDAO.GetByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrInvalidCredentials
		}
	}

	if !s.hasher.VerifySHAHash(req.Password, user.Password) || user.Name != req.Username {
		return "", ErrInvalidCredentials
	}

	return util.GeneratoeJWTToken(user.ID, user.Name, user.UserRole)
}
