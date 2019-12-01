package controllers

import (
	"barcodeMaker/model"
	"barcodeMaker/services"
	"image"
	"net/http"
)

func PrintBarcode(
	handlePostData model.HandlePostData,
	makeBarcodeLabelImage model.MakeBarcodeLabelImage,
	printBarcodeLabelImage model.PrintBarcodeLabelImage,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		postData, handlePostRequestErr := handlePostData(*r)
		if handlePostRequestErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		products, isProductsTypeValid := postData["products"].([]map[string]interface{})
		if isProductsTypeValid == false {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		quantitiesInterface, isQuantitiesInterfaceArray := postData["quantities"].([]interface{})
		if isQuantitiesInterfaceArray == false {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		quantitiesInt64, interfaceArrayToInt64ArrayErr := services.InterfaceArrayToInt64Array(quantitiesInterface)
		if interfaceArrayToInt64ArrayErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		barcodeLabels, mapToBarcodeLabelArrayErr := model.MapToBarcodeLabelArray(products)
		if mapToBarcodeLabelArrayErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		barcodeLabelImages := []image.Image{}
		for _, barcodeLabel := range barcodeLabels {
			barcodeLabelImage, makeBarcodeLabelImageErr := makeBarcodeLabelImage(barcodeLabel)
			if makeBarcodeLabelImageErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			barcodeLabelImages = append(barcodeLabelImages, barcodeLabelImage)
		}
		for index, barcodeLabelImage := range barcodeLabelImages {
			printBarcodeLabelImageErr := printBarcodeLabelImage(barcodeLabelImage, quantitiesInt64[index])
			if printBarcodeLabelImageErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	})
}
