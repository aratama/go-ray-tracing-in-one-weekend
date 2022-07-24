package raytracing

import "math/rand"

type Material interface {
	scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray, random *rand.Rand) bool
}

type Lambertian struct {
	albedo Color
}

func (lambertian *Lambertian) scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray, random *rand.Rand) bool {
	scatterDirection := add(rec.normal, randomUnitVector(random))
	*scattered = Ray{origin: rec.p, direction: scatterDirection}
	*attenuation = lambertian.albedo
	return true
}

type Metal struct {
	albedo Color
	fuzz   float64
}

func (metal *Metal) scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray, random *rand.Rand) bool {
	reflected := reflect(unit(rIn.direction), rec.normal)
	*scattered = Ray{origin: rec.p, direction: add(reflected, mul(randomInUnitSphere(random), metal.fuzz))}
	*attenuation = metal.albedo
	return (dot(scattered.direction, rec.normal) > 0)
}
