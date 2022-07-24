package raytracing

import (
	"math"
	"math/rand"
)

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
	*scattered = Ray{origin: rec.p, direction: add(reflected, mul(metal.fuzz, randomInUnitSphere(random)))}
	*attenuation = metal.albedo
	return (dot(scattered.direction, rec.normal) > 0)
}

func refract(uv Vec3, n Vec3, etai_over_etat float64) Vec3 {
	cos_theta := dot(negate(uv), n)
	r_out_parallel := mul(etai_over_etat, add(uv, mul(cos_theta, n)))
	r_out_perp := mul(-math.Sqrt(1.0-r_out_parallel.lengthSquared()), n)
	return add(r_out_parallel, r_out_perp)
}

type Dielectric struct {
	refIdx float64
}

func (dielectric *Dielectric) scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray, random *rand.Rand) bool {
	*attenuation = vec3(1.0, 1.0, 1.0)
	var etaiOverEtat float64
	if rec.frontFace {
		etaiOverEtat = 1.0 / dielectric.refIdx
	} else {
		etaiOverEtat = dielectric.refIdx
	}
	unitDirection := unit(rIn.direction)

	cosTheta := math.Min(dot(negate(unitDirection), rec.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaiOverEtat*sinTheta > 1.0 {
		refracted := reflect(unitDirection, rec.normal)
		*scattered = Ray{origin: rec.p, direction: refracted}
		return true
	}

	refracted := refract(unitDirection, rec.normal, etaiOverEtat)
	*scattered = Ray{origin: rec.p, direction: refracted}
	return true

}
