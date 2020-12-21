package raytracer

import (
	"fmt"
	"io"
)

// WriteColor writes color vector values out to an output stream
func WriteColor(w io.Writer, color *Vec3, samplesPerPixel int) error {
	r := color.X
	g := color.Y
	b := color.Z

	// Divide the colors by the number of samples
	scale := 1.0 / float64(samplesPerPixel)
	r *= scale
	g *= scale
	b *= scale

	// Write the translated [0,255] value of each color component
	output := fmt.Sprintf("%d %d %d\n", int(256*clamp(r, 0.0, 0.999)), int(256*clamp(g, 0.0, 0.999)), int(256*clamp(b, 0.0, 0.999)))
	_, err := w.Write([]byte(output))
	if err != nil {
		return err
	}
	return nil
}
