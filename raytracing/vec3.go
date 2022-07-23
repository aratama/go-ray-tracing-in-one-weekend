package raytracing

import "math"

type Vec3 struct {
	x float64
	y float64
	z float64
}

type Point = Vec3

type Color = Vec3

func vec3(x float64, y float64, z float64) Vec3 {
	return Vec3{x: x, y: y, z: z}
}

func addVec3(v Vec3, t Vec3) Vec3 {
	return Vec3{x: v.x + t.x, y: v.y + t.y, z: v.z + t.z}
}

func subVec3(v Vec3, t Vec3) Vec3 {
	return Vec3{x: v.x - t.x, y: v.y - t.y, z: v.z - t.z}
}

func (v *Vec3) sub(t Vec3) Vec3 {
	return Vec3{x: v.x - t.x, y: v.y - t.y, z: v.z - t.z}
}

func negateVec3(v Vec3) Vec3 {
	return Vec3{x: -v.x, y: -v.y, z: -v.z}
}

func mulVec3(v Vec3, t float64) Vec3 {
	return Vec3{x: v.x * t, y: v.y * t, z: v.z * t}
}

func (v *Vec3) mul(t float64) Vec3 {
	return Vec3{x: v.x * t, y: v.y * t, z: v.z * t}
}

func lenVec3(v Vec3) float64 {
	return math.Sqrt(lenSquaredVec3(v))
}

func lenSquaredVec3(v Vec3) float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func dot(v Vec3, t Vec3) float64 {
	return v.x*t.x + v.y*t.y + v.z*t.z
}

func cross(u *Vec3, v *Vec3) Vec3 {
	return Vec3{
		x: u.y*v.z - u.z*v.y,
		y: u.z*v.x - u.x*v.z,
		z: u.x*v.y - u.y*v.x,
	}
}

func unit(v Vec3) Vec3 {
	return mulVec3(v, 1/lenVec3(v))
}
