package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/engine"
)

type userController struct{}

func (receiver *userController) Login(c *gin.Context) engine.Result {
	return engine.SuccessRes(nil)
}
