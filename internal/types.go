package internal

type Product struct {
	//ID            string  `json:"id"`
	Title         string  `json:"title"`
	Price         float64 `json:"price"`
	PriceOriginal float64 `json:"price_original,omitempty"`
	Site          string  `json:"site"`
	URL           string  `json:"url"`
}

type ConfigScraping struct {
	Site      string
	Selectors SelectorsConfig
}

type SelectorsConfig struct {
	Title string
	Price string
}

type Resultscraping struct {
	Products []Product `json:"products"`
	Erro     string    `json:"error,omitempty"`
}

func (r *Resultscraping) AddProduct(product Product) {
	r.Products = append(r.Products, product)
}
