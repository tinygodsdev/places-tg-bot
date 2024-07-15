package formatter

import (
	"strconv"

	"github.com/tinygodsdev/cities/cities"
)

func formatWorldBankAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int
	switch label {
	case cities.AttributeCPI:
		label = cities.AttributeCPIShort
		cpi, err := strconv.ParseFloat(values, 64)
		if err == nil {
			switch {
			case cpi > 500:
				emoji = terrorEmoji
			case cpi > 200:
				emoji = emojiThought
			default:
				emoji = ""
			}
		}
	case cities.AttributeGDPPerCapita:
		label = cities.AttributeGDPPerCapitaShort
		gdp, err := strconv.ParseFloat(values, 64)
		if err == nil {
			switch {
			case gdp > 40000:
				emoji = emojiRich
			case gdp > 10000:
				emoji = ""
			default:
				emoji = emojiPoor
			}
		}
		values += "$"
	case cities.AttributeExports:
		label = cities.AttributeExportsShort
		exports, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = FormatLargeNumber(exports) + "$"
		}
	case cities.AttributeImports:
		label = cities.AttributeImportsShort
		imports, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = FormatLargeNumber(imports) + "$"
		}
	case cities.AttributeUnemployment:
		label = cities.AttributeUnemploymentShort
		values += "%"
	case cities.AttributeIndividualsUsingInternet:
		label = cities.AttributeIndividualsUsingInternetShort
		values += "%"
	case cities.AttributeTaxRevenue:
		label = cities.AttributeTaxRevenueShort
		values += "% of GDP"
	case cities.AttributeLifeExpectancy:
		label = cities.AttributeLifeExpectancyShort
		values += " years"
	case cities.AttributeMortalityRateUnder5:
		label = cities.AttributeMortalityRateUnder5Short
		values += " per 1000"
	case cities.AttributeGovtExpenditureEducation:
		label = cities.AttributeGovtExpenditureEducationShort
		values += "% of GDP"
	case cities.AttributeCO2Emissions:
		label = cities.AttributeCO2EmissionsShort
		values += " tons per capita"
	case cities.AttributeLiteracyRate:
		label = cities.AttributeLiteracyRateShort
		values += "%"
	case cities.AttributeCurrentHealthExpenditure:
		label = cities.AttributeCurrentHealthExpenditureShort
		values += "% of GDP"
	case cities.AttributeHealthExpenditurePerCapita:
		label = cities.AttributeHealthExpenditurePerCapitaShort
		values += "$"
	case cities.AttributePovertyHeadcount:
		label = cities.AttributePovertyHeadcountShort
		values += "%"
	default:
		return formatAttributeResult{}, false
	}

	return formatAttributeResult{
		attribute: formatSingleAttribute(label, values, emoji, comment),
		subgroup:  subgroup,
		order:     order,
	}, true
}
