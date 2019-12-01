package services

import (
	"barcodeMaker/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTextImage(t *testing.T) {
	txtImage := model.TextImageSetting{30, "Lato", 0, 0}
	actualImg, actualErr := MakeTextImage("Test", txtImage)
	assert.NotEmpty(t, actualImg)
	assert.Empty(t, actualErr)

	txtImageEmpty := model.TextImageSetting{}
	actualImg, actualErr = MakeTextImage("", txtImageEmpty)
	assert.Empty(t, actualImg)
	assert.NotEmpty(t, actualErr)

}
