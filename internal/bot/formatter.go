package bot

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/tinygodsdev/datasdk/pkg/citycountry"
	"github.com/tinygodsdev/datasdk/pkg/data"
	tele "gopkg.in/telebot.v3"
)

const (
	happyEmoji        = "😊"
	loveFaceEmoji     = "😍"
	heartsFaceEmoji   = "🥰"
	satisfiedEmoji    = "😌"
	happyCatEmoji     = "😸"
	partyEmoji        = "🥳"
	neutralEmoji      = "😐"
	sadEmoji          = "😞"
	terrorEmoji       = "😱"
	hotEmoji          = "🥵"
	veryHotEmoji      = hotEmoji + terrorEmoji
	coldEmoji         = "🥶"
	veryColdEmoji     = coldEmoji + terrorEmoji
	wetEmoji          = "💦"
	dryEmoji          = "🌵"
	normalHumEmoji    = "🌿"
	pressureEmoji     = "🌀"
	cloudyEmoji       = "☁️"
	clearEmoji        = "🌞"
	bankEmoji         = "🏦"
	stonksEmoji       = "📈"
	airEmoji          = "💨"
	weatherEmoji      = "🏞️"
	emojiRich         = "🤑"
	emojiPoor         = "💸"
	emojiThought      = "🤔"
	maskEmoji         = "😷"
	treeEmoji         = "🌳"
	skullEmoji        = "💀"
	thunderstormEmoji = "⛈️"

	// weather labels
	attributeTemperature = "temperature"
	attributeHumidity    = "humidity"
	attributePressure    = "pressure"
	attributeDescription = "description"

	// air quality labels
	attributeCo   = "co"
	attributeNo2  = "no2"
	attributeO3   = "o3"
	attributePm10 = "pm10"
	attributePm25 = "pm25"
	attributeSo2  = "so2"

	// world bank labels
	attributeCPI                             = "Consumer price index (2010 = 100)"
	attributeCPIShort                        = "Consumer price index (2010=100)"
	attributeGDPPerCapita                    = "GDP per capita (current US$)"
	attributeGDPPerCapitaShort               = "GDP per capita"
	attributeExports                         = "Merchandise exports (current US$)"
	attributeExportsShort                    = "Exports"
	attributeImports                         = "Merchandise imports (current US$)"
	attributeImportsShort                    = "Imports"
	attributeUnemployment                    = "Unemployment, total (% of total labor force) (modeled ILO estimate)"
	attributeUnemploymentShort               = "Unemployment"
	attributeIndividualsUsingInternet        = "Individuals using the Internet (% of population)"
	attributeIndividualsUsingInternetShort   = "Internet users"
	attributeTaxRevenue                      = "Tax revenue (% of GDP)"
	attributeTaxRevenueShort                 = "Tax revenue"
	attributeLifeExpectancy                  = "Life expectancy at birth, total (years)"
	attributeLifeExpectancyShort             = "Life expectancy"
	attributeMortalityRateUnder5             = "Mortality rate, under-5 (per 1,000 live births)"
	attributeMortalityRateUnder5Short        = "Infant mortality"
	attributeGovtExpenditureEducation        = "Government expenditure on education, total (% of GDP)"
	attributeGovtExpenditureEducationShort   = "Spending on education"
	attributeCO2Emissions                    = "CO2 emissions (metric tons per capita)"
	attributeCO2EmissionsShort               = "CO2 emissions"
	attributeLiteracyRate                    = "Literacy rate, adult total (% of people ages 15 and above)"
	attributeLiteracyRateShort               = "Literacy rate"
	attributeCurrentHealthExpenditure        = "Current health expenditure (% of GDP)"
	attributeCurrentHealthExpenditureShort   = "Health spending"
	attributePovertyHeadcount                = "Poverty headcount ratio at $2.15 a day (2017 PPP) (% of population)"
	attributePovertyHeadcountShort           = "Poverty"
	attributeHealthExpenditurePerCapita      = "Current health expenditure per capita, PPP (current international $)"
	attributeHealthExpenditurePerCapitaShort = "Health spending per capita"

	// categories
	catergoryWeather    = "weather"
	catergoryAirQuality = "air_quality"
	categoryWorldBank   = "world_bank"

	// provider
	provider = "tinygods.dev"
)

var categoryOrder = []string{catergoryWeather, catergoryAirQuality, categoryWorldBank}

func FormatCitiesReport(points []data.Point) []string {
	groupedData := groupDataByCityAndCategory(points)
	var messages []string
	for city, categories := range groupedData {
		messages = append(messages, bold(upper(city)))
		for _, category := range categoryOrder {
			if attrs, exists := categories[category]; exists {
				messages = append(messages, formatCatergoryTitle(category, city))
				messages = append(messages, formatCityAttributes(attrs))
			}
		}
	}
	return messages
}

