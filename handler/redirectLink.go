package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasedson/go-shortener-link/config"
	"github.com/lucasedson/go-shortener-link/schemas"
)

func RedirectLink(c *gin.Context) {

	// shortLink := c.Param("shortLink")

	var link schemas.Link

	// config.GetDB().Where("short_link is ?", shortLink).First(&link)
	config.GetDB().Where("short_link LIKE ?", "%"+c.Param("shortLink")).First(&link)
	print(link.Url)

	if link.Url == "" {
		print(link.Url)
		c.JSON(404, gin.H{"message": "link not found"})
		return
	}

	// redirect to external link
	c.Redirect(302, link.Url)

}
