package main

import (
	"gin_api/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.IndexApi)

	router.POST("/GetSysDict", apis.GetSysDict)

	return router
}
