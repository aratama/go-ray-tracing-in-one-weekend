package raytracing

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"
)

const samplesPerPixel = 100
const maxDepth = 50

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

func rayColor(ray Ray, world Hittable, depth int, random *rand.Rand) Color {
	rec := HitRecord{}

	if depth <= 0 {
		return vec3(0, 0, 0)
	}

	if world.hit(ray, 0.001, math.Inf(1), &rec) {

		var scattered Ray
		var attenuation Color
		if rec.material.scatter(&ray, &rec, &attenuation, &scattered, random) {
			return hadamard(attenuation, rayColor(scattered, world, depth-1, random))
		}

		target := add(add(rec.p, rec.normal), randomUnitVector(random))
		return mul(0.5, rayColor(Ray{origin: rec.p, direction: sub(target, rec.p)}, world, depth-1, random))
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
	var pixelColor Vec3
	random := rand.New(rand.NewSource(0))
	for s := 0; s < samplesPerPixel; s++ {
		u := (float64(i) + random.Float64()) / (imageWidth - 1)
		v := (float64(j) + random.Float64()) / (imageHeight - 1)
		pixelColor = add(pixelColor, rayColor(cam.getRay(u, v), &world, 50, random))
	}

	ch <- Pixel{x: i, y: j, color: pixelColor}
}

func Render() {

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	start := time.Now()

	ch := make(chan Pixel, imageWidth*imageHeight)

	world := HittableList{hittables: []Hittable{
		&Sphere{center: vec3(0, 0, -1), radius: 0.5, material: &Lambertian{albedo: vec3(0.1, 0.2, 0.5)}},
		&Sphere{center: vec3(0, -100.5, -1), radius: 100, material: &Lambertian{albedo: vec3(0.8, 0.8, 0.0)}},
		&Sphere{center: vec3(1, 0, -1), radius: 0.5, material: &Metal{albedo: vec3(0.8, 0.6, 0.2), fuzz: 0.3}},
		&Sphere{center: vec3(-1, 0, -1), radius: 0.5, material: &Dielectric{refIdx: 1.5}},
		&Sphere{center: vec3(-1, 0, -1), radius: -0.45, material: &Dielectric{refIdx: 1.5}},
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

			// with gumma correction (gumma = 2.0)
			px.color.x = math.Sqrt(px.color.x * scale)
			px.color.y = math.Sqrt(px.color.y * scale)
			px.color.z = math.Sqrt(px.color.z * scale)
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
