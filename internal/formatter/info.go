package formatter

import "github.com/tinygodsdev/cities/pkg/cities"

func formatGeneralInfoAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int

	switch label {
	case cities.AttributePopulationTotal:
		label = cities.AttributePopulationTotalShort
		order = 100
	case cities.AttributePopulationDensity:
		label = cities.AttributePopulationDensityShort
		values += " people/km²"
		order = 200
	case cities.AttributeAreaTotal:
		label = cities.AttributeAreaTotalShort
		values += " km²"
	case cities.AttributeElevation:
		label = cities.AttributeElevationShort
		values += " m"
		order = 300
	case cities.AttributeTimezone:
		label = cities.AttributeTimezoneShort
	default:
		return formatAttributeResult{}, false
	}

	return formatAttributeResult{
		attribute: formatSingleAttribute(label, values, emoji, comment),
		subgroup:  subgroup,
		order:     order,
	}, true
}
