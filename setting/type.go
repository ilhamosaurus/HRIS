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
	Port int
}
