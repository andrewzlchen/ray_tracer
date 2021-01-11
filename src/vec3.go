package raytracer

import (
	"errors"
	"math"
	"math/rand"
)

// Vec3 is a representation of a 3 dimensional vector
type Vec3 struct {
	X, Y, Z float64
}

// NewVec3 returns a vec3 struct from a tuple of values representing the vector's dimensions
func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{
		x, y, z,
	}
}

// Dot returns the dot product of two vec3 structs
func (v *Vec3) Dot(other *Vec3) float64 {
	return v.X*other.X +
		v.Y*other.Y +
		v.Z*other.Z
}

// AddVector returns a new Vec3 that is a returned by adding two Vec3 structs together
func (v *Vec3) AddVector(other *Vec3) *Vec3 {
	return &Vec3{
		X: other.X + v.X,
		Y: other.Y + v.Y,
		Z: other.Z + v.Z,
	}
}

// SubtractVector returns a new Vec3 that is a returned by dividing a Vec3 by a float
func (v *Vec3) SubtractVector(subtractor *Vec3) *Vec3 {
	return &Vec3{
		X: v.X - subtractor.X,
		Y: v.Y - subtractor.Y,
		Z: v.Z - subtractor.Z,
	}
}

// MultiplyVector returns a new vector that is obtained by multiplying the Vec3 that is by another vector
func (v *Vec3) MultiplyVector(multiplier *Vec3) *Vec3 {
	return &Vec3{
		X: v.X * multiplier.X,
		Y: v.Y * multiplier.Y,
		Z: v.Z * multiplier.Z,
	}
}

// MultiplyFloat returns a new Vec3 that is a returned by multiplying a Vec3 by a float
func (v *Vec3) MultiplyFloat(multiplier float64) *Vec3 {
	return &Vec3{
		X: v.X * multiplier,
		Y: v.Y * multiplier,
		Z: v.Z * multiplier,
	}
}

// DivideVector returns a new Vec3 that is a returned by dividing two Vec3 structs together
func (v *Vec3) DivideVector(other *Vec3) (*Vec3, error) {
	if other.X == 0 || other.Y == 0 || other.Z == 0 {
		return nil, errors.New("cannot divide by 0")
	}
	return &Vec3{
		X: v.X / other.X,
		Y: v.Y / other.Y,
		Z: v.Z / other.Z,
	}, nil
}

// DivideFloat returns a new Vec3 that is a returned by dividing a Vec3 by a float
func (v *Vec3) DivideFloat(divisor float64) (*Vec3, error) {
	if divisor == 0 {
		return nil, errors.New("cannot divide by 0")
	}
	return &Vec3{
		X: v.X / divisor,
		Y: v.Y / divisor,
		Z: v.Z / divisor,
	}, nil
}

// Length returns the length of the vector
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the squared length of the vector
func (v *Vec3) LengthSquared() float64 {
	return v.Dot(v)
}

// Unit returns a unit vector of the caller
func (v *Vec3) Unit() (*Vec3, error) {
	v, err := v.DivideFloat(v.Length())
	if err != nil {
		return nil, errors.New("cannot get the unit vector of a zero vector")
	}
	return v, nil
}

// Random returns a vec3 with random x, y and z values
func Random() *Vec3 {
	return NewVec3(rand.Float64(), rand.Float64(), rand.Float64())
}

// RandomBound returns a vec3 with a random x, y, and z values between min and max
func RandomBound(min, max float64) *Vec3 {
	return NewVec3(randomFloat(min, max), randomFloat(min, max), randomFloat(min, max))
}

// RandomUnitInUnitSphere returns a vector that touches the unit sphere
func RandomUnitInUnitSphere() *Vec3 {
	for {
		p := RandomBound(-1.0, 1.0)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

// RandomUnitVector returns the unit vector of a vector that touches the the unit sphere
func RandomUnitVector() (*Vec3, error) {
	unit, err := RandomUnitInUnitSphere().Unit()
	if err != nil {
		return nil, err
	}
	return unit, nil
}

// RandomInHemisphere returns a random vector within the same hemisphere of the normal
func RandomInHemisphere(normal *Vec3) *Vec3 {
	inUnitSphere := RandomUnitInUnitSphere()
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.MultiplyFloat(-1.0)
}
