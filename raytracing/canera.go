package raytracing

type Camera struct {
	origin          Vec3
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
}

func camera() Camera {
	const aspectRatio = 16.0 / 9.0
	const imageWidth = 384
	const imageHeight = imageWidth / aspectRatio

	const viewportHeight = 2.0
	const viewportWidth = aspectRatio * viewportHeight
	const focalLength = 1

	var origin = vec3(0, 0, 0)
	var horizontal = vec3(viewportWidth, 0, 0)
	var vertical = vec3(0, viewportHeight, 0)
	var lowerLeftCorner = sub(sub((origin.sub(horizontal.mul(0.5))), vertical.mul(0.5)), vec3(0, 0, focalLength))

	return Camera{origin: origin, lowerLeftCorner: lowerLeftCorner, horizontal: horizontal, vertical: vertical}
}

func (camera *Camera) getRay(u float64, v float64) Ray {
	return Ray{
		origin:    camera.origin,
		direction: sub(add(add(lowerLeftCorner, horizontal.mul(u)), vertical.mul(v)), origin),
	}
}
