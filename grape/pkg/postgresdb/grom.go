package postgresdb

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbconn *gorm.DB
)

func Connect(dsn string) error {
	var err error
	l := logger.New(
		log.New(os.Stdout, "", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	dbconn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: l,
	})
	return err
}

func GetDB() *gorm.DB {
	return dbconn
}
