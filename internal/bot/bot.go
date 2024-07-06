package bot

import (
	"time"

	"github.com/google/uuid"
	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/datasdk/pkg/server"
	"github.com/tinygodsdev/places-tg-bot/internal/config"
	tele "gopkg.in/telebot.v3"
)

const (
	CommandStart  = "/start"
	CommandReport = "/report"
	CommandCities = "/cities"

	TagCategory = "category"
	TagCity     = "city"

	callbackCity = "city_callback"
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

func (b *Bot) Setup() error {
	b.t.SetCommands([]tele.Command{
		{Text: CommandStart, Description: "Start the bot"},
		{Text: CommandReport, Description: "Get report on all available cities"},
		{Text: CommandCities, Description: "Get available cities list"},
		// TODO: add help and settings commands
	})

	return nil
}

func (b *Bot) Start() {
	b.t.Start()
}

func (b *Bot) Stop() {
	b.t.Stop()
}

func (b *Bot) linkHandlers() {
	b.t.Handle(CommandStart, b.getHandler(b.handleStart))
	b.t.Handle(CommandReport, b.getHandler(b.handleReport))
	b.t.Handle(CommandCities, b.getHandler(b.handleCities))

	// callbacks
	b.t.Handle(&tele.InlineButton{Unique: callbackCity}, b.getHandler(b.handleCitiesCallback))
}

func (b *Bot) getHandler(fn func(tele.Context) error) tele.HandlerFunc {
	return func(c tele.Context) error {
		requestID := uuid.New().String()
		start := time.Now()
		var unique string
		if c.Callback() != nil {
			unique = c.Callback().Unique
		}

		b.log.Info(
			"received request",
			"id", requestID,
			"text", c.Message().Text,
			"callback", unique,
			"data", c.Data(),
			"recipient", c.Recipient().Recipient(),
		)
		defer func() {
			b.log.Info(
				"request processed",
				"id", requestID,
				"elapsed", time.Since(start),
			)
		}()
		return fn(c)
	}
}
