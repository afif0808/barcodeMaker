package services

import (
	"barcodeMaker/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeBarcodeLabelImage2(t *testing.T) {
	_, makeBarcodeLabelImage2Err := MakeBarcodeLabelImage2(
		model.MainImageSetting{Width: 3150, Height: 1260},
		// model.BarcodeSetting{Width: 2397, Height: 567, X: 20, Y: int(float64(0.3) * float64(1260))},
		model.BarcodeSetting{Width: 2397, Height: 567},
		model.TextImageSettings{},
		MakeBarcodeImage)(model.BarcodeLabel{Code: "V012"})
	assert.Empty(t, makeBarcodeLabelImage2Err)
}
