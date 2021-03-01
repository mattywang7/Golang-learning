package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

// Define your own Image type, implement the necessary methods, and call pic.ShowImage.
// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
// ColorModel should return color.RGBAModel.
// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

type Image struct {
	Height int
	Width int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Height, m.Width)
}

func (m Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
