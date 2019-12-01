package controllers

import (
	"barcodeMaker/mocks"
	"image"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestPrintBarcode(t *testing.T) {
	printBarcodeLabelMock := mocks.PrintBarcodeLabelImageMock
	makeBarcodeLabelMock := mocks.MakeBarcodeLabelImageMock
	handlePostDataMock := mocks.HandlePostDataMock

	expectedStatusCode := http.StatusOK
	var img image.Image
	printBarcodeHandler := PrintBarcode(
		handlePostDataMock(map[string]interface{}{
			"products":   []map[string]interface{}{},
			"quantities": []interface{}{}},
			nil),
		makeBarcodeLabelMock(img, nil),
		printBarcodeLabelMock(nil),
	)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/printBarcode", nil)

	router := mux.NewRouter()

	router.Handle("/printBarcode", printBarcodeHandler)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, expectedStatusCode, recorder.Code)
}
