package usercontroller

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/application/userservice"
	"github.com/gin-gonic/gin"
)

type IUserController interface{
	Init()
}

type UserController struct {
	engine  *gin.Engine
	service userservice.IUserService
}