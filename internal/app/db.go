package app

import (
	"fmt"
	"github.com/Nikkoz/mp.gateway/internal/configs"
	"github.com/Nikkoz/mp.gateway/internal/domain/store"
	"github.com/Nikkoz/mp.gateway/pkg/store/db"
	"gorm.io/gorm"
	"log"
)

func connectionDB() (*gorm.DB, func()) {
	conn, err := db.New(settingsDB(config.Db))
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	return conn, func() {
		sqlDB, _ := conn.DB()
		err := sqlDB.Close()
		if err != nil {
			log.Println(fmt.Errorf("Error close connection: %v\n", err))
		}
	}
}

func settingsDB(config configs.Db) db.Settings {
	sslMode := "disable"
	if config.SslMode {
		sslMode = "enable"
	}

	return db.NewSettings(
		config.Host,
		config.Port,
		config.Name,
		config.User,
		config.Password,
		sslMode,
		"mp_",
		4,
		1000,
	)
}

func migrate(connection *gorm.DB) {
	err := connection.AutoMigrate(&store.Store{})
	if err != nil {
		log.Fatalf("error init models: %v\n", err)
	}

	if err := store.Migrate(connection); err != nil {
		log.Fatalf("error migrate stores table: %v\n", err)
	}
}
