package raytracer

import (
	"fmt"
	"io"
	"math"
)

// WriteColor writes color vector values out to an output stream
func WriteColor(w io.Writer, color *Vec3, samplesPerPixel int) error {
	r := color.X
	g := color.Y
	b := color.Z

	// Divide the color by the number of samples and gamma-correct for gamma = 2.0
	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(r * scale)
	g = math.Sqrt(g * scale)
	b = math.Sqrt(b * scale)

	// Write the translated [0,255] value of each color component
	output := fmt.Sprintf("%d %d %d\n", int(256*clamp(r, 0.0, 0.999)), int(256*clamp(g, 0.0, 0.999)), int(256*clamp(b, 0.0, 0.999)))
	_, err := w.Write([]byte(output))
	if err != nil {
		return err
	}
	return nil
}
