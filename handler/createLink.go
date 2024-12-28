package handler

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/lucasedson/go-shortener-link/config"
	"github.com/lucasedson/go-shortener-link/schemas"
)

func CreateLink(c *gin.Context) {
	var requestBody schemas.Link

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	hash := sha256.Sum256([]byte(requestBody.Url))
	base64Hash := base64.URLEncoding.EncodeToString(hash[:8])
	data := schemas.Link{Url: requestBody.Url, ShortLink: "http://localhost:8080/" + base64Hash}
	config.GetDB().Create(&data)

	c.JSON(200, gin.H{"message": "success", "data": data.ShortLink})
}