func FormatSources(sources []data.Source) string {
	if len(sources) == 0 {
		return ""
	}

	uniqueNames := make(map[string]struct{})
	for _, source := range sources {
		uniqueNames[source.Name] = struct{}{}
	}

	var names []string
	for name := range uniqueNames {
		names = append(names, name)
	}
	sort.Strings(names)

	return fmt.Sprintf("%s: %s", bold("Sources"), strings.Join(names, ", "))
}

func FormatFetchDuration(d time.Duration) string {
	return italic(fmt.Sprintf("Fetched in %s", d))
}

func FormatProvider() string {
	return "By " + underline(provider)
}

func FormatDeveloperPlain() string {
	return "Developed by " + provider
}

func FormatMessageFooter(sources []data.Source, startTime time.Time) string {
	var res []string
	if len(sources) != 0 {
		res = append(res, FormatSources(sources))
	}

	if !startTime.IsZero() {
		res = append(res, FormatFetchDuration(time.Since(startTime)))
	}

	res = append(res, FormatProvider())
	return strings.Join(res, "\n")
}

func formatCityAttributes(attributes []data.Attribute) string {
	var result []string
	var latestTime time.Time
	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Label < attributes[j].Label
	})
	for _, attr := range attributes {
		if len(attr.Values) == 0 || len(attr.Timestamps) == 0 {
			continue
		}

		lastValue := attr.Values[len(attr.Values)-1]
		lastTimestamp := attr.Timestamps[len(attr.Timestamps)-1]
		if lastTimestamp.After(latestTime) {
			latestTime = lastTimestamp
		}
		attributeString := formatAttribute(attr.Label, lastValue, attr.Comment)
		result = append(result, attributeString)
	}
	if !latestTime.IsZero() {
		result = append(result, italic(fmt.Sprintf("updated at %s", latestTime.Format("2006-01-02 15:04"))))
	}
	return strings.Join(result, "\n")
}

func formatWorldBankAttribute(label string, values string, comment string) (string, bool) {
	var emoji string
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
			values = formatLargeNumber(exports) + "$"
		}
	case attributeImports:
		label = attributeImportsShort
		imports, err := strconv.ParseFloat(values, 64)
		if err == nil {
			values = formatLargeNumber(imports) + "$"
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
		return "", false
	}

	if comment != "" {
		values = fmt.Sprintf("%s (%s)", values, comment)
	}

	return fmt.Sprintf("%s: %s %s", bold(capitalize(label)), values, emoji), true
}

func formatWeatherAttribute(label string, values string, comment string) (string, bool) {
	var emoji string
	switch label {
	case attributeTemperature:
		temp, err := strconv.ParseFloat(values, 64)
		if err == nil {
			switch {
			case temp > 34:
				emoji = veryHotEmoji
			case temp > 26:
				emoji = hotEmoji
			case temp > 15:
				emoji = randomHappyEmoji()
			case temp > 5:
				emoji = neutralEmoji
			case temp >= 0:
				emoji = coldEmoji
			default:
				emoji = veryColdEmoji
			}
			values += "°"
		}
	case attributeHumidity:
		hum, err := strconv.ParseFloat(values, 64)
		if err == nil {
			if hum > 80 {
				emoji = wetEmoji
			} else if hum < 20 {
				emoji = dryEmoji
			} else {
				emoji = normalHumEmoji
			}
			values += "%"
		}
	case attributePressure:
		press, err := strconv.ParseFloat(values, 64)
		if err == nil {
			switch {
			case press > 1020:
				emoji = pressureEmoji
			case press > 1000:
				emoji = randomHappyEmoji()
			default:
				emoji = sadEmoji
			}
			values += " hPa"
		}
	case attributeDescription:
		if strings.Contains(values, "cloud") {
			emoji = cloudyEmoji
		} else if strings.Contains(values, "clear") {
			emoji = clearEmoji
		} else if strings.Contains(values, "thunderstorm") {
			emoji = thunderstormEmoji
		}
	default:
		return "", false
	}

	if comment != "" {
		values = fmt.Sprintf("%s (%s)", values, comment)
	}

	return fmt.Sprintf("%s: %s %s", bold(capitalize(label)), values, emoji), true
}

