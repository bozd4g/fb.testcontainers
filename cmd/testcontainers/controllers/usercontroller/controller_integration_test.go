package usercontroller

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntegrationSuite struct {
	suite.Suite
	controller IUserController
}

func TestInit(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (suite *IntegrationSuite) SetupSuite() {}

func (suite *IntegrationSuite) AfterTest(_, _ string) {}

func (suite *IntegrationSuite) Test_CreateUser_WhenIsCreated_ReturnsSuccess() {}
