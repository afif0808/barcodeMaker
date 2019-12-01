package mocks

import (
	"barcodeMaker/model"
	"image"
)

func MakeTextImageMock(img image.Image, err error) model.MakeTextImage {
	return func(model.TextImageSetting) (image.Image, error) {
		return img, err
	}
}
