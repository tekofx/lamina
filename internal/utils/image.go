package utils

import (
	"github.com/sunshineplan/imgconv"
)

func ConvertToSticker(filepath string) error {
	src, err := imgconv.Open(filepath)
	if err != nil {
		return err
	}

	resized := imgconv.Resize(src, &imgconv.ResizeOption{Width: 512})

	if err := imgconv.Save("sticker.webp", resized, &imgconv.FormatOption{Format: imgconv.WEBP}); err != nil {
		return err
	}

	return nil
}

func ConvertToImage(filepath string) error {
	src, err := imgconv.Open(filepath)
	if err != nil {
		return err
	}

	if err := imgconv.Save("sticker.png", src, &imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
		return err
	}

	return nil

}
