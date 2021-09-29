package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)

// your implementation of the image.Image interface
type Image struct{}
 
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 360, 360)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{ uint8( (x + y) / 2), uint8( x * y), uint8( x ^ y), 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
