package bot

import (
	"context"
	"sort"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) handleStart(c tele.Context) error {
	return b.handleHelp(c)
}

func (b *Bot) handleHelp(c tele.Context) error {
	return c.Send(b.getHelpMessage())
}

func (b *Bot) handleCitiesCallback(c tele.Context) error {
	citiesData, err := b.fetchCitiesData(fetchCitiesDataInput{
		city: c.Data(),
	})
	if err != nil {
		return err
	}

	cityReports := FormatCitiesReport(citiesData.points)
	report := strings.Join([]string{
		strings.Join(cityReports, "\n\n"),
		FormatMessageFooter(citiesData.sources, citiesData.start),
	}, "\n\n")
	return c.Send(report, &tele.SendOptions{ParseMode: tele.ModeHTML})
}

func (b *Bot) handleCities(c tele.Context) error {
	tags, err := b.placesClient.GetTags(context.TODO())
	if err != nil {
		return err
	}

	var cities []string
	for _, tag := range tags {
		if tag.Label == TagCity {
			cities = append(cities, tag.Value)
		}
	}

	sort.Strings(cities)

	r := &tele.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	const buttonsPerRow = 3
	var btns []tele.Btn
	for _, city := range cities {
		btns = append(btns, r.Data(capitalize(city), callbackCity, city))
	}

	var rows []tele.Row
	for i := 0; i < len(btns); i += buttonsPerRow {
		end := i + buttonsPerRow
		if end > len(btns) {
			end = len(btns)
		}
		rows = append(rows, r.Row(btns[i:end]...))
	}
	r.Inline(rows...)

	return c.Send("Choose a city to get report:", &tele.SendOptions{
		ReplyMarkup: r,
	})
}
