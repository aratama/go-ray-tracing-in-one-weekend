package raytracing

import (
	"math"
	"math/rand"
)

type Camera struct {
	origin          Vec3
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
	u               Vec3
	v               Vec3
	w               Vec3
	lensRadius      float64
}

func camera(
	lookfrom Vec3,
	lookat Vec3,
	vup Vec3,
	vfov float64,
	aspectRatio float64,
	aperture float64,
	focusDist float64,
) Camera {
	theta := degreesToRadians(vfov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := unit(sub(lookfrom, lookat))
	u := unit(cross(vup, w))
	v := cross(w, u)

	origin := lookfrom
	horizontal := mul(focusDist*viewportWidth, u)
	vertical := mul(focusDist*viewportHeight, v)
	lowerLeftCorner := sub(sub((sub(origin, mul(0.5, horizontal))), mul(0.5, vertical)), mul(focusDist, w))
	lensRadius := aperture / 2.0

	return Camera{
		origin:          origin,
		lowerLeftCorner: lowerLeftCorner,
		horizontal:      horizontal,
		vertical:        vertical,
		u:               u,
		v:               v,
		w:               w,
		lensRadius:      lensRadius,
	}
}

func (camera *Camera) getRay(u float64, v float64, random *rand.Rand) Ray {
	rd := mul(camera.lensRadius, randomInUnitDisk(random))
	offset := add(mul(rd.x, camera.u), mul(rd.y, camera.v))
	return Ray{
		origin:    camera.origin,
		direction: sub(sub(add(add(camera.lowerLeftCorner, mul(u, camera.horizontal)), mul(v, camera.vertical)), camera.origin), offset),
	}
}
