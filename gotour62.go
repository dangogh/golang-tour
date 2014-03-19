package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	c    color.Model
	w, h int
}

func (img Image) ColorModel() color.Model {
	return img.c
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 255, 255}
}

func main() {
	m := Image{color.RGBAModel, 10, 20}
	pic.ShowImage(m)
}
