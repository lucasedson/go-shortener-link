package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/lucasedson/go-shortener-link/handler"
)

func initializeRoutes(router *gin.Engine) {
	prefix := "/api/v1"
	router.GET(prefix+"/links", handler.GetLinks)
	router.POST(prefix+"/link", handler.CreateLink)
	router.GET(prefix+"/link/:shortLink", handler.GetLink)

	router.GET("/:shortLink", handler.RedirectLink)
}
