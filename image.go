package main

import (
	"image"
	"image/draw"
)

// transform an image into an RGBA object
func toRGBA(i image.Image) *image.RGBA {
	bounds := i.Bounds()
	ni := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(ni, ni.Bounds(), i, bounds.Min, draw.Src)
	return ni
}
