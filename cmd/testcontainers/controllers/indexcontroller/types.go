package indexcontroller

import (
	"github.com/gin-gonic/gin"
)

type IIndexController interface {
	Init(e *gin.Engine)
}

type IndexController struct {
}
