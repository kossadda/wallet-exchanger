package configs

func NewDefaultConfig() *Config {
	return &Config{
		DBHost:     "localhost",
		DBPort:     "5436",
		DBUser:     "postgres",
		DBPassword: "qwerty",
		DBName:     "postgres",
		DBSSLMode:  "disable",
		ServerPort: "8000",
	}
}
