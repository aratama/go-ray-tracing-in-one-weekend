package raytracing

type Ray struct {
	origin    Vec3
	direction Vec3
}

func At(ray *Ray, t float64) Point {
	return add(ray.origin, mul(ray.direction, t))
}

func (ray *Ray) At(t float64) Point {
	return add(ray.origin, mul(ray.direction, t))
}
