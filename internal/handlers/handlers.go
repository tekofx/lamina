package handlers

import (
	"log"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/tekofx/lamina/internal/logger"
	"github.com/tekofx/lamina/internal/utils"
)

func AddHandlers(bh *th.BotHandler, bot *telego.Bot) {
	newMessage(bh, bot)

}

func newMessage(bh *th.BotHandler, bot *telego.Bot) {
	bh.Handle(func(ctx *th.Context, update telego.Update) error {

		var fileID string
		var fileExtension string

		if len(update.Message.Photo) > 0 {
			fileID = update.Message.Photo[len(update.Message.Photo)-1].FileID // Highest quality photo
			fileExtension = "jpg"
		}

		if update.Message.Document != nil {
			fileID = update.Message.Document.FileID
			fileExtension = "png"
		}

		if update.Message.Sticker != nil {
			if update.Message.Sticker.IsAnimated || update.Message.Sticker.IsVideo {
				logger.Error("Must be still sticker")
			}
			fileID = update.Message.Sticker.FileID
			fileExtension = "webp"
		}

		bot.SendMessage(update.Context(), &telego.SendMessageParams{
			ChatID: update.Message.Chat.ChatID(),
			Text:   "Converting image...",
		})
		logger.Log("test")

		downloadedImageFileName, err := utils.DownloadFile(fileID, fileExtension, update.Context(), bot)
		if err != nil {
			log.Fatal(err)
		}

		convertedImageFilename, err := utils.ConvertToImage(*downloadedImageFileName)
		if err != nil {
			log.Fatal(err)
		}

		file, _ := os.Open(*convertedImageFilename)
		defer file.Close()

		params := tu.Document(tu.ID(update.Message.Chat.ID), tu.File(file))
		_, err = bot.SendDocument(ctx, params)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Remove(*downloadedImageFileName)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Remove(*convertedImageFilename)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}, th.AnyMessage())
}
