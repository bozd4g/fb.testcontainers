package indexcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() IIndexController {
	return &IndexController{}
}

func (controller IndexController) Init(e *gin.Engine) {
	e.GET("/", controller.indexHandler)
}

// @Summary redirectToSwaggerUi
// @Description This method redirects to swagger ui
// @Accept  json
// @Produce  json
// @tags IndexController
// @Success 308 {string} string	"Redirect"
// @Router / [get]
func (controller IndexController) indexHandler(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}