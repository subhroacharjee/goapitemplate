package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/subhroacharjee/appname/pkg/controllers/health"
)

func InitRouter(g *gin.Engine) {
	router := g.Group("/api")
	// define routes and controller.
	health.InitHealthController(router)
}
