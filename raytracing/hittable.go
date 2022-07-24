package raytracing

type Hittable interface {
	hit(r Ray, tMin float64, tMax float64, hitRecord *HitRecord) bool
}
