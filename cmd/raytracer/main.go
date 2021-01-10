package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	rt "github.com/andrewzlchen/raytracer/src"
)

const (
	// IMAGE
	aspectRatio     = 16.0 / 9.0
	imageWidth      = 400
	imageHeight     = int(imageWidth / aspectRatio)
	samplesPerPixel = 100
	maxDepth        = 50
)

// RENDER
func main() {
	// Add elements to the world
	world := &rt.HittableList{}
	world.Add(&rt.Sphere{Center: rt.NewVec3(0, 0, -1), Radius: 0.5})
	world.Add(&rt.Sphere{Center: rt.NewVec3(0, -100.5, -1), Radius: 100})

	// Print the p3 metadata
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	// Set up buffered stdout
	bufferedStdout := bufio.NewWriter(os.Stdout)
	defer bufferedStdout.Flush()

	// Set up camera
	camera, err := rt.NewCamera(rt.NewVec3(0, 0, 0))
	if err != nil {
		panic("could not set up the camera")
	}

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			pixelColor := rt.NewVec3(0, 0, 0)
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)

				ray := camera.GetRay(u, v)
				currentColor, err := ray.Color(world, maxDepth)
				if err != nil {
					panic(fmt.Sprintf("could not get color: %s", err))
				}
				pixelColor = pixelColor.AddVector(currentColor)
			}
			rt.WriteColor(bufferedStdout, pixelColor, samplesPerPixel)
		}
	}
	fmt.Fprint(os.Stderr, "Done!\n")
}
