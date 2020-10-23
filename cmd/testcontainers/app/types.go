package app

import (
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IApplication interface {
	Build() IApplication
	Run() error
}

type Application struct {
	logger logrus.Logger
	engine *gin.Engine
	broker rabbitmq.IRabbitMq
	db     *gorm.DB
}
