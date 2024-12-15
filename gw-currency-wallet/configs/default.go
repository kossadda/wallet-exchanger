package configs

func NewDefaultConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "5436",
		User:     "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
}
