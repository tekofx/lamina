package main

import (
	"context"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/tekofx/lamina/internal/config"
	"github.com/tekofx/lamina/internal/handlers"
	"github.com/tekofx/lamina/internal/logger"
)

func main() {

	config.InitializeConfig()

	bot, botErr := telego.NewBot(config.Conf.Token)

	if botErr != nil {
		logger.Fatal(botErr)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(context.Background(), nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates) // Stop handling updates
	defer func() { _ = bh.Stop() }()

	handlers.AddHandlers(bh, bot)

	logger.Log("Bot started")
	bh.Start()

}
