package model

import (
	"fmt"
	"log"

	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	db *gorm.DB
}

func NewModel() (*Model, error) {
	var dialector gorm.Dialector
	var err error

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
	default:
		log.Fatalf("database type %s is not supported", setting.Database.Type)
	}

	opts := gorm.Config{
		TranslateError: true,
	}

	db, err = gorm.Open(dialector, &opts)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{}, &Attendance{}, &Overtime{}, &Reimburse{}, &Payslip{}, &UserActivity{})
	if err != nil {
		return nil, err
	}

	model := &Model{db: db}
	if err := Seed(model); err != nil {
		return nil, err
	}
	return model, nil
}
