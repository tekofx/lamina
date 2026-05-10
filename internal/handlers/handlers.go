package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/tekofx/lamina/internal/logger"
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
			err := DownloadFile(fileID, "photo.jpg", update.Context(), bot)

			if err != nil {
				log.Fatal(err)
			}
		}

		if update.Message.Document != nil {
			logger.Log("Document")
			err := DownloadFile(update.Message.Document.FileID, update.Message.Document.FileName, update.Context(), bot)

			if err != nil {
				log.Fatal(err)
			}
		}

		if update.Message.Sticker != nil {
			logger.Log("Sticker")
			if update.Message.Sticker.IsAnimated || update.Message.Sticker.IsVideo {
				logger.Error("Must be still sticker")
			}

			err := DownloadFile(update.Message.Sticker.FileID, "sticker.webp", update.Context(), bot)

			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}, th.AnyMessage())
}

func DownloadFile(fileID string, filepath string, ctx context.Context, bot *telego.Bot) error {

	file, err := bot.GetFile(ctx, &telego.GetFileParams{FileID: fileID})
	if err != nil {
		log.Fatal(err)
	}

	fileUrl := bot.FileDownloadURL(file.FilePath)

	// Get the response from the URL
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the response body to the file
	_, err = io.Copy(out, resp.Body)
	return err

}
