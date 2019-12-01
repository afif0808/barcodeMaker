package model

import (
	"image"
	"net/http"
)

type HandlePostData func(http.Request) (map[string]interface{}, error)
type MakeBarcodeLabelImage func(BarcodeLabel) (image.Image, error)
type PrintBarcodeLabelImage func(barcodeLabelImage image.Image, quantity int64) error
type MakeTextImage func(text string, textImageSetting TextImageSetting) (image.Image, error)
type MakeBarcodeImage func(barcodeData interface{}, width, height int) (image.Image, error)
type ArrangeImage func(MainImageSetting, Images) (image.Image, error)
