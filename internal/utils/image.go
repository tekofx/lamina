package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sunshineplan/imgconv"
	"github.com/tekofx/lamina/internal/config"
)

func ConvertToImage(filepath string) (*string, error) {
	src, err := imgconv.Open(filepath)
	if err != nil {
		return nil, err
	}

	filename := fmt.Sprintf("%s/%s.png", config.Conf.MediaFolder, uuid.NewString())

	if err := imgconv.Save(
		filename,
		src,
		&imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
		return nil, err
	}

	return &filename, nil

}
