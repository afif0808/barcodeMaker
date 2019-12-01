package services

import (
	"fmt"
	"image"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func MakeBarcodeImage(content interface{}, width, height int) (image.Image, error) {
	var barcodeImg barcode.Barcode
	barcodeImg, code128EncodeErr := code128.Encode(fmt.Sprint(content))
	if code128EncodeErr != nil {
		return nil, code128EncodeErr
	}
	barcodeImgScaled, scaleBarcodeErr := barcode.Scale(barcodeImg, width, height)
	if scaleBarcodeErr != nil {
		return nil, scaleBarcodeErr
	}
	return barcodeImgScaled, nil
}
