package container

import (
	userdao "github.com/ilhamosaurus/HRIS/internal/modules/user/dao"
	userhandler "github.com/ilhamosaurus/HRIS/internal/modules/user/handler"
	userservice "github.com/ilhamosaurus/HRIS/internal/modules/user/service"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"gorm.io/gorm"
)

type Container struct {
	DB     *gorm.DB
	Hasher *util.Hasher

	UserDAO userdao.UserDAO

	AuthService userservice.AuthService
	UserService userservice.UserService

	AuthHandler userhandler.AuthHandler
	UserHandler userhandler.UserHandler
}

func NewContainer(db *gorm.DB, hasher *util.Hasher) (*Container, error) {
	c := &Container{
		DB:     db,
		Hasher: hasher,
	}

	c.initDAO()
	c.initService()
	c.initHandler()

	return c, nil
}

func (c *Container) initDAO() {
	c.UserDAO = userdao.NewUserDAO(c.DB)
}

func (c *Container) initService() {
	c.AuthService = userservice.NewAuthService(c.UserDAO, c.Hasher)
	c.UserService = userservice.NewUserService(c.UserDAO, c.Hasher)
}

func (c *Container) initHandler() {
	c.AuthHandler = userhandler.NewAuthHandler(c.AuthService)
	c.UserHandler = userhandler.NewUserHandler(c.UserService)
}
