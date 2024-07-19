package bot

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/tinygodsdev/cities/pkg/cities"
	"github.com/tinygodsdev/datasdk/pkg/bot/format"
	"github.com/tinygodsdev/datasdk/pkg/data"
	"github.com/tinygodsdev/places-tg-bot/internal/formatter"
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

	cityReports := formatter.FormatCitiesReport(citiesData.points)
	report := strings.Join([]string{
		strings.Join(cityReports, "\n\n"),
		formatter.FormatMessageFooter(citiesData.sources, citiesData.start),
	}, "\n\n")

	// if err := c.Delete(); err != nil {
	// 	b.log.Error("failed to delete message", "error", err)
	// }

	return c.Send(report, &tele.SendOptions{ParseMode: tele.ModeHTML})
}

func (b *Bot) handleCities(c tele.Context) error {
	f := format.New(format.ModeHTML)
	tags, err := b.placesClient.GetTags(context.TODO(), data.Filter{
		From: time.Now().Add(-24 * time.Hour),
		To:   time.Now(),
	})
	if err != nil {
		return err
	}

	var citiesStrs []string
	for _, tag := range tags {
		if tag.Label == cities.TagCity {
			citiesStrs = append(citiesStrs, tag.Value)
		}
	}

	sort.Strings(citiesStrs)

	r := &tele.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	const buttonsPerRow = 3
	var btns []tele.Btn
	for _, city := range citiesStrs {
		btns = append(btns, r.Data(f.Capitalize(city), callbackCity, city))
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

// handle random message
func (b *Bot) handleRandom(c tele.Context) error {
	return c.Send(fmt.Sprintf("You message is acknowledged. %s Otherwise type /help to get instructions.", suggestionInfo))
}
