package formatter

import (
	"fmt"
	"strconv"

	"github.com/tinygodsdev/places-tg-bot/internal/util"
)

func formatPriceAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int
	switch label {
	case attributePairOfJeans:
		label = attributePairOfJeansShort
	case attributeApples:
		label = attributeApplesShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeMonthlyNetSalary:
		label = attributeMonthlyNetSalaryShort
	case attributeBanana:
		label = attributeBananaShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeBasicUtilities:
		label = attributeBasicUtilitiesShort
		subgroup = apartment
	case attributeBeefRound:
		label = attribureBeefRoundShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeBottleOfWine:
		label = attributeBottleOfWineShort
		subgroup = foodAndDrink
	case attributeCappuccino:
		label = attributeCappuccinoShort
		subgroup = foodAndDrink
	case attributeChickenFillets:
		label = attributeChickenFilletsShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeCigarettes:
		label = attributeCigarettesShort
	case attributeDomesticBeerBottle:
		label = attributeDomesticBeerBottleShort
		subgroup = foodAndDrink
	case attributeEggs:
		label = attributeEggsShort
		subgroup = foodAndDrink
	case attributeGasoline:
		label = attributeGasolineShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerLiterFromGallons(price))
		}
		subgroup = transport
	case attributeInternationalPrimarySchool:
		label = attributeInternationalPrimarySchoolShort
		subgroup = education
	case attributeInternet:
		label = attributeInternetShort
	case attributeBread:
		label = attributeBreadShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeLocalCheese:
		label = attributeLocalCheeseShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeMcMeal:
		label = attributeMcMealShort
		subgroup = foodAndDrink
	case attributeMealFor2:
		label = attributeMealFor2Short
		subgroup = foodAndDrink
	case attributeMilk:
		label = attributeMilkShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerLiterFromGallons(price))
		}
		subgroup = foodAndDrink
	case attributeMobilePlan:
		label = attributeMobilePlanShort
	case attributeMonthlyPass:
		label = attributeMonthlyPassShort
		subgroup = transport
	case attributeOneWayTicket:
		label = attributeOneWayTicketShort
		subgroup = transport
	case attributePotato:
		label = attributePotatoShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributePreschool:
		label = attributePreschoolShort
		subgroup = education
	case attributeOranges:
		label = attributeOrangesShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeTaxi1Mile:
		label = attributeTaxi1MileShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilometerFromMiles(price))
		}
		subgroup = transport
	case attributeTomato:
		label = attributeTomatoShort
		price, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = fmt.Sprintf("%.2f", util.PricePerKilogramFromPounds(price))
		}
		subgroup = foodAndDrink
	case attributeWaterBottle:
		label = attributeWaterBottleShort
		subgroup = foodAndDrink
	case attributeApartment1BedroomOutsideCentre:
		label = attributeApartment1BedroomOutsideCentreShort
		subgroup = apartment
	case attributeApartment1BedroomCityCentre:
		label = attributeApartment1BedroomCityCentreShort
		subgroup = apartment
	case attributeMortgageInterestRate:
		label = attributeMortgageInterestRateShort
		subgroup = apartment
	default:
		return formatAttributeResult{}, false
	}

	if label == attributeMortgageInterestRate || label == attributeMortgageInterestRateShort {
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
