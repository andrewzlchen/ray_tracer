package raytracer

import (
	"errors"
	"math"
)

// Sphere is a struct that represents a sphere in 3d space
type Sphere struct {
	Center *Vec3
	Radius float64
}

// Hit returns whether the passed-in ray hits the current sphere
func (s *Sphere) Hit(ray *Ray, tMin, tMax float64) (*HitRecord, bool, error) {
	// Use quadratic equation to solve for t (a point on ray r such that it intersects with the sphere)
	// Using vector math, it works out that we need to solve for:
	//
	// (P(t)-C) dot (P(t) - C) = r^2 : subbing P(t) for its definition P(t) = A + tb, A = origin vector, b = direction vector
	// (A + tb - C) dot (A + tb - C) = r^2
	// ((A-C) + tb) dot ((A-C) + tb) = r^2
	// ((A-C) dot (A-C)) + (2tb dot (A-C)) + t^2(b dot b) - r^2 = 0 : rearranging the formula to get a, b, and c
	// t^2(b dot b) + (2tb dot (A-C)) + ((A-C) dot (A-C)) - r^2 = 0 : a = (b dot b), b = 2b dot (A-C), c = ((A-C) dot (A-C))
	//
	// note that we have a factor of 2 in b, so we can do the following in the quadratic formula to simplify computation:
	// t = (-b +- sqrt(b^2 - 4ac)) / 2a : swap out b with 2h
	// t = (-2h +- sqrt((2h)^2 - 4ac)) / 2a : swap out b with 2h
	// t = (-2h +- 2 * sqrt( h^2 - ac )) / 2a
	// t = (-h +- sqrt( h^2 - ac )) / a : this is our final formula where h = b/2 = b dot (A-C)
	oc := ray.Origin().SubtractVector(s.Center)
	a := ray.Direction().LengthSquared()
	halfB := oc.Dot(ray.Direction())
	c := oc.LengthSquared() - s.Radius*s.Radius

	// Use the discriminant to see whether there are any solutions
	// if discriminant < 0, no solutions
	// if discriminant = 0, 1 solution
	// if discriminant > 0, 2 solutions
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return nil, false, nil
	}
	// At this point, there must be at least one root
	// Find the nearest root that lies in the acceptable range
	squareRootDiscriminant := math.Sqrt(discriminant)
	root := (-halfB - squareRootDiscriminant) / a
	if root < tMin || tMax < root {
		// this root does not fall into the acceptable range; try the other root instead
		root = (-halfB + squareRootDiscriminant) / a
		if root < tMin || tMax < root {
			return nil, false, nil
		}
	}

	hitRecord := &HitRecord{}
	hitRecord.T = root
	hitRecord.P = ray.At(hitRecord.T)
	outwardNormal, err := hitRecord.P.SubtractVector(s.Center).DivideFloat(s.Radius)
	if err != nil {
		return nil, false, errors.New("could not find the normal vector")
	}
	hitRecord.SetFaceNormal(ray, outwardNormal)

	return hitRecord, true, nil
}
