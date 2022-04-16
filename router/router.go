package router

import (
	"BenchMarkDataBase/apis"
	"BenchMarkDataBase/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	router.Use(middleware.LoggerToFile())

	router.Use(gin.Recovery())
	v1 := router.Group("/apis/v1")
	{
		v1.GET("/ping", apis.Ping)
		hostsGroup := v1.Group("/hosts")
		hostsGroup.POST("/addOne", apis.HostAddOne)
		hostsGroup.POST("/deleteOne", apis.HostDeleteOne)
		hostsGroup.POST("/ListOne", apis.HostListOne)
		hostsGroup.POST("/List", apis.HostList)
	}
	return router
}
