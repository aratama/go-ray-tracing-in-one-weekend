package raytracing

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	x float64
	y float64
	z float64
}

type Point = Vec3

func vec3(x float64, y float64, z float64) Vec3 {
	return Vec3{x: x, y: y, z: z}
}

func add(v Vec3, t Vec3) Vec3 {
	return Vec3{x: v.x + t.x, y: v.y + t.y, z: v.z + t.z}
}

func sub(v Vec3, t Vec3) Vec3 {
	return Vec3{x: v.x - t.x, y: v.y - t.y, z: v.z - t.z}
}

func (v *Vec3) sub(t Vec3) Vec3 {
	return Vec3{x: v.x - t.x, y: v.y - t.y, z: v.z - t.z}
}

func negate(v Vec3) Vec3 {
	return Vec3{x: -v.x, y: -v.y, z: -v.z}
}

func mul(t float64, v Vec3) Vec3 {
	return Vec3{x: v.x * t, y: v.y * t, z: v.z * t}
}

func (v *Vec3) mul(t float64) Vec3 {
	return Vec3{x: v.x * t, y: v.y * t, z: v.z * t}
}

func length(v Vec3) float64 {
	return math.Sqrt(lengthSquared(v))
}

func (v *Vec3) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func lengthSquared(v Vec3) float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v *Vec3) lengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func dot(v Vec3, t Vec3) float64 {
	return v.x*t.x + v.y*t.y + v.z*t.z
}

func cross(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.y*v.z - u.z*v.y,
		y: u.z*v.x - u.x*v.z,
		z: u.x*v.y - u.y*v.x,
	}
}

func hadamard(u Vec3, v Vec3) Vec3 {
	return Vec3{
		x: u.x * v.x,
		y: u.y * v.y,
		z: u.z * v.z,
	}
}

func unit(v Vec3) Vec3 {
	return mul(1/length(v), v)
}

func lerp(u Vec3, v Vec3, t float64) Vec3 {
	return add(mul(1-t, u), mul(t, v))
}

func randomVec3(random *rand.Rand) Vec3 {
	return vec3(random.Float64(), random.Float64(), random.Float64())
}

func randomVec3MinMax(min float64, max float64, random *rand.Rand) Vec3 {
	return vec3(
		randomMinMax(min, max, random),
		randomMinMax(min, max, random),
		randomMinMax(min, max, random),
	)
}

func reflect(v Vec3, n Vec3) Vec3 {
	return sub(v, mul(2.0*dot(v, n), n))
}

func randomInUnitDisk(random *rand.Rand) Vec3 {
	for {
		p := vec3(randomMinMax(-1, 1, random), randomMinMax(-1, 1, random), 0)
		if p.lengthSquared() >= 1 {
			continue
		}
		return p
	}
}
