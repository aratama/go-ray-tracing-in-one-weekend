package raytracing

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
