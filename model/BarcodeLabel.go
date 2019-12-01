package model

import "fmt"

type BarcodeLabel struct {
	Model     string
	Code      string
	Size      string
	Price     string
	Variant   string // or color
	Material  string
	StyleName string
}

func MapToBarcodeLabelArray(products []map[string]interface{}) ([]BarcodeLabel, error) {
	barcodeLabels := []BarcodeLabel{}
	for _, product := range products {
		if product["code"] == nil || product["price"] == nil {
			return nil, fmt.Errorf("Error : barcode label input doesn't meet the requirement")
		}
		barcodeLabel := BarcodeLabel{
			fmt.Sprint(product["styleName"]),
			fmt.Sprint(product["model"]),
			fmt.Sprint(product["code"]),
			fmt.Sprint(product["color"]),
			fmt.Sprint(product["material"]),
			fmt.Sprint(product["size"]),
			fmt.Sprint(product["price"]),
		}

		barcodeLabels = append(barcodeLabels, barcodeLabel)
	}

	return barcodeLabels, nil
}
