package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/mymmrac/telego"
	"github.com/tekofx/lamina/internal/config"
)

func DownloadFile(fileID string, fileExtension string, ctx context.Context, bot *telego.Bot) (*string, error) {

	file, err := bot.GetFile(ctx, &telego.GetFileParams{FileID: fileID})
	if err != nil {
		return nil, err
	}

	fileUrl := bot.FileDownloadURL(file.FilePath)

	// Get the response from the URL
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create the file
	fileName := fmt.Sprintf("%s/%s.%s", config.Conf.MediaFolder, uuid.NewString(), fileExtension)
	out, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	// Write the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}
	return &fileName, nil

}
