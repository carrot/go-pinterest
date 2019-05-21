package models

// 60x60 images
type Images struct {
	Size_60x60 Image `json:"60x60"`
}

// Image with width and height
type Image struct {
	Url    string `json:"url"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}
