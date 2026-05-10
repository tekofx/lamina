package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sunshineplan/imgconv"
	"github.com/tekofx/lamina/internal/config"
	"github.com/tekofx/lamina/internal/logger"
)

func ConvertToImage(filepath string) (*string, error) {
	src, err := imgconv.Open(filepath)
	if err != nil {
		return nil, err
	}

	logger.Log(src.Bounds().Dx())
	logger.Log(src.Bounds().Dy())

	filename := fmt.Sprintf("%s/%s.png", config.Conf.MediaFolder, uuid.NewString())

	if src.Bounds().Dx() > src.Bounds().Dy() {
		src = imgconv.Resize(src, &imgconv.ResizeOption{Width: 512})
	} else {
		src = imgconv.Resize(src, &imgconv.ResizeOption{Height: 512})
	}

	if err := imgconv.Save(
		filename,
		src,
		&imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
		return nil, err
	}

	return &filename, nil

}
