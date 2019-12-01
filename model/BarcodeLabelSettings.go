package model

type BarcodeLabelSettings struct {
	BarcodeSetting
	CodeSetting, ModelSetting, PriceSetting                        TextImageSetting
	StyleNameSetting, VariantSetting, SizeSetting, MaterialSetting TextImageSetting
}

type BarcodeLabelPositions struct {
	BarcodePos, CodePos, ModelPos, StyleNamePos ImagePosition
	VariantPost, MaterialPos, SizePos, PricePos ImagePosition
}
