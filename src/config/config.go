package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
	SSLMode  bool
}

//GetConfig icin sad sa
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "kandula.db.elephantsql.com",
			Port:     5432,
			Username: "plrvuppn",
			Password: "DyhDQ6VlBGElGdX-qTJSjB5mR1fAvkrd!",
			Name:     "plrvuppn",
			Charset:  "utf8",
			SSLMode:  false,
		},
	}
}
