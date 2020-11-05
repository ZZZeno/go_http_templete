package server

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/controllers"
	"go-gin-boilerplate/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.LoggerToFile())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.Use(middlewares.FillSession())

	router.GET("/health", middlewares.AuthMiddleware(), health.Status)


	return router

}