func formatAirQualityAttribute(label string, values string, comment string) (string, bool) {
	var emoji string
	value, err := strconv.ParseFloat(values, 64)
	if err != nil {
		return "", false
	}

	switch label {
	case attributeCo:
		switch {
		case value > 200:
			emoji = skullEmoji
		case value > 100:
			emoji = maskEmoji
		case value <= 50:
			emoji = treeEmoji
		default:
			emoji = ""
		}
	case attributeNo2:
		switch {
		case value >= 101:
			emoji = skullEmoji
		case value >= 40:
			emoji = maskEmoji
		case value <= 30:
			emoji = treeEmoji
		default:
			emoji = ""
		}
	case attributeO3:
		switch {
		case value >= 100:
			emoji = skullEmoji
		case value >= 50:
			emoji = maskEmoji
		case value <= 40:
			emoji = treeEmoji
		default:
			emoji = ""
		}
	case attributePm10:
		switch {
		case value >= 51:
			emoji = skullEmoji
		case value >= 21:
			emoji = maskEmoji
		case value <= 20:
			emoji = treeEmoji
		default:
			emoji = ""
		}
	case attributePm25:
		switch {
		case value >= 26:
			emoji = skullEmoji
		case value >= 11:
			emoji = maskEmoji
		case value <= 10:
			emoji = treeEmoji
		default:
			emoji = ""
		}
	case attributeSo2:
		switch {
		case value >= 76:
			emoji = skullEmoji
		case value >= 21:
			emoji = maskEmoji
		case value <= 20:
			emoji = treeEmoji
		default:
			emoji = ""
		}
	default:
		return "", false
	}

	if comment != "" {
		values = fmt.Sprintf("%s (%s)", values, comment)
	}

	return fmt.Sprintf("%s: %s %s", bold(capitalize(label)), values, emoji), true
}

func formatAttribute(label string, values string, comment string) string {
	if result, formatted := formatWeatherAttribute(label, values, comment); formatted {
		return result
	}

	if result, formatted := formatWorldBankAttribute(label, values, comment); formatted {
		return result
	}

	if result, formatted := formatAirQualityAttribute(label, values, comment); formatted {
		return result
	}

	// not a special case, just return the label and values
	return fmt.Sprintf("%s: %s", bold(capitalize(label)), values)
}

func formatCatergoryTitle(category, city string) string {
	switch category {
	case catergoryWeather:
		return underline("Weather") + " " + weatherEmoji
	case catergoryAirQuality:
		return underline("Air Quality") + " " + airEmoji
	case categoryWorldBank:
		title := "National Economy"
		flag, ok := citycountry.GetFlagByCity(city)
		if ok {
			return underline(title) + " " + flag
		}
		return underline(title) + " " + stonksEmoji
	default:
		return capitalize(category)
	}
}

func FormatCityList(cities []string) string {
	var formattedCities []string
	for _, c := range cities {
		formattedCities = append(formattedCities,
			capitalize(c),
		)
	}

	return strings.Join(formattedCities, ", ")
}

func FormatCommands(commands []tele.Command) string {
	var formattedCommands []string
	for _, c := range commands {
		formattedCommands = append(formattedCommands,
			fmt.Sprintf("%s - %s", c.Text, c.Description),
		)
	}

	return strings.Join(formattedCommands, "\n")
}

func groupDataByCityAndCategory(points []data.Point) map[string]map[string][]data.Attribute {
	result := make(map[string]map[string][]data.Attribute)
	for _, point := range points {
		city := getCityFromTags(point.Tags)
		category := getCategoryFromTags(point.Tags)
		if city == "" || category == "" {
			continue
		}
		if _, exists := result[city]; !exists {
			result[city] = make(map[string][]data.Attribute)
		}
		result[city][category] = append(result[city][category], point.Attributes...)
	}
	return result
}

func getCityFromTags(tags []data.Tag) string {
	for _, tag := range tags {
		if tag.Label == "city" {
			return tag.Value
		}
	}
	return ""
}

func getCategoryFromTags(tags []data.Tag) string {
	for _, tag := range tags {
		if tag.Label == "category" {
			return tag.Value
		}
	}
	return ""
}

func bold(s string) string {
	return fmt.Sprintf("<b>%s</b>", s)
}

func underline(s string) string {
	return fmt.Sprintf("<u>%s</u>", s)
}

func italic(s string) string {
	return fmt.Sprintf("<i>%s</i>", s)
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func randomHappyEmoji() string {
	happyEmojis := []string{
		happyEmoji,
		heartsFaceEmoji,
		satisfiedEmoji,
		loveFaceEmoji,
		happyCatEmoji,
		partyEmoji,
	}

	return happyEmojis[rand.Intn(len(happyEmojis))]
}

func formatLargeNumber(num float64) string {
	absNum := math.Abs(num)
	var divisor float64
	var suffix string

	switch {
	case absNum >= 1e12:
		divisor = 1e12
		suffix = "T"
	case absNum >= 1e9:
		divisor = 1e9
		suffix = "B"
	case absNum >= 1e6:
		divisor = 1e6
		suffix = "M"
	case absNum >= 1e3:
		divisor = 1e3
		suffix = "K"
	default:
		divisor = 1
		suffix = ""
	}

	formattedNumber := num / divisor
	if divisor > 1 {
		return fmt.Sprintf("%.0f%s", formattedNumber, suffix)
	}
	return fmt.Sprintf("%.0f", formattedNumber)
}
