package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	Cfg       *ini.File
	Server    ServerCfg
	Database  DatabaseCfg
	AdminUser Administrator
)

func init() {
	var err error
	Cfg, err = ini.Load("./conf/HRIS.cfg")
	if err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}

	LoadServer()
	LoadDatabase()
	LoadAdministrator()
}

func LoadServer() {
	sec, err := Cfg.GetSection("SERVER")
	if err != nil {
		log.Fatalf("failed to get section SERVER: %v", err)
	}

	Server.Port = sec.Key("Port").MustInt(8080)
	Server.Secret = sec.Key("Secret").MustString("@!@#!#@!#!#!")
	Server.JWTSecret = sec.Key("JWTSecret").MustString("@!@#!#@!#!#!")
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("DATABASE")
	if err != nil {
		log.Fatalf("failed to get section DATABASE: %v", err)
	}

	Database.Type = sec.Key("Type").MustString("postgres")
	Database.Host = sec.Key("Host").MustString("localhost")
	Database.Port = sec.Key("Port").MustInt(5432)
	Database.User = sec.Key("User").MustString("postgres")
	Database.Pass = sec.Key("Password").MustString("@@!#!@@#@!")
	Database.Name = sec.Key("Name").MustString("hris")
	Database.DSN = sec.Key("DSN").MustString("hris.db")
}

func LoadAdministrator() {
	sec, err := Cfg.GetSection("ADMINISTRATOR")
	if err != nil {
		log.Fatalf("failed to get section ADMINISTRATOR: %v", err)
	}

	AdminUser.Username = sec.Key("Username").MustString("admin")
	AdminUser.Password = sec.Key("Password").MustString("@!@#!#@!@#")
}
