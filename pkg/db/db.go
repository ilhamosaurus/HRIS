package db

import (
	"fmt"
	"log"

	"github.com/ilhamosaurus/HRIS/internal/model"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch setting.Database.Type {
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			setting.Database.Host,
			setting.Database.Port,
			setting.Database.User,
			setting.Database.Pass,
			setting.Database.Name,
		)
		log.Printf("dsn: %s", dsn)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(setting.Database.DSN)
	default:
		log.Fatalf("database type %s is not supported", setting.Database.Type)
	}

	opts := gorm.Config{
		TranslateError: true,
	}

	db, err := gorm.Open(dialector, &opts)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Attendance{}, &model.Overtime{}, &model.Reimburse{}, &model.Payslip{}, &model.UserActivity{})
	if err != nil {
		return nil, err
	}

	// if err := Seed(model); err != nil {
	// 	return nil, err
	// }
	return db, nil
}
