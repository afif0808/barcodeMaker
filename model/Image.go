package model

import "image"

type Image struct {
	image.Image
	X, Y float64
}

type Images struct {
	BarcodeImage, CodeImage, ModelImage, StyleNameImage Image
	VariantImage, PriceImage, SizeImage, MaterialImage  Image
}
