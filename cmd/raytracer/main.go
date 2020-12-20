package main

import (
	"fmt"
	"os"

	rt "github.com/andrewzlchen/raytracer/src"
)

const (
	// IMAGE
	aspectRatio = 16.0 / 9.0
	imageWidth  = 400
	imageHeight = int(imageWidth / aspectRatio)

	// CAMERA
	viewportHeight = 2.0
	viewportWidth  = aspectRatio * viewportHeight
	focalLength    = 1.0
)

var (
	origin     = rt.NewVec3(0, 0, 0)
	horizontal = rt.NewVec3(viewportWidth, 0, 0)
	vertical   = rt.NewVec3(0, viewportHeight, 0)

	horizontalDivTwo, _ = horizontal.DivideFloat(2)
	verticalDivTwo, _   = horizontal.DivideFloat(2)
	lowerLeftCorner     = origin.SubtractVector(horizontalDivTwo).SubtractVector(verticalDivTwo).SubtractVector(rt.NewVec3(0, 0, focalLength))
)

// RENDER
func main() {
	// Print the p3 metadata
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)

			direction := lowerLeftCorner.
				AddVector(horizontal.MultiplyFloat(u)).
				AddVector(vertical.MultiplyFloat(v)).
				SubtractVector(origin)

			ray := rt.NewRay(origin, direction)
			currentColor, err := ray.Color()
			if err != nil {
				panic(fmt.Sprintf("could not get color: %s", err))
			}
			rt.WriteColor(os.Stdout, currentColor)
		}
	}
	fmt.Fprint(os.Stderr, "Done!\n")
}
