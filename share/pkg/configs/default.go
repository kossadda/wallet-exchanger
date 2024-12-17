package configs

import "time"

const (
	DefaultGrpcPort          = "44044"
	DefaultWalletServicePort = "40404"
	DefaultPostgresPort      = "5436"
	DefaultTokenTTL          = time.Hour * 24
	DefaultENV               = "local"
)
