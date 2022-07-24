package raytracing

import "image/color"

type Color = Vec3

func Vec3ToColor(r float64, g float64, b float64, a float64) color.RGBA {
	ir := uint8(255.999 * r)
	ig := uint8(255.999 * g)
	ib := uint8(255.999 * b)
	ia := uint8(255.999 * a)
	return color.RGBA{R: ir, G: ig, B: ib, A: ia}
}

func VecToColor(v Vec3) color.RGBA {
	return Vec3ToColor(v.x, v.y, v.z, 1)
}
