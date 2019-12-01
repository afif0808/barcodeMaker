package mocks

import "image"

func MakeBarcodeImageMock(img image.Image, err error) func(interface{}) (image.Image, error) {
	return func(interface{}) (image.Image, error) {
		return img, err
	}
}
