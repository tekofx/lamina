package handlers

import (
	"fmt"
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
			fmt.Println(update.Message.Document.MimeType)

			fileExtension = "png"
		}

		if update.Message.Sticker != nil {
			if update.Message.Sticker.IsAnimated || update.Message.Sticker.IsVideo {
				logger.Error("Must be still sticker")
				return nil
			}
			fileID = update.Message.Sticker.FileID
			fileExtension = "webp"
		}

		bot.SendMessage(update.Context(), &telego.SendMessageParams{
			ChatID: update.Message.Chat.ChatID(),
			Text:   "Converting image...",
		})

		downloadedImageFileName, err := utils.DownloadFile(fileID, fileExtension, update.Context(), bot)
		if err != nil {
			logger.Log(err)
			bot.SendMessage(update.Context(), &telego.SendMessageParams{
				ChatID: update.Message.Chat.ChatID(),
				Text:   "Error downloading image. Check the image is not larger than 20 mb",
			})
			return nil
		}

		convertedImageFilename, err := utils.ConvertToImage(*downloadedImageFileName)
		if err != nil {
			logger.Log(err)
			return nil
		}

		file, _ := os.Open(*convertedImageFilename)
		defer file.Close()

		logger.Log("Sending image")
		params := tu.Document(tu.ID(update.Message.Chat.ID), tu.File(file))
		_, err = bot.SendDocument(ctx, params)
		if err != nil {
			logger.Log(err)
			return nil
		}
		logger.Log("Image sent")

		err = os.Remove(*downloadedImageFileName)
		if err != nil {
			logger.Log(err)
			return nil
		}

		err = os.Remove(*convertedImageFilename)
		if err != nil {
			logger.Log(err)
			return nil
		}

		return nil
	}, th.AnyMessage())
}
