package models

type Images struct {
	Size_60x60 Image `json:"60x60"`
}

type Image struct {
	Url    string `json:"url"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}
