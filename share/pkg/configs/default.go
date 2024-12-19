package configs

import "time"

const (
	DefaultGrpcPort          = "44044"
	DefaultWalletServicePort = "40404"
	DefaultPostgresPort      = "5432"
	DefaultTokenExpire       = time.Hour * 24
	DefaultCacheExpire       = time.Minute
	DefaultENV               = "local"
)
