package raytracer

import "errors"

// Camera is a representation of the virtual camera system
type Camera struct {
	origin          *Vec3
	horizontal      *Vec3
	vertical        *Vec3
	lowerLeftCorner *Vec3
}

// NewCamera returns a new camera struct
func NewCamera(origin *Vec3) (*Camera, error) {
	c := &Camera{origin: origin}
	horizontal := NewVec3(c.ViewportWidth(), 0, 0)
	vertical := NewVec3(0, c.ViewportHeight(), 0)
	halfHorizontal, err := horizontal.DivideFloat(2)
	if err != nil {
		return nil, errors.New("could not compute half of horizontal")
	}
	halfVertical, err := vertical.DivideFloat(2)
	if err != nil {
		return nil, errors.New("could not compute half of vertical")
	}

	lowerLeftCorner := origin.
		SubtractVector(halfHorizontal).
		SubtractVector(halfVertical).
		SubtractVector(NewVec3(0, 0, c.FocalLength()))

	c.horizontal = horizontal
	c.vertical = vertical
	c.lowerLeftCorner = lowerLeftCorner

	return c, nil
}

// AspectRatio returns the current aspect ratio of the camera
func (c *Camera) AspectRatio() float64 {
	return 16.0 / 9.0
}

// ViewportHeight returns the viewport height of the camera
func (c *Camera) ViewportHeight() float64 {
	return 2.0
}

// ViewportWidth returns the viewport width of the camera
func (c *Camera) ViewportWidth() float64 {
	return c.AspectRatio() * c.ViewportHeight()
}

// FocalLength returns the focal length of the camera
func (c *Camera) FocalLength() float64 {
	return 1.0
}

// GetRay returns the ray that should be rendered on the (u,v) point on a flat canvas
func (c *Camera) GetRay(u, v float64) *Ray {
	direction := c.lowerLeftCorner.
		AddVector(c.horizontal.MultiplyFloat(u)).
		AddVector(c.vertical.MultiplyFloat(v)).
		SubtractVector(c.origin)

	return NewRay(c.origin, direction)
}
