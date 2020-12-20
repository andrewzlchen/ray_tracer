package raytracer

import (
	"fmt"
	"io"
)

// WriteColor writes color vector values out to an output stream
func WriteColor(w io.Writer, v *Vec3) error {
	output := fmt.Sprintf("%d %d %d\n", int(v.X*255.999), int(v.Y*255.999), int(v.Z*255.999))
	_, err := w.Write([]byte(output))
	if err != nil {
		return err
	}
	return nil
}
