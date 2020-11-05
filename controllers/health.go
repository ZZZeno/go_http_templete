package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/middlewares"
	"net/http"
)

type HealthController struct{}

var logger = middlewares.Logger("xxx")

func (h HealthController) Status(c *gin.Context) {
	user := c.GetString("CAS_USERNAME")
	logger.Info("记录一下日志")
	c.String(http.StatusOK, fmt.Sprintf("Working for %+v !!!", user))
}

