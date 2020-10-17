package app

import "github.com/gin-gonic/gin"

func (application *Application) AddRouter() {
	application.engine = gin.Default()
}
