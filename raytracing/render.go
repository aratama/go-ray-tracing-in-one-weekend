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

const samplesPerPixel = 1

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

func rayColor(ray Ray, world Hittable) Color {
	rec := HitRecord{}
	if world.hit(ray, 0, math.Inf(1), &rec) {
		return mul(add(rec.normal, vec3(1, 1, 1)), 0.5)
	}
	unitDirection := unit(ray.direction)
	t := 0.5 * (unitDirection.y + 1.0)
	return lerp(vec3(1, 1, 1), vec3(0.5, 0.7, 1.0), t)
}

type Pixel struct {
	x     int
	y     int
	color Color
}

func pathTrace(world HittableList, i int, j int, ch chan Pixel, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	cam := camera()
	u := float64(i) / (imageWidth - 1)
	v := float64(j) / (imageHeight - 1)
	ch <- Pixel{x: i, y: j, color: rayColor(cam.getRay(u, v), &world)}
}

func Render() {

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	start := time.Now()

	ch := make(chan Pixel, imageWidth*imageHeight)

	world := HittableList{hittables: []Hittable{
		&Sphere{center: vec3(0, 0, -1), radius: 0.5},
		&Sphere{center: vec3(0, -100.5, -1), radius: 100},
	}}

	var waitGroup sync.WaitGroup

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			waitGroup.Add(1)
			go pathTrace(world, i, j, ch, &waitGroup)
		}
	}

	waitGroup.Wait()

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			px := <-ch

			scale := 1.0 / float64(samplesPerPixel)
			px.color.x *= scale
			px.color.y *= scale
			px.color.z *= scale
			img.SetRGBA(px.x, imageHeight-1-px.y, vecColorToRGBA(px.color))
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
