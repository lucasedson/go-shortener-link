package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasedson/go-shortener-link/config"
	"github.com/lucasedson/go-shortener-link/schemas"
)

func GetLinks(c *gin.Context) {

	var links []schemas.Link

	config.GetDB().Find(&links)

	c.JSON(200, gin.H{"message": "success", "data": links})
}
