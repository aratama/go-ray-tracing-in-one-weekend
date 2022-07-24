package raytracing

import "math/rand"

func random_scene(random *rand.Rand) HittableList {

	ground_material := Lambertian{vec3(0.5, 0.5, 0.5)}

	world := []Hittable{
		&Sphere{center: vec3(0, -1000, 0), radius: 1000, material: &ground_material},
	}

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			choose_mat := random.Float64()
			center := vec3(float64(a)+0.9*random.Float64(), 0.2, float64(b)+0.9*random.Float64())

			if (length(sub(center, vec3(4, 0.2, 0)))) > 0.9 {
				if choose_mat < 0.7 {
					// diffuse
					albedo := hadamard(randomVec3(random), randomVec3(random))
					material := Lambertian{albedo: albedo}
					world = append(world, &Sphere{center: center, radius: 0.2, material: &material})
				} else if choose_mat < 0.85 {
					// metal
					albedo := randomVec3MinMax(0.5, 1.0, random)
					fuzz := randomMinMax(0, 0.5, random)
					material := Metal{albedo, fuzz}
					world = append(world, &Sphere{center, 0.2, &material})
				} else {
					// glass
					material := Dielectric{1.5}
					world = append(world, &Sphere{center, 0.2, &material})
				}
			}
		}
	}

	material1 := Dielectric{1.5}
	world = append(world, &Sphere{vec3(0, 1, 0), 1.0, &material1})
	material2 := Lambertian{vec3(0.4, 0.2, 0.1)}
	world = append(world, &Sphere{vec3(-4, 1, 0), 1.0, &material2})
	material3 := Metal{vec3(0.7, 0.6, 0.5), 0.0}
	world = append(world, &Sphere{vec3(4, 1, 0), 1.0, &material3})

	return HittableList{world}
}
