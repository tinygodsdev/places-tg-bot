package bot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/data"
	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/datasdk/pkg/server"
	"github.com/tinygodsdev/places-tg-bot/internal/config"
	tele "gopkg.in/telebot.v3"
)

const (
	CommandStart  = "/start"
	CommandReport = "/report"
)

type Bot struct {
	cfg          *config.Config
	placesClient server.Client
	log          logger.Logger
	t            *tele.Bot
}

func New(cfg *config.Config, placesClient server.Client, l logger.Logger) (*Bot, error) {
	settings := tele.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	t, err := tele.NewBot(settings)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		cfg:          cfg,
		placesClient: placesClient,
		t:            t,
		log:          l,
	}

	b.linkHandlers()
	return b, nil
}

func (b *Bot) Start() {
	b.t.Start()
}

func (b *Bot) Stop() {
	b.t.Stop()
}

func (b *Bot) linkHandlers() {
	b.t.Handle("/start", b.handleStart)
	b.t.Handle("/report", b.handleReport)
}

func (b *Bot) handleStart(c tele.Context) error {
	return c.Send("Hello! I'm a bot that can help you find places!")
}

func (b *Bot) handleReport(c tele.Context) error {
	b.log.Info("got report command", fmt.Sprintf("user: %+v", c.Recipient()))
	if err := c.Send("Fetching today's data about locations"); err != nil {
		return err
	}

	start := time.Now()
	points, err := b.placesClient.GetPoints(
		context.TODO(),
		data.Filter{
			From: time.Now().Add(-24 * time.Hour),
			To:   time.Now(),
		},
		data.Group{
			TagLabels: []string{"city"},
		},
	)
	if err != nil {
		if err := c.Send("Failed to fetch data" + err.Error()); err != nil {
			return err
		}
	}

	sources, err := b.placesClient.GetSources(context.TODO())
	if err != nil {
		if err := c.Send("Failed to fetch sources" + err.Error()); err != nil {
			return err
		}
	}

	sourcesReport := FormatSources(sources)

	cityReports := FormatCitiesReport(points)
	report := strings.Join([]string{
		"Here is the data:",
		strings.Join(cityReports, "\n\n"),
		sourcesReport,
		FormatFetchDuration(time.Since(start)),
	}, "\n\n")
	return c.Send(report, &tele.SendOptions{ParseMode: tele.ModeHTML})
}
