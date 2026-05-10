package main

import (

	"github.com/lamina/internal/config"
	"github.com/lamina/internal/logger"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
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
	bh, _ := th.NewBotHandler(bot, u	// Stop handling updates
	defer func() { _ = bh.Stop() }()



	logger.Log("Bot started")
	bh.Start()

}
