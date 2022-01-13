package config

type Config struct {
	App struct {
		Port        uint16 `default:"50051"`
		Session_key string `default:"pwdkey"`
	}
	MySQL struct {
		Addr     string `default:"localhost:3306"`
		User     string `default:"root"`
		Password string `default:"12345678"`
		DBName   string `default:"pwdbox"`
	}
	Redis struct {
		Addr     string `default:"localhost:6379"`
		Password string `default:""`
		DB       uint16 `default:"0"`
		Prefix   string `default:""`
	}
}

var cfg Config

func GetConfig() *Config {
	return &cfg
}
