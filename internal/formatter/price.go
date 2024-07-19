package formatter

import (
	"fmt"
	"strconv"

	"github.com/tinygodsdev/cities/pkg/cities"
	"github.com/tinygodsdev/places-tg-bot/internal/util"
)

func formatPriceAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int
	switch label {
	case cities.AttributePairOfJeans:
		label = cities.AttributePairOfJeansShort
		subgroup = clothing
	case cities.AttributeApples:
		label = cities.AttributeApplesShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeMonthlyNetSalary:
		label = cities.AttributeMonthlyNetSalaryShort
		subgroup = finance
	case cities.AttributeBanana:
		label = cities.AttributeBananaShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeBasicUtilities:
		label = cities.AttributeBasicUtilitiesShort
		subgroup = apartment
	case cities.AttributeBeefRound:
		label = cities.AttributeBeefRoundShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeBottleOfWine:
		label = cities.AttributeBottleOfWineShort
		subgroup = foodAndDrink
	case cities.AttributeCappuccino:
		label = cities.AttributeCappuccinoShort
		subgroup = foodAndDrink
	case cities.AttributeChickenFillets:
		label = cities.AttributeChickenFilletsShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeCigarettes:
		label = cities.AttributeCigarettesShort
		subgroup = various
	case cities.AttributeDomesticBeerBottle:
		label = cities.AttributeDomesticBeerBottleShort
		subgroup = foodAndDrink
	case cities.AttributeEggs:
		label = cities.AttributeEggsShort
		subgroup = foodAndDrink
	case cities.AttributeGasoline:
		label = cities.AttributeGasolineShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerLiterFromGallons(price))
		}
		subgroup = transport
	case cities.AttributeInternationalPrimarySchool:
		label = cities.AttributeInternationalPrimarySchoolShort
		subgroup = education
	case cities.AttributeInternet:
		label = cities.AttributeInternetShort
		subgroup = communication
	case cities.AttributeBread:
		label = cities.AttributeBreadShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeLocalCheese:
		label = cities.AttributeLocalCheeseShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeMcMeal:
		label = cities.AttributeMcMealShort
		subgroup = foodAndDrink
	case cities.AttributeMealFor2:
		label = cities.AttributeMealFor2Short
		subgroup = foodAndDrink
	case cities.AttributeMilk:
		label = cities.AttributeMilkShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerLiterFromGallons(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeMobilePlan:
		label = cities.AttributeMobilePlanShort
		subgroup = communication
	case cities.AttributeMonthlyPass:
		label = cities.AttributeMonthlyPassShort
		subgroup = transport
	case cities.AttributeOneWayTicket:
		label = cities.AttributeOneWayTicketShort
		subgroup = transport
	case cities.AttributePotato:
		label = cities.AttributePotatoShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributePreschool:
		label = cities.AttributePreschoolShort
		subgroup = education
	case cities.AttributeOranges:
		label = cities.AttributeOrangesShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeTaxi1Mile:
		label = cities.AttributeTaxi1MileShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilometerFromMiles(price))
		}
		subgroup = transport
	case cities.AttributeTomato:
		label = cities.AttributeTomatoShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case cities.AttributeWaterBottle:
		label = cities.AttributeWaterBottleShort
		subgroup = foodAndDrink
	case cities.AttributeApartment1BedroomOutsideCentre:
		label = cities.AttributeApartment1BedroomOutsideCentreShort
		subgroup = apartment
	case cities.AttributeApartment1BedroomCityCentre:
		label = cities.AttributeApartment1BedroomCityCentreShort
		subgroup = apartment
	case cities.AttributeMortgageInterestRate:
		label = cities.AttributeMortgageInterestRateShort
		subgroup = finance
	default:
		return formatAttributeResult{}, false
	}

	if label == cities.AttributeMortgageInterestRate || label == cities.AttributeMortgageInterestRateShort {
		values += "%"
	} else {
		values += "$"
	}

	return formatAttributeResult{
		attribute: formatSingleAttribute(label, values, emoji, comment),
		subgroup:  subgroup,
		order:     order,
	}, true
}
