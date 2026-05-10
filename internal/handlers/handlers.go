package handlers

import (
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
		}

		if update.Message.Document != nil {
			logger.Log("Document")
		}

		if update.Message.Sticker != nil {
			logger.Log("Sticker")
			fileID := update.Message.Sticker.FileID
			file, err := bot.GetFile(update.Context(), &telego.GetFileParams{FileID: fileID})
			if err != nil {
				log.Fatal(err)
			}

			if update.Message.Sticker.IsAnimated || update.Message.Sticker.IsVideo {
				logger.Error("Must be still sticker")
			}

			fileUrl := bot.FileDownloadURL(file.FilePath)
			logger.Log(file.FileID)

			err = DownloadFile("sticker.webp", fileUrl)

			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}, th.AnyMessage())
}

func DownloadFile(filepath string, url string) error {
	// Get the response from the URL
	resp, err := http.Get(url)
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
