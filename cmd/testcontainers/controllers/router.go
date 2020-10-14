package controllers

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/controllers/usercontroller"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/application/userservice"
	_ "github.com/bozd4g/fb.testcontainers/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type IRouter interface {
	Get() *gin.Engine
	InitRoutes()
	InitMiddlewares()
}

type Router struct {
	engine *gin.Engine

	userService userservice.IUserService
}

func New(userService userservice.IUserService) IRouter {
	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ""))
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	})

	return &Router{engine: engine, userService: userService}
}

func (router *Router) Get() *gin.Engine {
	return router.engine
}

func (router *Router) InitRoutes() {
	usercontroller.New(router.engine, router.userService).Init()
}

func (router *Router) InitMiddlewares() {
	router.engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}
