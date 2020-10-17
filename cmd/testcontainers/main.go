package main

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/app"
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
// @BasePath /
func main() {
	application := app.New()
	err := application.Build().Run()
	if err != nil {
		panic(fmt.Sprintf("app cannot be started! Error %+v", err))
	}
}
