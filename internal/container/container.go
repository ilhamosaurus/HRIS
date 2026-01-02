package container

import (
	userdao "github.com/ilhamosaurus/HRIS/internal/modules/user/dao"
	userhandler "github.com/ilhamosaurus/HRIS/internal/modules/user/handler"
	userservice "github.com/ilhamosaurus/HRIS/internal/modules/user/service"

	attendancedao "github.com/ilhamosaurus/HRIS/internal/modules/attendance/dao"
	attendancehandler "github.com/ilhamosaurus/HRIS/internal/modules/attendance/handler"
	attendanceservice "github.com/ilhamosaurus/HRIS/internal/modules/attendance/service"

	useractivitydao "github.com/ilhamosaurus/HRIS/internal/modules/userActivity/dao"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"gorm.io/gorm"
)

type Container struct {
	DB     *gorm.DB
	Hasher *util.Hasher

	UserDAO         userdao.UserDAO
	AttendanceDAO   attendancedao.AttendanceDAO
	UserActivityDAO useractivitydao.UserActivityDAO

	AuthService       userservice.AuthService
	UserService       userservice.UserService
	AttendanceService attendanceservice.AttendanceService

	AuthHandler       userhandler.AuthHandler
	UserHandler       userhandler.UserHandler
	AttendanceHandler attendancehandler.AttendanceHandler
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
	c.UserActivityDAO = useractivitydao.NewUserActivityDAO(c.DB)
	c.AttendanceDAO = attendancedao.NewAttendanceDAO(c.DB)
}

func (c *Container) initService() {
	c.AuthService = userservice.NewAuthService(c.UserDAO, c.Hasher)
	c.UserService = userservice.NewUserService(c.UserDAO, c.Hasher)
	c.AttendanceService = attendanceservice.NewAttendanceService(c.AttendanceDAO)
}

func (c *Container) initHandler() {
	c.AuthHandler = userhandler.NewAuthHandler(c.AuthService)
	c.UserHandler = userhandler.NewUserHandler(c.UserService)
	c.AttendanceHandler = attendancehandler.NewAttendanceHandler(c.AttendanceService)
}
