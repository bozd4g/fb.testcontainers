package usercontroller

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/application/userservice"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New(service userservice.IUserService) IUserController {
	return &UserController{service: service}
}

func (controller UserController) Init(e *gin.Engine) {
	group := e.Group("api/users")
	{
		group.POST("", controller.createHandler)
		group.GET("", controller.getAllHandler)
	}
}

// @Summary Create a user
// @Description This method creates a new user
// @Accept  json
// @Produce  json
// @tags UserController
// @param UserCreateRequestDto body userservice.UserCreateRequestDto true "Create a user"
// @Success 201 {string} string	"Success"
// @Router /api/users [post]
func (controller UserController) createHandler(c *gin.Context) {
	var userDto userservice.UserCreateRequestDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.service.Create(userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occured while creating the user! Please try again later."})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary Get all users
// @Description This method returns all users recorded in the database
// @Accept  json
// @Produce  json
// @tags UserController
// @Success 200 {object} []userservice.UserDto "Success"
// @Router /api/users [get]
func (controller UserController) getAllHandler(c *gin.Context) {
	users, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured while retrieving the users!"})
		return
	}

	c.JSON(http.StatusOK, users)
}
