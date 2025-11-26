package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sabari7320/cli-url-shortner/internal/db"
	"github.com/sabari7320/cli-url-shortner/internal/models"
)

func Start() {

	r := gin.Default()
	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		var url models.URL
		err := db.DB.Where("short_code =?", code).First(&url).Error
		if err != nil {
			c.String(http.StatusNotFound, "Short URL not found")
			return
		}

		//UPDATE CLICK COUNT && TIME STAMPS

		now := time.Now()
		db.DB.Model(&url).Updates(models.URL{
			Clicks: url.Clicks + 1,
			FirstClick: func() *time.Time {
				if url.FirstClick == nil {
					return &now
				}
				return url.FirstClick
			}(),
			LastClick: &now,
		})

		c.Redirect(http.StatusMovedPermanently, url.LongUrl)

	})

	r.Run(":8080")

}
