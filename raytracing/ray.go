package raytracing

type Ray struct {
	origin    Vec3
	direction Vec3
}

func at(ray *Ray, t float64) Point {
	return add(ray.origin, mul(t, ray.direction))
}

func (ray *Ray) at(t float64) Point {
	return add(ray.origin, mul(t, ray.direction))
}
