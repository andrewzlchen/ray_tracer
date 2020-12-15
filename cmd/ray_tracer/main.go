package main

import (
	"fmt"
	"os"

	"github.com/andrewzlchen/ray_tracer/src/color"
	"github.com/andrewzlchen/ray_tracer/src/vector"
)

// size of the output image
const (
	imageWidth  = 256
	imageHeight = 256
)

func main() {
	fmt.Fprint(os.Stderr, "Starting to write values\n")

	// print the p3 metadata
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for i := imageHeight - 1; i >= 0; i-- {
		for j := 0; j < imageWidth; j++ {
			r := float64(j) / (imageWidth - 1)
			g := float64(i) / (imageHeight - 1)
			b := 0.25
			currentColor := vector.NewVec3(r, g, b)

			color.WriteColor(os.Stdout, currentColor)
		}
	}
	fmt.Fprint(os.Stderr, "Done!\n")
}
