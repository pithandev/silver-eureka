package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pithandev/silver-eureka/config"
	"github.com/pithandev/silver-eureka/internal"
	"github.com/pithandev/silver-eureka/utils"

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

func (s *ScraperKabum) ScraperProduct(url string) (*internal.Product, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error trying to create request: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("status code erro: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear HTML: %v", err)
	}

	return s.extractData(doc, url)
}

func (s *ScraperKabum) extractData(doc *goquery.Document, url string) (*internal.Product, error) {
	product := &internal.Product{
		URL:  url,
		Site: "Kabum",
	}

	title := doc.Find(s.config.Selectors.Title).First().Text()

	if title == "" {
		return nil, fmt.Errorf("title not found")
	}
	product.Title = strings.TrimSpace(title)

	priceText := doc.Find(s.config.Selectors.Price).First().Text()

	if priceText == "" {
		return product, fmt.Errorf("title not found")
	}
	priceText = strings.TrimSpace(priceText)
	price, err := utils.ConvertPrice(priceText)

	if err != nil {
		return nil, fmt.Errorf("erro trying to converto price %s: %v", priceText, err)
	}

	product.Price = price

	return product, nil
}
