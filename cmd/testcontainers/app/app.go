package app

import (
	"fmt"
	"os"
)

func New() IApplication {
	return &Application{}
}

func (application *Application) Build() IApplication {
	application.AddRouter()
	application.AddControllers().InitMiddlewares().AddSwagger()

	application.AddRabbitMq()
	
	return application
}

func (application *Application) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := application.engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	return err
}
