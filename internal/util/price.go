package util

const (
	// conversion coefficients
	poundsToKilograms = 0.453592
	gallonsToLiters   = 3.78541
	ouncesToLiters    = 0.0295735
	milesToKilometers = 1.60934
)

func PricePerKilogramFromPounds(pricePerPound float64) float64 {
	return pricePerPound / poundsToKilograms
}

func PricePerLiterFromGallons(pricePerGallon float64) float64 {
	return pricePerGallon / gallonsToLiters
}

func PricePerLiterFromOunces(pricePerOunce float64) float64 {
	return pricePerOunce / ouncesToLiters
}

func PricePerKilometerFromMiles(pricePerMile float64) float64 {
	return pricePerMile / milesToKilometers
}
