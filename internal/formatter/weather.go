package formatter

import (
	"strconv"
	"strings"
)

func formatWeatherAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int
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
				emoji = RandomHappyEmoji()
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
				emoji = RandomHappyEmoji()
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
		} else if strings.Contains(values, "rain") {
			emoji = rainEmoji
		} else if strings.Contains(values, "snow") {
			emoji = snowEmoji
		}
	default:
		return formatAttributeResult{}, false
	}

	return formatAttributeResult{
		attribute: formatSingleAttribute(label, values, emoji, comment),
		subgroup:  subgroup,
		order:     order,
	}, true
}
