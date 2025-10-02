package main

import (
	"fmt"
	"log"

	"github.com/pithandev/silver-eureka/internal"
	scraper "github.com/pithandev/silver-eureka/internal/scraper"
)

// TestarConversao testa vários formatos de preço
func TestarConversao() {
	urls := []string{
		"https://www.kabum.com.br/produto/922662",
		"https://www.kabum.com.br/produto/590662",
		"https://www.kabum.com.br/produto/695107",
	}
	scraperKabum := scraper.NewScrapeKabum()
	var result []internal.Product

	for i, url := range urls {
		fmt.Printf("\n🔍 Processando %d/%d: %s\n", i+1, len(urls), url)

		produto, err := scraperKabum.ScrapeProduct(url)
		if err != nil {
			log.Printf("❌ Erro em %s: %v", url, err)
			continue
		}

		result = append(result, *produto)
		fmt.Printf("✅ Extraído: %s - R$ %.2f\n", produto.Title, produto.Price)
	}
}

func main() {
	TestarConversao()
}
