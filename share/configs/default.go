package configs

func NewDefaultConfigDB() *ConfigDB {
	return &ConfigDB{
		DBHost:     "localhost",
		DBPort:     "5436",
		DBUser:     "postgres",
		DBPassword: "qwerty",
		DBName:     "postgres",
		DBSSLMode:  "disable",
	}
}

func NewDefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Env:         "local",
		StoragePath: "./",
		TokenTTL:    "10h",
		Port:        "8080",
	}
}
