package raytracing

import (
	"math"
	"math/rand"
)

type Sphere struct {
	center   Vec3
	radius   float64
	material Material
}

func (sphere *Sphere) hit(ray Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	oc := sub(ray.origin, sphere.center)
	a := ray.direction.lengthSquared()
	halfB := dot(oc, ray.direction)
	c := oc.lengthSquared() - sphere.radius*sphere.radius
	discriminant := halfB*halfB - a*c
	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		temp := (-halfB - root) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = ray.at(rec.t)
			outwardNormal := mul(sub(rec.p, sphere.center), 1/sphere.radius)
			rec.setFaceNormal(&ray, &outwardNormal)
			rec.material = sphere.material
			return true
		}
		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = ray.at(rec.t)
			outwardNormal := mul(sub(rec.p, sphere.center), 1/sphere.radius)
			rec.setFaceNormal(&ray, &outwardNormal)
			rec.material = sphere.material
			return true
		}
	}
	return false

}

func randomInUnitSphere(random *rand.Rand) Vec3 {
	for {
		p := randomVec3MinMax(-1, 1, random)
		if p.length() >= 1.0 {
			continue
		}
		return p
	}
}

func randomUnitVector(random *rand.Rand) Vec3 {
	a := randomMinMax(0, 2.0*math.Pi, random)
	z := randomMinMax(-1, 1, random)
	r := math.Sqrt(1.0 - z*z)
	return vec3(r*math.Cos(a), r*math.Sin(a), z)
}
