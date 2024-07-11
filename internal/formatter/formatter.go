package formatter

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/citycountry"
	"github.com/tinygodsdev/datasdk/pkg/data"
	tele "gopkg.in/telebot.v3"
)

var categoryOrder = []string{catergoryWeather, catergoryAirQuality, categoryPrices, categoryWorldBank}

type formatAttributeResult struct {
	attribute string
	subgroup  string
	order     int
}

func getSkipLabels() map[string]struct{} {
	return map[string]struct{}{
		// world bank
		attributeCO2Emissions: {},
		// prices
		attributeMenLeatherBusinessShoes:         {},
		attributeNikeRunningShoes:                {},
		attributeSummerDressChainStore:           {},
		attributeCinemaSeat:                      {},
		attributeTaxi1HourWaiting:                {},
		attributeToyotaCorolla:                   {},
		attributeVolkswagenGolf:                  {},
		attributeWaterSmallBottle:                {},
		attributeLettuce:                         {},
		attributeTennisCourtRent:                 {},
		attributeTaxiStart:                       {},
		attributeDomesticBeerPint:                {},
		attributeImportedBeer:                    {},
		attributeCokePepsi:                       {},
		attributeRice:                            {},
		attributeOnion:                           {},
		attributeFitnessClubFee:                  {},
		attributeApartment3BedroomsOutsideCentre: {},
		attributeApartment3BedroomsCityCentre:    {},
		attributeMealInexpensiveRestaurant:       {},
		attributePricePerSqFtOutsideCentre:       {},
		attributePricePerSqFtCityCentre:          {},
		attributeCigarettes:                      {},
		attributeOranges:                         {},
		attributeTomato:                          {},
		attributeWaterBottle:                     {},
		attributePairOfJeans:                     {},
	}
}

func FormatCitiesReport(points []data.Point) []string {
	groupedData := groupDataByCityAndCategory(points)
	var messages []string
	for city, categories := range groupedData {
		messages = append(messages, Bold(city)+" - city report")
		for _, category := range categoryOrder {
			if attrs, exists := categories[category]; exists {
				messages = append(messages, formatCatergoryTitle(category, city))
				messages = append(messages, formatCityAttributes(attrs, getSkipLabels()))
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

	return fmt.Sprintf("%s: %s", Bold("Sources"), strings.Join(names, ", "))
}

func FormatFetchDuration(d time.Duration) string {
	return Italic(fmt.Sprintf("Fetched in %s", d))
}

func FormatProvider() string {
	return "By " + Underline(provider)
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

func formatCityAttributes(attributes []data.Attribute, skipLabels map[string]struct{}) string {
	var result []formatAttributeResult
	var latestTime time.Time
	for _, attr := range attributes {
		if len(attr.Values) == 0 || len(attr.Timestamps) == 0 {
			continue
		}

		if _, skip := skipLabels[attr.Label]; skip {
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

	subgroupMap := make(map[string][]formatAttributeResult)
	for _, r := range result {
		subgroup := r.subgroup
		if subgroup == "" {
			subgroup = noSubgroup
		}
		subgroupMap[subgroup] = append(subgroupMap[subgroup], r)
	}

	for _, subgroupResults := range subgroupMap {
		sort.Slice(subgroupResults, func(i, j int) bool {
			if subgroupResults[i].order == subgroupResults[j].order {
				return subgroupResults[i].attribute < subgroupResults[j].attribute
			}
			return subgroupResults[i].order < subgroupResults[j].order
		})
	}

	var subgroups []string
	for subgroup := range subgroupMap {
		subgroups = append(subgroups, subgroup)
	}
	sort.Strings(subgroups)

	var formattedResult []string
	for i, subgroup := range subgroups {
		var prefix, suffix string
		if i != 0 {
			prefix = "\n"
		}

		if i != len(subgroups)-1 {
			suffix = ""
		}

		if subgroup != noSubgroup {
			formattedResult = append(formattedResult, prefix+Underline(subgroup)+suffix)
		}

		for _, r := range subgroupMap[subgroup] {
			formattedResult = append(formattedResult, r.attribute)
		}
	}

	if !latestTime.IsZero() {
		formattedResult = append(formattedResult, Italic(fmt.Sprintf("updated at %s", latestTime.Format("2006-01-02 15:04"))))
	}
	return strings.Join(formattedResult, "\n")
}

func formatAttribute(label string, values string, comment string) formatAttributeResult {
	if result, formatted := formatWeatherAttribute(label, values, comment); formatted {
		return result
	}

	if result, formatted := formatWorldBankAttribute(label, values, comment); formatted {
		return result
	}

	if result, formatted := formatAirQualityAttribute(label, values, comment); formatted {
		return result
	}

	if result, formatted := formatPriceAttribute(label, values, comment); formatted {
		return result
	}

	// not a special case, just return the label and values
	return formatAttributeResult{
		attribute: formatSingleAttribute(label, values, "", comment),
	}
}

func formatSingleAttribute(label, values, emoji, comment string) string {
	if comment != "" {
		comment = " (" + comment + ")"
	}
	return fmt.Sprintf("%s: %s%s %s", Capitalize(label), Bold(values), Italic(comment), emoji)
}

func formatCatergoryTitle(category, city string) string {
	switch category {
	case catergoryWeather:
		return formatSingleCategoryTitle("Weather") + " " + weatherEmoji
	case catergoryAirQuality:
		return formatSingleCategoryTitle("Air Quality") + " " + airEmoji
	case categoryWorldBank:
		title := "National Economy"
		flag, ok := citycountry.GetFlagByCity(city)
		if ok {
			return formatSingleCategoryTitle(title) + " " + flag
		}
		return formatSingleCategoryTitle(title) + " " + stonksEmoji
	case categoryPrices:
		return formatSingleCategoryTitle("Prices") + " " + dollarEmoji
	default:
		return formatSingleCategoryTitle(category)
	}
}

func formatSingleCategoryTitle(category string) string {
	return Bold(Upper((category)))
}

func FormatCityList(cities []string) string {
	var formattedCities []string
	for _, c := range cities {
		formattedCities = append(formattedCities,
			Capitalize(c),
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
