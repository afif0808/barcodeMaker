package services

import (
	"barcodeMaker/model"
	"image"

	cairo "github.com/ungerik/go-cairo"
)

func MakeBarcodeLabelImage2(
	mainImageSetting model.MainImageSetting,
	barcodeSetting model.BarcodeSetting,
	textImageSettings model.TextImageSettings,
	makeBarcodeImage model.MakeBarcodeImage,
) model.MakeBarcodeLabelImage {
	return func(barcodeLabel model.BarcodeLabel) (image.Image, error) {

		makeText := func(
			surface *cairo.Surface,
			text string,
			textImageSetting model.TextImageSetting,
			imagePosition model.ImagePosition,
		) {
			surface.MoveTo(imagePosition.X, imagePosition.Y)
			surface.SelectFontFace(textImageSetting.FontStyle, textImageSetting.FontSlant, textImageSetting.FontWeight)
			surface.SetFontSize(textImageSetting.FontSize)
			surface.SetSourceRGB(0, 0, 0)
			surface.ShowText(text)
		}

		surface := cairo.NewSurface(cairo.FORMAT_ARGB32, mainImageSetting.Width, mainImageSetting.Height)
		surface.Rectangle(0, 0, float64(mainImageSetting.Width), float64(mainImageSetting.Height))
		surface.SetSourceRGB(100, 100, 100)
		surface.Fill()

		codeImage := model.TextImageSetting{
			FontSize:  20,
			FontStyle: "Lato",
		}

		makeText(surface, barcodeLabel.Code, codeImage, model.ImagePosition{X: 0, Y: 800})

		//Make barcode image
		barcodeImage, makeBarcodeImageErr := makeBarcodeImage(
			barcodeLabel.Code,
			barcodeSetting.Width, barcodeSetting.Height)
		if makeBarcodeImageErr != nil {
			return nil, makeBarcodeImageErr
		}
		// barcodeImagePositioned := barcodeImage.Bounds().Add(image.Pt(barcodeSetting.X, barcodeSetting.Y))
		surface.SetImage(barcodeImage)
		// surface.SetFontSize(size)
		// surface.
		surface.WriteToPNG("/home/afif0808/Music/test.png")

		return nil, nil
	}
}
