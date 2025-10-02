package internal

import (
	"github.com/pithandev/silver-eureka/config"

	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ScraperKabum struct {
	client *http.Client
	config config.ConfigScraping
}

func NewScraperKabum() *ScraperKabum {
	return &ScraperKabum{
		client: &http.Client{Timeout: 30 * time.Second},
		config: config.GetConfigKabum(),
	}
}

func extractData(doc *goquery.Document, url string) (*Product, error)
