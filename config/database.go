package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {
	// init connection mysql
	dsn := "root:password@tcp(127.0.0.1:3306)/test_user?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})

	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("ERROR Setup pooling: %s", err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	log.Printf("INFO: Connected to DB")

	return db, nil

}
