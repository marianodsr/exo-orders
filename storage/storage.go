package storage

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

//InitDB func
func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	dsn := "host=localhost user=postgres password=auth159123456 dbname=exo-orders port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

}

//GetDbConnection func
func GetDbConnection() *gorm.DB {
	return db
}
