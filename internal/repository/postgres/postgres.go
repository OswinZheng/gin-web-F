package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/OswinZheng/gin-web-F/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func New() {
	var err error
	Db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		configs.Get().Database.Host,
		configs.Get().Database.User,
		configs.Get().Database.Password,
		configs.Get().Database.Name,
		configs.Get().Database.Port,
		configs.Get().Database.SslMode,
		configs.Get().Database.TimeZone,
	)), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	if configs.Get().Database.ShowDbLog {
		Db.Logger = logger.Default.LogMode(logger.Info)
	}

	// add update scopes

	if configs.Get().Database.MaxIdleConns > 0 && configs.Get().Database.MaxOpenConns > 0 {
		sqlDb, err := Db.DB()
		if err != nil {
			log.Fatalf("models.Setup err: %v", err)
		}
		sqlDb.SetMaxIdleConns(configs.Get().Database.MaxIdleConns)
		sqlDb.SetMaxOpenConns(configs.Get().Database.MaxOpenConns)
		sqlDb.SetConnMaxLifetime(time.Duration(configs.Get().Database.ConnMaxLifetime) * time.Second)
	}
}
