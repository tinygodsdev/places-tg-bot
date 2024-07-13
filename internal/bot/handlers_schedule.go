package bot

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/data"
	"github.com/tinygodsdev/places-tg-bot/internal/formatter"
	"github.com/tinygodsdev/places-tg-bot/internal/util"
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) handleScheduleCallback(c tele.Context) error {
	userID := c.Sender().ID
	schedule := c.Data()

	user, err := b.userStore.GetUserByID(context.TODO(), fmt.Sprint(userID))
	if err != nil {
		b.log.Error("failed to get user", "error", err, "user_id", userID)
		return c.Send("Failed to get user profile")
	}

	user.Preferences.ReportSchedule = schedule

	ctx := context.TODO()
	if err := b.userStore.SaveOrUpdateUser(ctx, user); err != nil {
		b.log.Error("failed to save user", "error", err, "user_id", userID)
		return c.Send("Failed to save user profile")
	}

	if err := c.Delete(); err != nil {
		b.log.Error("failed to delete message", "error", err)
	}

	return c.Send(fmt.Sprintf("Reporting frequency set to %s", schedule))
}

func (b *Bot) handleSchedule(c tele.Context) error {
	userID := c.Sender().ID
	user, err := b.userStore.GetUserByID(context.TODO(), fmt.Sprint(userID))
	if err != nil {
		b.log.Error("failed to get user", "error", err, "user_id", userID)
		return c.Send("Failed to get user profile")
	}

	r := &tele.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	schedules := b.getSchedules()
	const buttonsPerRow = 2
	var btns []tele.Btn
	for _, schedule := range schedules {
		scheduleText := schedule
		if user.Preferences.ReportSchedule == schedule {
			scheduleText = fmt.Sprintf("%s %s", schedule, selectedEmoji)
		}

		btns = append(btns, r.Data(scheduleText, callbackSchedule, schedule))
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

	return c.Send("Choose reporting frequency:", &tele.SendOptions{
		ReplyMarkup: r,
	})
}

func (b *Bot) handleReportCities(c tele.Context) error {
	userID := c.Sender().ID
	user, err := b.userStore.GetUserByID(context.TODO(), fmt.Sprint(userID))
	if err != nil {
		b.log.Error("failed to get user", "error", err, "user_id", userID)
		return c.Send("Failed to get user profile")
	}

	tags, err := b.placesClient.GetTags(context.TODO(), data.Filter{
		From: time.Now().Add(-24 * time.Hour),
		To:   time.Now(),
	})
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
		cityText := formatter.Capitalize(city)
		if util.ContainsString(user.Preferences.ReportCities, city) {
			cityText = fmt.Sprintf("%s %s", cityText, selectedEmoji)
		}

		btns = append(btns, r.Data(formatter.Capitalize(cityText), callbackReportCities, city))
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

	return c.Send("Select cities for regular reporting:", &tele.SendOptions{
		ReplyMarkup: r,
	})
}

func (b *Bot) handleReportCitiesCallback(c tele.Context) error {
	userID := c.Sender().ID
	city := c.Data()

	user, err := b.userStore.GetUserByID(context.TODO(), fmt.Sprint(userID))
	if err != nil {
		b.log.Error("failed to get user", "error", err, "user_id", userID)
		return c.Send("Failed to get user profile")
	}

	// Добавить или удалить город из списка предпочтений
	if util.ContainsString(user.Preferences.ReportCities, city) {
		// Удалить город
		user.Preferences.ReportCities = util.RemoveString(user.Preferences.ReportCities, city)
	} else {
		// Добавить город
		user.Preferences.ReportCities = append(user.Preferences.ReportCities, city)
	}

	ctx := context.TODO()
	if err := b.userStore.SaveOrUpdateUser(ctx, user); err != nil {
		b.log.Error("failed to save user", "error", err, "user_id", userID)
		return c.Send("Failed to save user profile")
	}

	if err := c.Delete(); err != nil {
		b.log.Error("failed to delete message", "error", err)
	}

	// Обновить сообщение с выбором
	return b.handleReportCities(c)
}
