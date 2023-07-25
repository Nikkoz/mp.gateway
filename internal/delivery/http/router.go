package http

import (
	"github.com/Nikkoz/mp.gateway/internal/configs"
	"github.com/Nikkoz/mp.gateway/internal/configs/types/logger"
	"github.com/Nikkoz/mp.gateway/internal/delivery/http/middlewares"
	"github.com/gin-gonic/gin"
)

func (d *Delivery) initRouter(config configs.Config) {
	if config.App.Environment.IsProduction() {
		switch config.Log.Level {
		case logger.Debug:
			gin.SetMode(gin.DebugMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.Use(middlewares.Auth(config.Auth))

	d.stores(router.Group("/stores"))

	d.router = router
}

func (d *Delivery) stores(router *gin.RouterGroup) {
	router.POST("/", d.handlers.Store.Create)
	//router.GET("/", d.handlers.Store.List)

	d.store(router.Group("/:id"))
}

func (d *Delivery) store(router *gin.RouterGroup) {
	//router.PUT("/", d.handlers.Store.Update)
	//router.DELETE("/", d.handlers.Store.Delete)
}
