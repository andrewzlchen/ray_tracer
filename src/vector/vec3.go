package vector

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

// Add is a helper for adding two Vec3 structs together
func (v *Vec3) Add(other *Vec3) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

// Multiply is a helper for multiplying two Vec3 structs together
func (v *Vec3) Multiply(m float64) {
	v.X *= m
	v.Y *= m
	v.Z *= m
}

// Divide is a helper for multiplying two Vec3 structs together
func (v *Vec3) Divide(d float64) {
	v.Multiply(1 / d)
}

// Length returns the length of the vector
func (v *Vec3) Length() float64 {
	return v.X + v.Y + v.Z
}

// Dot returns the dot product of two vec3 structs
func Dot(v1, v2 *Vec3) float64 {
	return v1.X*v2.X +
		v1.Y*v2.Y +
		v1.Z*v2.Z
}
