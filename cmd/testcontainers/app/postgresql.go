package app

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSqlOpts struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
}

func (application *Application) AddPostgreSql(opts PostgreSqlOpts) *Application {
	defaultDsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	dsn := fmt.Sprintf(defaultDsn, opts.Host, opts.User, opts.Password, opts.Database, opts.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		application.logger.Error("An error occured while connection to postgresql! ", err)
		return application
	}

	application.db = db

	application.migrate()
	return application
}

func (application *Application) migrate() {
	err := application.db.AutoMigrate(user.Entity{})
	if err != nil {
		application.logger.Error("An error occured while migrating to postgresql! ", err)
	}
}
