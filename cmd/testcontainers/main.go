package main

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/controllers"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/application/userservice"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/repository/userrepository"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"os"
)

// @title User API
// @version 1.0
// @description This is a user microservice.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email me@furkanbozdag.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	broker, err := rabbitmq.New("amqp://guest:123456@localhost:5672/demand")
	if err != nil {
		panic(err)
	}

	userRepository, err := userrepository.New()
	if err != nil {
		panic(err)
	}
	userService := userservice.New(broker, userRepository)

	router := controllers.New(userService)
	router.InitRoutes()
	router.InitMiddlewares()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = router.Get().Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
