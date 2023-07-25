package db

import (
	"fmt"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func New(settings Settings) (*gorm.DB, error) {
	connection, err := gorm.Open(driver.Open(toDNSForPgSql(settings)), settings.Config())
	if err != nil {
		_, err = fmt.Printf("Can't open connection: %v", err)

		return nil, err
	}

	sqlDB, err := connection.DB()
	if err != nil {
		_, err = fmt.Printf("Can't get sql db: %v", err)

		return nil, err
	}

	sqlDB.SetMaxIdleConns(settings.MaxConnections)
	sqlDB.SetMaxOpenConns(settings.MaxConnections)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return connection, nil
}

func toDNSForPgSql(settings Settings) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Password,
		settings.Database,
		settings.SSLMode,
	)
}
