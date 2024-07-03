package citycountry

var cityCountryMap = map[string]string{
	"Moscow":   "RU",
	"Tel Aviv": "IL",
	"Tbilisi":  "GE",
	"Belgrade": "RS",
	"Istanbul": "TR",
	"London":   "GB",
}

// GetCountryByCity возвращает код страны по названию города.
func GetCountryByCity(city string) (string, bool) {
	country, exists := cityCountryMap[city]
	return country, exists
}
