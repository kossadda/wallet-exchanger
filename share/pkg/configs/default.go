package configs

import "time"

const (
	DefaultGrpcPort          = "44044"
	DefaultCachePort         = "6379"
	DefaultWalletServicePort = "40404"
	DefaultPostgresPort      = "5436"
	DefaultTokenExpire       = time.Hour * 24
	DefaultCacheExpire       = time.Minute
	DefaultENV               = "local"
)
