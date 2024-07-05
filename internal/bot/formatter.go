package bot

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/data"
)

const (
	happyEmoji      = "😊"
	loveFaceEmoji   = "😍"
	heartsFaceEmoji = "🥰"
	satisfiedEmoji  = "😌"
	happyCatEmoji   = "😸"
	partyEmoji      = "🥳"
	neutralEmoji    = "😐"
	sadEmoji        = "😞"
	terrorEmoji     = "😱"
	hotEmoji        = "🥵"
	veryHotEmoji    = hotEmoji + terrorEmoji
	coldEmoji       = "🥶"
	veryColdEmoji   = coldEmoji + terrorEmoji
	wetEmoji        = "💦"
	dryEmoji        = "🌵"
	normalHumEmoji  = "🌿"
	pressureEmoji   = "🌀"
	cloudyEmoji     = "☁️"
	clearEmoji      = "🌞"

	attributeTemperature = "temperature"
	attributeHumidity    = "humidity"
	attributePressure    = "pressure"
	attributeDescription = "description"
)

func FormatCityReport(point data.Point) string {
	city := getCityFromTags(point.Tags)
	attributesMessage := formatCityAttributes(point.Attributes)
	return fmt.Sprintf("%s\n%s", bold(underline(city)), attributesMessage)
}

func FormatCitiesReport(points []data.Point) []string {
	var messages []string
	sort.Slice(points, func(i, j int) bool {
		return getCityFromTags(points[i].Tags) < getCityFromTags(points[j].Tags) // Сортировка по городу
	})
	for _, point := range points {
		messages = append(messages, FormatCityReport(point))
	}
	return messages
}

func FormatSources(sources []data.Source) string {
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

func getCityFromTags(tags []data.Tag) string {
	for _, tag := range tags {
		if tag.Label == "city" {
			return tag.Value
		}
	}
	return ""
}

func formatCityAttributes(attributes []data.Attribute) string {
	var result []string
	var latestTime time.Time

	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Label < attributes[j].Label
	})

	for _, attr := range attributes {
		values := strings.Join(attr.Values, ", ")
		lastTimestamp := attr.Timestamps[len(attr.Timestamps)-1]
		if lastTimestamp.After(latestTime) {
			latestTime = lastTimestamp
		}
		attributeString := formatAttribute(attr.Label, values)
		result = append(result, attributeString)
	}

	if !latestTime.IsZero() {
		result = append(result, italic(fmt.Sprintf("updated at %s", latestTime.Format("2006-01-02"))))
	}
	return strings.Join(result, "\n")
}

func formatAttribute(label string, values string) string {
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
		}
	}
	return fmt.Sprintf("%s: %s %s", bold(strings.Trim(label, " ")), values, emoji)
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
