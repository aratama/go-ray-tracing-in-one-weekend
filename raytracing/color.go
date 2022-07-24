package raytracing

import (
	"image/color"
)

type Color = Vec3

func vecColorToRGBA(v Vec3) color.RGBA {
	r := v.x
	g := v.y
	b := v.z
	a := 1.0
	ir := uint8(255.999 * clamp(r, 0.0, 0.999))
	ig := uint8(255.999 * clamp(g, 0.0, 0.999))
	ib := uint8(255.999 * clamp(b, 0.0, 0.999))
	ia := uint8(255.999 * clamp(a, 0.0, 0.999))
	return color.RGBA{R: ir, G: ig, B: ib, A: ia}
}
