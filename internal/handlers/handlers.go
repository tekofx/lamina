package handlers

import (
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/tekofx/lamina/internal/logger"
	"github.com/tekofx/lamina/internal/utils"
)

func AddHandlers(bh *th.BotHandler, bot *telego.Bot) {
	newMessage(bh, bot)

}

// func newStickerMessage(bh *th.BotHandler, bot *telego.Bot) {
// 	bh.Handle(func(ctx *th.Context, update telego.Update) error {
// 		logger.Log(fmt.Sprintf("Admission: Left member %s", update.Message.LeftChatMember.Username))
// 		return nil
// 	}, th.AnyMessage())
// }

func newMessage(bh *th.BotHandler, bot *telego.Bot) {
	bh.Handle(func(ctx *th.Context, update telego.Update) error {

		if len(update.Message.Photo) > 0 {
			logger.Log("Photo")
			fileID := update.Message.Photo[len(update.Message.Photo)-1].FileID // Highest quality photo
			err := utils.DownloadFile(fileID, "photo.jpg", update.Context(), bot)

			if err != nil {
				log.Fatal(err)
			}
		}

		if update.Message.Document != nil {
			logger.Log("Document")
			err := utils.DownloadFile(update.Message.Document.FileID, update.Message.Document.FileName, update.Context(), bot)

			if err != nil {
				log.Fatal(err)
			}
		}

		if update.Message.Sticker != nil {
			logger.Log("Sticker")
			if update.Message.Sticker.IsAnimated || update.Message.Sticker.IsVideo {
				logger.Error("Must be still sticker")
			}

			err := utils.DownloadFile(update.Message.Sticker.FileID, "sticker.webp", update.Context(), bot)

			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}, th.AnyMessage())
}
