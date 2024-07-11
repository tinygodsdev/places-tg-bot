package formatter

import "strconv"

func formatWorldBankAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int
	switch label {
	case attributeCPI:
		label = attributeCPIShort
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
	case attributeGDPPerCapita:
		label = attributeGDPPerCapitaShort
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
	case attributeExports:
		label = attributeExportsShort
		exports, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = FormatLargeNumber(exports) + "$"
		}
	case attributeImports:
		label = attributeImportsShort
		imports, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = FormatLargeNumber(imports) + "$"
		}
	case attributeUnemployment:
		label = attributeUnemploymentShort
		values += "%"
	case attributeIndividualsUsingInternet:
		label = attributeIndividualsUsingInternetShort
		values += "%"
	case attributeTaxRevenue:
		label = attributeTaxRevenueShort
		values += "% of GDP"
	case attributeLifeExpectancy:
		label = attributeLifeExpectancyShort
		values += " years"
	case attributeMortalityRateUnder5:
		label = attributeMortalityRateUnder5Short
		values += " per 1000"
	case attributeGovtExpenditureEducation:
		label = attributeGovtExpenditureEducationShort
		values += "% of GDP"
	case attributeCO2Emissions:
		label = attributeCO2EmissionsShort
		values += " tons per capita"
	case attributeLiteracyRate:
		label = attributeLiteracyRateShort
		values += "%"
	case attributeCurrentHealthExpenditure:
		label = attributeCurrentHealthExpenditureShort
		values += "% of GDP"
	case attributeHealthExpenditurePerCapita:
		label = attributeHealthExpenditurePerCapitaShort
		values += "$"
	case attributePovertyHeadcount:
		label = attributePovertyHeadcountShort
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
