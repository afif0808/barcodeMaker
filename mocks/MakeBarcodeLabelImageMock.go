package mocks

import (
	"barcodeMaker/model"
	"image"
)

func MakeBarcodeLabelImageMock(img image.Image, err error) model.MakeBarcodeLabelImage {
	return func(model.BarcodeLabel) (image.Image, error) {
		return img, err
	}
}
