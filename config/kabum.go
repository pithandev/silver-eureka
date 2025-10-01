package config

type ConfigScraping struct {
	Site      string
	Selectors struct {
		Title string
		Price string
	}
}

func GetConfigKabum() ConfigScraping {
	return ConfigScraping{
		Site: "Kabum",
		Selectors: struct {
			Title string
			Price string
		}{
			Title: ".text-sm.desktop\\:text-xl.text-black-800.font-bold.desktop\\:font-bold",
			Price: ".text-4xl.text-secondary-500.font-bold.transition-all.duration-500",
		},
	}
}
