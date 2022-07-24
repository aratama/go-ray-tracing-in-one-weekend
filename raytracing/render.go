package raytracing

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"sync"
	"time"
)

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

func lerp(u Vec3, v Vec3, t float64) Vec3 {
	return add(mul(u, 1-t), mul(v, t))
}

func rayColor(ray Ray) Color {
	t := hitSphere(vec3(0, 0, -1), 0.5, ray)
	if t > 0.0 {
		n := unit(sub(ray.At(t), vec3(0, 0, -1)))
		return mul(vec3(n.x+1, n.y+1, n.z+1), 0.5)
	} else {
		unitDirection := unit(ray.direction)
		t := 0.5 * (unitDirection.y + 1.0)
		return lerp(vec3(1, 1, 1), vec3(0.5, 0.7, 1.0), t)
	}
}

func hitSphere(center Point, radius float64, ray Ray) float64 {
	oc := sub(ray.origin, center)
	a := ray.direction.lengthSquared()
	halfB := dot(oc, ray.direction)
	c := oc.lengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1
	} else {
		return (-halfB - math.Sqrt(discriminant)) / a
	}
}

type Pixel struct {
	x     int
	y     int
	color Color
}

func pathTrace(i int, j int, ch chan Pixel, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	u := float64(i) / (imageWidth - 1)
	v := float64(j) / (imageHeight - 1)
	direction := sub(add(add(lowerLeftCorner, horizontal.mul(u)), vertical.mul(v)), origin)
	r := Ray{origin: origin, direction: direction}
	ch <- Pixel{x: i, y: j, color: rayColor(r)}
}

func Render() {

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	start := time.Now()

	ch := make(chan Pixel, imageWidth*imageHeight)

	var waitGroup sync.WaitGroup

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			waitGroup.Add(1)
			go pathTrace(i, j, ch, &waitGroup)
		}
	}

	waitGroup.Wait()

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			px := <-ch
			img.SetRGBA(px.x, imageHeight-1-px.y, VecToColor(px.color))
		}
	}

	fmt.Printf("rendering time: %s\n", time.Now().
		Sub(start).String())

	encodeStart := time.Now()

	out, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}

	fmt.Printf("encodeing time: %s\n", time.Now().Sub(encodeStart).String())

}
