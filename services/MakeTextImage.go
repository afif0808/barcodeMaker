package services

import (
	"barcodeMaker/model"
	"bytes"
	"fmt"
	"image"
	"image/png"

	cairo "github.com/ungerik/go-cairo"
)

func MakeTextImage(text string, textImageSetting model.TextImageSetting) (image.Image, error) {
	// if textImageSetting.FontSize == 0 || textImageSetting.Text == "" || textImageSetting.FontStyle == ""{
	// }
	textImageSettingHeight := int(textImageSetting.FontSize + 10)
	textImageSettingWidth := len(text) * int(textImageSetting.FontSize)

	cairoSurface := cairo.NewSurface(cairo.FORMAT_ARGB32, textImageSettingWidth, textImageSettingHeight)
	cairoSurface.MoveTo(0, textImageSetting.FontSize)
	cairoSurface.SetFontSize(textImageSetting.FontSize)
	cairoSurface.SelectFontFace(textImageSetting.FontStyle, textImageSetting.FontSlant, textImageSetting.FontWeight)
	cairoSurface.ShowText(text)

	txtImageBytes, cairoStatus := cairoSurface.WriteToPNGStream()

	if cairoStatus != cairo.STATUS_SUCCESS {
		return nil, fmt.Errorf("MakeTextImage() : error occured on making text image")
	}

	TextImageSettingImg, imageDecodeErr := png.Decode(bytes.NewReader(txtImageBytes))
	if imageDecodeErr != nil {
		return nil, imageDecodeErr
	}
	return TextImageSettingImg, nil
}
