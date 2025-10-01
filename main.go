package main

import (
	"fmt"

	"github.com/pithandev/silver-eureka/config"
)

func main() {
	configKabum := config.GetConfigKabum()

	fmt.Println("🔧 Kabum configuration:")
	fmt.Printf("Site: %s ", configKabum.Site)
	fmt.Printf("Title: %s ", configKabum.Selectors.Title)
	fmt.Printf("Price: %s ", configKabum.Selectors.Price)

}
