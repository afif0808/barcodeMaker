package model

type TextImageSetting struct {
	FontSize   float64
	FontStyle  string
	FontWeight int
	FontSlant  int
}

type TextImageSettings map[string]TextImageSetting
