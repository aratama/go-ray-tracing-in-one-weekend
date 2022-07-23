package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const imageHeight = 256
const imageWidth = 256

func main() {

	img := image.NewRGBA(image.Rect(0, 0, 128, 128))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25
			ir := uint8(255.999 * r)
			ig := uint8(255.999 * g)
			ib := uint8(255.999 * b)
			img.SetRGBA(i, j, color.RGBA{R: ir, G: ig, B: ib, A: 255})
		}
	}

	out, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}

}
