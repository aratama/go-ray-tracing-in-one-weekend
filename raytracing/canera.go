package raytracing

import (
	"math"
)

type Camera struct {
	origin          Vec3
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
}

func camera(lookfrom Vec3, lookat Vec3, vup Vec3, vfov float64, aspectRatio float64) Camera {
	theta := degreesToRadians(vfov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := unit(sub(lookfrom, lookat))
	u := unit(cross(vup, w))
	v := cross(w, u)

	origin := lookfrom
	horizontal := mul(viewportWidth, u)

	vertical := mul(viewportHeight, v)
	lowerLeftCorner := sub(sub((sub(origin, mul(0.5, horizontal))), mul(0.5, vertical)), w)

	return Camera{origin: origin, lowerLeftCorner: lowerLeftCorner, horizontal: horizontal, vertical: vertical}
}

func (camera *Camera) getRay(u float64, v float64) Ray {
	return Ray{
		origin:    camera.origin,
		direction: sub(add(add(camera.lowerLeftCorner, mul(u, camera.horizontal)), mul(v, camera.vertical)), camera.origin),
	}
}
