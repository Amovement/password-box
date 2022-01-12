package config

type Config struct {
	App struct {
		Mode            string `default:"production"`
		Port            uint16 `default:"50051"`
		Oauth_Login_Url string `default:"localhost"`
		Wechat_port     uint16 `default:"80"`
	}
	MySQL struct {
		Addr     string `default:"localhost:3306"`
		User     string `default:"root"`
		Password string `default:"12345678"`
		DBName   string `default:"turingstar"`
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
