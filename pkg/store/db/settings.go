package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Settings struct {
	Host           string
	Port           uint16
	Database       string
	User           string
	Password       string
	SSLMode        string
	MaxConnections int
	LifeTime       time.Duration
	Prefix         string
	BatchSize      int
}

func NewSettings(host string, port uint16, db, user, pass, sslMode, prefix string, maxConn, batchSize int) Settings {
	return Settings{
		Host:           host,
		Port:           port,
		Database:       db,
		User:           user,
		Password:       pass,
		SSLMode:        sslMode,
		Prefix:         prefix,
		MaxConnections: maxConn,
		LifeTime:       time.Hour,
		BatchSize:      batchSize,
	}
}

func (s Settings) Config() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   s.Prefix,
			SingularTable: false,
		},
		CreateBatchSize: s.BatchSize,
		//DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction: true,
		//DryRun: true,
	}
}
