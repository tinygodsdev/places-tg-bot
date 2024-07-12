package formatter

func formatGeneralInfoAttribute(label string, values string, comment string) (formatAttributeResult, bool) {
	var emoji string
	var subgroup string
	var order int

	switch label {
	case attributePopulationTotal:
		label = attributePopulationTotalShort
		order = 100
	case attributePopulationDensity:
		label = attributePopulationDensityShort
		values += " people/km²"
		order = 200
	case attributeAreaTotal:
		label = attributeAreaTotalShort
		values += " km²"
	case attributeElevation:
		label = attributeElevationShort
		values += " m"
		order = 300
	case attributeTimezone:
		label = attributeTimezoneShort
	default:
		return formatAttributeResult{}, false
	}

	return formatAttributeResult{
		attribute: formatSingleAttribute(label, values, emoji, comment),
		subgroup:  subgroup,
		order:     order,
	}, true
}
