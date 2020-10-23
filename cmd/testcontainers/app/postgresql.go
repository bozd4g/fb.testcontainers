package app

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (application *Application) AddPostgreSql(dsnString string) *Application {
	db, err := gorm.Open(postgres.Open(dsnString), &gorm.Config{})
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
