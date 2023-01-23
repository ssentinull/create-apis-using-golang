package config

import "time"

const (
	DefaultPostgresMaxIdleConns    = 3
	DefaultPostgresMaxOpenConns    = 5
	DefaultPostgresConnMaxLifetime = 1 * time.Hour
	DefaultPostgresPingInterval    = 1 * time.Second
	DefaultPostgresRetryAttempts   = 3
)
