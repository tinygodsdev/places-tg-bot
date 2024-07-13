package bot

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/datasdk/pkg/server"
	"github.com/tinygodsdev/places-tg-bot/internal/config"
	"github.com/tinygodsdev/places-tg-bot/internal/formatter"
	"github.com/tinygodsdev/places-tg-bot/internal/user"
	tele "gopkg.in/telebot.v3"
)

const (
	CommandStart        = "/start"
	CommandCities       = "/cities"
	CommandHelp         = "/help"
	CommandSchedule     = "/schedule"
	CommandReportCities = "/report_cities"

	TagCategory = "category"
	TagCity     = "city"

	callbackCity         = "city_callback"
	callbackSchedule     = "schedule_callback"
	callbackReportCities = "report_cities_callback"

	actionUnknown = "unknown"

	botBioEn       = "Insights on cities with large Russian expat communities. Perfect for potential movers or the curious!"
	suggestionInfo = "If you want to suggest more cities or something else, please send it to the bot."

	scheduleMinutely   = "@minutely" // for testing
	scheduleHourly     = "@hourly"
	scheduleDaily      = "@daily"
	scheduleWeekly     = "@weekly"
	scheduleMonthly    = "@monthly"
	scheduleBiannually = "@biannually"

	selectedEmoji = "✅"
)

type Bot struct {
	cfg          *config.Config
	placesClient server.Client
	log          logger.Logger
	t            *tele.Bot
	commands     []tele.Command
	userStore    user.Storage
	mu           sync.Mutex
}

func New(
	cfg *config.Config,
	placesClient server.Client,
	l logger.Logger,
	userStore user.Storage,
) (*Bot, error) {
	settings := tele.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	t, err := tele.NewBot(settings)
	if err != nil {
		return nil, err
	}

	var commands []tele.Command = []tele.Command{
		{Text: CommandStart, Description: "Start the bot"},
		{Text: CommandCities, Description: "Get available cities list"},
		{Text: CommandHelp, Description: "Get bot instructions"},
		{Text: CommandSchedule, Description: "Set a schedule to get city insights"},
		{Text: CommandReportCities, Description: "Select cities for regular reports"},
	}

	b := &Bot{
		cfg:          cfg,
		placesClient: placesClient,
		t:            t,
		log:          l,
		commands:     commands,
		userStore:    userStore,
	}

	b.linkHandlers()
	return b, nil
}

func (b *Bot) Setup() error {
	b.log.Info("setting commands")
	err := b.t.SetCommands(b.commands)
	if err != nil {
		return err
	}

	b.log.Info("setting in-chat description")
	err = b.t.SetMyDescription(b.getHelpMessage(), "en")
	if err != nil {
		return err
	}

	b.log.Info("setting bio")
	err = b.t.SetMyShortDescription(botBioEn, "en")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) Start() {
	b.t.Start()
}

func (b *Bot) Stop() {
	b.t.Stop()
}

func (b *Bot) linkHandlers() {
	b.t.Handle(CommandStart, b.getHandler(CommandStart, b.handleStart))
	b.t.Handle(CommandCities, b.getHandler(CommandCities, b.handleCities))
	b.t.Handle(CommandHelp, b.getHandler(CommandHelp, b.handleHelp))
	b.t.Handle(CommandSchedule, b.getHandler(CommandSchedule, b.handleSchedule))
	b.t.Handle(CommandReportCities, b.getHandler(CommandReportCities, b.handleReportCities))

	// callbacks
	b.t.Handle(&tele.InlineButton{Unique: callbackCity}, b.getHandler(callbackCity, b.handleCitiesCallback))
	b.t.Handle(&tele.InlineButton{Unique: callbackSchedule}, b.getHandler(callbackSchedule, b.handleScheduleCallback))
	b.t.Handle(&tele.InlineButton{Unique: callbackReportCities}, b.getHandler(callbackReportCities, b.handleReportCitiesCallback))

	b.t.Handle(tele.OnText, b.getHandler(actionUnknown, b.handleRandom))
}

func (b *Bot) getHandler(name string, fn func(tele.Context) error) tele.HandlerFunc {
	return func(c tele.Context) error {
		requestID := uuid.New().String()
		start := time.Now()
		var unique string
		if c.Callback() != nil {
			unique = c.Callback().Unique
		}

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if c.Sender() == nil {
				b.log.Error("sender is nil", "callback", unique)
				return
			}

			b.mu.Lock()
			defer b.mu.Unlock()

			usr, err := b.userStore.GetUserByID(ctx, fmt.Sprint(c.Sender().ID))
			if err != nil {
				b.log.Error("failed to get user", "error", err, "request_id", requestID)
			}
			if usr == nil {
				if err := b.userStore.SaveOrUpdateUser(ctx, &user.User{
					ID:        fmt.Sprint(c.Sender().ID),
					FirstName: c.Sender().FirstName,
					LastName:  c.Sender().LastName,
					Username:  c.Sender().Username,
				}); err != nil {
					b.log.Error("failed to save user", "error", err, "request_id", requestID)
					return
				}
			} else {
				usr.FirstName = c.Sender().FirstName
				usr.LastName = c.Sender().LastName
				usr.Username = c.Sender().Username
				if err := b.userStore.SaveOrUpdateUser(ctx, usr); err != nil {
					b.log.Error("failed to update user", "error", err, "request_id", requestID)
					return
				}
			}

			actionLog := &user.UserActionLog{
				UserID: fmt.Sprint(c.Sender().ID),
				Action: name,
				Details: map[string]interface{}{
					"request_id": requestID,
					"callback":   unique,
					"text":       c.Message().Text,
					"data":       c.Data(),
				},
			}

			if err := b.userStore.LogUserAction(ctx, actionLog); err != nil {
				b.log.Error("failed to log user action", "error", err, "request_id", requestID)
				return
			}

			b.log.Info("user input saved", "id", c.Sender().ID)
		}()

		recipient := c.Recipient().Recipient()
		if b.cfg.Env == "dev" {
			recipient = fmt.Sprintf("%+v", c.Recipient())
		}

		b.log.Info(
			"received request",
			"id", requestID,
			"text", c.Message().Text,
			"callback", unique,
			"data", c.Data(),
			"recipient", recipient,
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

func (b *Bot) getHelpMessage() string {
	return strings.Join([]string{
		botBioEn,
		formatter.FormatCommands(b.commands),
		suggestionInfo,
		formatter.FormatDeveloperPlain(),
	}, "\n\n")
}

func (b *Bot) getSchedules() []string {
	return []string{
		scheduleMinutely,
		scheduleHourly,
		scheduleDaily,
		scheduleWeekly,
		scheduleMonthly,
		scheduleBiannually,
	}
}
