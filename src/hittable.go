package raytracer

import (
	"fmt"
)

// HitRecord is a struct that stores information relevant to a ray hitting a Hittable
type HitRecord struct {
	P, Normal *Vec3
	T         float64
	FrontFace bool
}

// SetFaceNormal sets whether the surface normal should face outwards or inwards
func (hr *HitRecord) SetFaceNormal(ray *Ray, outwardNormal *Vec3) {
	if ray.Direction().Dot(outwardNormal) < 0 {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.MultiplyFloat(-1.0)
	}
}

// Hittable is an interface that types will interface if they are able to be hit by a ray
type Hittable interface {
	Hit(ray *Ray, tMin, tMax float64) (*HitRecord, bool, error)
}

// HittableList is a list of hittable objects
type HittableList struct {
	Objects []Hittable
}

// Clear clears out the list of hittable objects but does not reclaim the allocatted memory
func (hl *HittableList) Clear() {
	hl.Objects = hl.Objects[:0]
}

// Add adds a hittable object to the list of hittable objects
func (hl *HittableList) Add(obj Hittable) {
	hl.Objects = append(hl.Objects, obj)
}

// Hit determines whether the input ray intersects anything
func (hl *HittableList) Hit(ray *Ray, tMin, tMax float64) (*HitRecord, bool, error) {
	hitRecord := &HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, object := range hl.Objects {
		tmpRecord, didHit, err := object.Hit(ray, tMin, closestSoFar)
		if err != nil {
			return nil, false, fmt.Errorf("could not determine if object was hit: %s", err)
		}
		if didHit {
			hitAnything = true
			closestSoFar = tmpRecord.T
			hitRecord = tmpRecord
		}
	}

	return hitRecord, hitAnything, nil
}
