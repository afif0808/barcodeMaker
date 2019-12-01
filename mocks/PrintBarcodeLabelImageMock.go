package mocks

import (
	"barcodeMaker/model"
	"image"
)

func PrintBarcodeLabelImageMock(err error) model.PrintBarcodeLabelImage {
	return func(image.Image, int64) error {
		return err
	}
}
