package model

import (
	"fmt"
	"log"

	"github.com/ilhamosaurus/HRIS/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var dialector gorm.Dialector

	switch setting.Database.Type {
	case "postgres":
		dialector = postgres.Open(fmt.Sprintf("postgress://%s:%s@%s:%d/%s?sslmode=disable",
			setting.Database.User,
			setting.Database.Pass,
			setting.Database.Host,
			setting.Database.Port,
			setting.Database.Name,
		))
	default:
		log.Fatalf("database type %s is not supported", setting.Database.Type)
	}

	opts := gorm.Config{
		TranslateError: true,
	}

	db, err := gorm.Open(dialector, &opts)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&User{}, &Attendance{}, &Overtime{}, &Reimburse{}, &Payslip{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
