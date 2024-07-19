package formatter

import (
	"strconv"

	"github.com/tinygodsdev/cities/pkg/cities"
)

func formatAirQualityAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int
	value, err := strconv.ParseFloat(values, 64)
	if err != nil {
		return formatAttributeResult{}, false
	}

	switch label {
	case cities.AttributeCo:
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
	case cities.AttributeNo2:
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
	case cities.AttributeO3:
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
	case cities.AttributePm10:
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
	case cities.AttributePm25:
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
	case cities.AttributeSo2:
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
		return formatAttributeResult{}, false
	}

	return formatAttributeResult{
		attribute: formatSingleAttribute(label, values, emoji, comment),
		subgroup:  subgroup,
		order:     order,
	}, true
}
