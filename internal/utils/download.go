package utils

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
)

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
