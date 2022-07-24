package raytracing

type Material interface {
	scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool
}

type Lambertian struct {
	albedo Color
}

func (lambertian *Lambertian) scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool {
	scatterDirection := add(rec.normal, randomUnitVector())
	*scattered = Ray{origin: rec.p, direction: scatterDirection}
	*attenuation = lambertian.albedo
	return true
}

type Metal struct {
	albedo Color
}

func (metal *Metal) scatter(rIn *Ray, rec *HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := reflect(unit(rIn.direction), rec.normal)
	*scattered = Ray{origin: rec.p, direction: reflected}
	*attenuation = metal.albedo
	return (dot(scattered.direction, rec.normal) > 0)
}
