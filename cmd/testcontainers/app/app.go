package app

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"github.com/sirupsen/logrus"
	"os"
)

func New() IApplication {
	return &Application{}
}

func (application *Application) Build() IApplication {
	application.logger = *logrus.New()
	application.AddRabbitMq(rabbitmq.Opts{
		Username:    "guest",
		Password:    "123456",
		Host:        "localhost",
		VirtualHost: "demand",
	})

	application.AddRouter()
	application.AddControllers().InitMiddlewares().AddSwagger()

	return application
}

func (application *Application) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	defer application.broker.Close()

	err := application.engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	return err
}
