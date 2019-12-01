package services

import (
	"barcodeMaker/model"
	"fmt"
	"image"
)

func MakeBarcodeLabelImage(
	imageSetting model.ImageSettings,
	barcodeLabelSettings model.BarcodeLabelSettings,
	barcodeLabelPosition model.BarcodeLabelPositions,
	makeTextImage model.MakeTextImage,
	makeBarcodeImage model.MakeBarcodeImage,
	arrangeImage model.ArrangeImage,
) func(barcodeLabel model.BarcodeLabel) (image.Image, error) {

	return func(barcodeLabel model.BarcodeLabel) (image.Image, error) {
		if barcodeLabel == (model.BarcodeLabel{}) {
			return nil, fmt.Errorf("MakeBarcodeLabelImage() : expecting 'BarcodeLabel' to be not empty")
		}

		barcodeImage, makeBarcodeImageErr := makeBarcodeImage(
			barcodeLabel.Code,
			barcodeLabelSettings.BarcodeSetting.Width,
			barcodeLabelSettings.BarcodeSetting.Height)

		codeImage, makeCodeImageErr := makeTextImage(barcodeLabel.Code, barcodeLabelSettings.CodeSetting)
		if makeCodeImageErr != nil {
			return nil, fmt.Errorf("MakeBarcodeLabelImage() : %v", makeCodeImageErr)
		}

		modelImage, makeModelImageErr := makeTextImage(barcodeLabel.Model, barcodeLabelSettings.ModelSetting)
		if makeModelImageErr != nil {
			return nil, fmt.Errorf("MakeBarcodeLabelImage() : %v", makeModelImageErr)
		}

		return nil, nil
	}

}
