package setting

type DatabaseCfg struct {
	Type string
	Host string
	Port int
	User string
	Pass string
	Name string
}

type ServerCfg struct {
	Port      int
	Secret    string
	JWTSecret string
}

type Administrator struct {
	Username string
	Password string
}
