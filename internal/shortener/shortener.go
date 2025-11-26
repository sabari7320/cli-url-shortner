package shortener

import (
	"github.com/jcoene/go-base62"
	"github.com/sabari7320/cli-url-shortner/internal/db"
	"github.com/sabari7320/cli-url-shortner/internal/models"
)

var baseDomain = "https://s.ly"

func Create(longURL string) (string, error) {

	var count int64
	db.DB.Model(&models.URL{}).Count(&count)
	shortCode := base62.EncodeUint64(uint64(count + 9000000000))

	url := models.URL{
		LongUrl:   longURL,
		ShortCode: shortCode,
		Clicks:    0,
	}

	db.DB.Create(&url)
	return baseDomain + "/" + shortCode, nil

}
