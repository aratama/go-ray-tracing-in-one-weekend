package raytracing

type HitRecord struct {
	p         Vec3
	normal    Vec3
	t         float64
	frontFace bool
	material  Material
}

func (hitRecord *HitRecord) setFaceNormal(ray *Ray, outwardNormal *Vec3) {
	hitRecord.frontFace = dot(ray.direction, *outwardNormal) < 0
	if hitRecord.frontFace {
		hitRecord.normal = *outwardNormal
	} else {
		hitRecord.normal = negate(*outwardNormal)
	}

}
