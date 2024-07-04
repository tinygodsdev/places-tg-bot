package bot

import (
	"fmt"
	"strings"

	"github.com/tinygodsdev/datasdk/pkg/data"
)

func FormatCityReport(point data.Point) string {
	city := ""
	for _, tag := range point.Tags {
		if tag.Label == "city" {
			city = tag.Value
			break
		}
	}

	attributesMessage := formatCityAttributes(point.Attributes)
	return fmt.Sprintf("%s\n%s", city, attributesMessage)
}

func FormatCitiesReport(points []data.Point) []string {
	var messages []string
	for _, point := range points {
		messages = append(messages, FormatCityReport(point))
	}
	return messages
}

func formatCityAttributes(attributes []data.Attribute) string {
	var result []string
	for _, attr := range attributes {
		values := strings.Join(attr.Values, ", ")
		lastTimestamp := attr.Timestamps[len(attr.Timestamps)-1]
		attributeString := fmt.Sprintf(
			"%s: %s (updated at %s)",
			strings.Trim(attr.Label, " "),
			values,
			lastTimestamp.Format("2006-01-02"),
		)
		result = append(result, attributeString)
	}
	return strings.Join(result, "\n")
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

	return fmt.Sprintf("Sources: %s", strings.Join(names, ", "))
}
