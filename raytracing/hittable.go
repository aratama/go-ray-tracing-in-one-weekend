package raytracing

import "math"

type HitRecord struct {
	p         Vec3
	normal    Vec3
	t         float64
	frontFace bool
}

func (hitRecord *HitRecord) setFaceNormal(ray *Ray, outwardNormal *Vec3) {
	hitRecord.frontFace = dot(ray.direction, *outwardNormal) < 0
	if hitRecord.frontFace {
		hitRecord.normal = *outwardNormal
	} else {
		hitRecord.normal = negate(*outwardNormal)
	}

}

type Hittable interface {
	hit(r Ray, tMin float64, tMax float64, hitRecord *HitRecord) bool
}

type Sphere struct {
	center Vec3
	radius float64
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
			rec.p = ray.At(rec.t)
			outwardNormal := mul(sub(rec.p, sphere.center), 1/sphere.radius)
			rec.setFaceNormal(&ray, &outwardNormal)
			return true
		}
		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = ray.At(rec.t)
			outwardNormal := mul(sub(rec.p, sphere.center), 1/sphere.radius)
			rec.setFaceNormal(&ray, &outwardNormal)
			return true
		}
	}
	return false

}

type HittableList struct {
	hittables []Hittable
}

func (list *HittableList) hit(ray Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	tempRec := HitRecord{}
	hitAnything := false
	closestSoFar := tMax
	for _, object := range list.hittables {
		if object.hit(ray, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}
	return hitAnything
}
