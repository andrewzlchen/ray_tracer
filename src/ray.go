package raytracer

import (
	"fmt"
	"math"
)

// Ray is a struc that contains a origin and a direction and can be described the formula
// P(t) = A + tb where
// - P represents the position of a point in 3D space
// - t is a real number that scales the point along the direction vector
// - b is a vector that represents the direction that the ray is facing
type Ray struct {
	// Origin is the location of a point in 3 dimensional space, denoted by a Vec3
	origin *Vec3
	// Direction is a Vec3 that represents the direction of the ray
	direction *Vec3
}

// NewRay constructs a new ray from an origin and direction vector
func NewRay(origin, direction *Vec3) *Ray {
	return &Ray{
		origin:    origin,
		direction: direction,
	}
}

// Origin returns a vector representing the origin of the ray
func (r *Ray) Origin() *Vec3 {
	return r.origin
}

// Direction returns a vector representing the direction of the ray
func (r *Ray) Direction() *Vec3 {
	return r.direction
}

// At returns the point on the ray given the coefficient, t
func (r *Ray) At(t float64) *Vec3 {
	return r.Origin().
		AddVector(
			r.Direction().
				MultiplyFloat(t),
		)
}

// Color computes the color of the ray.
func (r *Ray) Color(world Hittable) (*Vec3, error) {
	hitRecord, didHit, err := world.Hit(r, 0, math.Inf(1))
	if err != nil {
		return nil, fmt.Errorf("could not compute collision: %s", err)
	}
	if didHit {
		return hitRecord.Normal.AddVector(NewVec3(1, 1, 1)).MultiplyFloat(0.5), nil
	}
	// there is no intersection
	return r.linearBlueGradient()
}

// linearBlueGradient blends color to be a linear gradient on the Y direction.
// We first get the unit vector, of which, the values will be -1 < x < 1.
// We then add 1 and divide by 2 to scale the values to be from 0 to 1.
// 0 = white
// 1 = blue
func (r *Ray) linearBlueGradient() (*Vec3, error) {
	unitDirection, err := r.Direction().Unit()
	if err != nil {
		return nil, fmt.Errorf("cannot get unit vector of vector: %s", err)
	}

	// scale unit vector values to be 0 < x < 1 to determine what color from white to blue to choose
	blueness := 0.5 * (unitDirection.Y + 1.0)

	blue := NewVec3(1.0, 1.0, 1.0)
	shadeOfBlue := blue.MultiplyFloat(1.0 - blueness)

	// This is similar to a seed in a RNG to determine the base level of blueness
	baseBlue := NewVec3(0.5, 0.7, 1.0).MultiplyFloat(blueness)

	return shadeOfBlue.AddVector(baseBlue), nil
}

// hitsSphere determines whether or not the ray, will at some point, given P(t) = A +tb, where P is some point on the ray,
// with intersect with a sphere at point 't'
//
// The equation for a sphere centered at (x, y, z) with radius r is
// ``x^2 + y^2 + z^2 = r^2``
//
// If a given point (a,b,c) is on the surface of the sphere, then a^2 + b^2 + c^2 = r^2
// If a given point (a,b,c) is inside of the sphere, then a^2 + b^2 + c^2 < r^2
// If a given point (a,b,c) is outside of the sphere, then a^2 + b^2 + c^2 > r^2
//
// If we need to solve for a point t such that P(t) - C, where P is a point along the ray and C is the center of a sphere, the
// equation is (P_x,y,z-C_x,y,z) dot (P_x,y,z-C_x,y,z) = r^2
//
// If there are 0 solutions where (P(t)-C) dot (P(t)-C) = r^2, then this ray does not intersect with the sphere
// If there is 1 solution where (P(t)-C) dot (P(t)-C) = r^2, then this ray only intersects with the sphere in one spot, and this ray is tangent to the sphere's surface
// If there are 2 solution where (P(t)-C) dot (P(t)-C) = r^2, then this ray goes through the sphere, and there is a point on the front side where this ray intersects and one more in the back
//
// The final equation to solve for t given x,y,z is:
// t^2(b ⋅ b) + 2tb ⋅ (A−C) + (A−C) ⋅ (A−C) − r^2 = 0
//
// We can use the quadratic formula to solve for t where a = (b ⋅ b), b = 2(b ⋅ (A-C))), and c = ((A-C)) ⋅ (A-C) - r^2)
// but if we utilize the fact that if b = 2h, the formula becomes:
//
// (-h +- sqrt(h^2 - ac)) / a
func (r *Ray) hitsSphere(center *Vec3, radius float64) float64 {
	oc := r.Origin().SubtractVector(center)
	a := r.Direction().LengthSquared()
	halfB := oc.Dot(r.Direction())
	c := oc.LengthSquared() - radius*radius

	// discriminant of the quadratic formula
	//     discriminant > 0 -> 2 solutions
	//     discriminant = 0 -> 1 solution
	//     discriminant < 0 -> 0 solutions
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1.0
	}
	// quadratic formula: solve for t given t = (-b +- sqrt( b^2 - 4ac )) / 2a
	return (-halfB - math.Sqrt(discriminant)) / a
}
