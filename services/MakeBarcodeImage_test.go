package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeBarcodeImage(t *testing.T) {

	barcodeImg, makeBarcodeImageErr := MakeBarcodeImage("02V1003", 1000, 500)

	assert.Empty(t, makeBarcodeImageErr)
	assert.NotEmpty(t, barcodeImg)

	_, emptyInputErr := MakeBarcodeImage("", 1000, 500)
	assert.NotEmpty(t, emptyInputErr)
}
