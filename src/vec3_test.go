package raytracer_test

import (
	"testing"

	raytracer "github.com/andrewzlchen/raytracer/src"
	"github.com/stretchr/testify/assert"
)

func TestVec3(t *testing.T) {
	t.Run("Creating a Vec3 works", func(t *testing.T) {
		vec := raytracer.NewVec3(0, 0, 0)
		assert.Equal(t, vec.X, 0, "vector.x should be 0")
		assert.Equal(t, vec.Y, 0, "vector.y should be 0")
		assert.Equal(t, vec.Z, 0, "vector.z should be 0")
	})
}

func TestVec3Add(t *testing.T) {
	t.Run("Adding using the Vec3 Add method", func(t *testing.T) {
		for _, tc := range []struct {
			desc         string
			a, b, result *raytracer.Vec3
		}{
			{desc: "two positive", a: raytracer.NewVec3(1, 1, 1), b: raytracer.NewVec3(1, 1, 1), result: raytracer.NewVec3(2, 2, 2)},
			{desc: "one positive, one negative", a: raytracer.NewVec3(1, 1, 1), b: raytracer.NewVec3(-1, -1, -1), result: raytracer.NewVec3(0, 0, 0)},
			{desc: "two negative", a: raytracer.NewVec3(-1, -1, -1), b: raytracer.NewVec3(-1, -1, -1), result: raytracer.NewVec3(-2, -2, -2)},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				out := tc.a.AddVector(tc.b)
				assert.Equal(t, tc.result, out)
			})
		}
	})
}

func TestVec3Multiply(t *testing.T) {
	t.Run("Multiplying two vectors", func(t *testing.T) {
		for _, tc := range []struct {
			desc         string
			a, b, result *raytracer.Vec3
		}{
			{desc: "two unit vectors", a: raytracer.NewVec3(1, 1, 1), b: raytracer.NewVec3(1, 1, 1), result: raytracer.NewVec3(1, 1, 1)},
			{desc: "a vector and a zero vector", a: raytracer.NewVec3(1, 1, 1), b: raytracer.NewVec3(0, 0, 0), result: raytracer.NewVec3(0, 0, 0)},
			{desc: "one negative and one positive", a: raytracer.NewVec3(-1, -1, -1), b: raytracer.NewVec3(2, 2, 2), result: raytracer.NewVec3(-2, -2, -2)},
			{desc: "two negatives", a: raytracer.NewVec3(-1, -1, -1), b: raytracer.NewVec3(-1, -1, -1), result: raytracer.NewVec3(1, 1, 1)},
			{desc: "random numbers", a: raytracer.NewVec3(1, 2, 3), b: raytracer.NewVec3(4, 5, 6), result: raytracer.NewVec3(4, 10, 18)},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				out := tc.a.MultiplyVector(tc.b)
				assert.Equal(t, tc.result, out)
			})
		}
	})

	t.Run("Multiplying vectors by a constant", func(t *testing.T) {
		for _, tc := range []struct {
			desc      string
			a, result *raytracer.Vec3
			b         float64
		}{
			{desc: "multiplying a vector by 0", a: raytracer.NewVec3(1, 1, 1), b: float64(0), result: raytracer.NewVec3(0, 0, 0)},
			{desc: "multiplying a vector by -3", a: raytracer.NewVec3(1, 1, 1), b: float64(-3), result: raytracer.NewVec3(-3, -3, -3)},
			{desc: "multiplying negative unit vectory by a negative constant", a: raytracer.NewVec3(-1, -1, -1), b: float64(3), result: raytracer.NewVec3(-3, -3, -3)},
			{desc: "multiplying negative unit vector by a positive constant", a: raytracer.NewVec3(-1, -1, -1), b: float64(10), result: raytracer.NewVec3(-10, -10, -10)},
			{desc: "multiplying vector by 1", a: raytracer.NewVec3(1, 2, 3), b: float64(1), result: raytracer.NewVec3(1, 2, 3)},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				out := tc.a.MultiplyFloat(tc.b)
				assert.Equal(t, tc.result, out)
			})
		}
	})
}

func TestVec3Divide(t *testing.T) {
	t.Run("Dividing two vectors", func(t *testing.T) {
		for _, tc := range []struct {
			desc         string
			a, b, result *raytracer.Vec3
			isError      bool
		}{
			{desc: "dividing a vector by a zero vector", a: raytracer.NewVec3(1, 1, 1), b: raytracer.NewVec3(0, 0, 0), isError: true},
			{desc: "dividing two unit vectors", a: raytracer.NewVec3(1, 1, 1), b: raytracer.NewVec3(1, 1, 1), result: raytracer.NewVec3(1, 1, 1)},
			{desc: "dividing one negative by one positive", a: raytracer.NewVec3(-1, -1, -1), b: raytracer.NewVec3(2, 2, 2), result: raytracer.NewVec3(-1.0/2, -1.0/2, -1.0/2)},
			{desc: "dividing two negatives", a: raytracer.NewVec3(-1, -1, -1), b: raytracer.NewVec3(-1, -1, -1), result: raytracer.NewVec3(1, 1, 1)},
			{desc: "dividing and getting fractional vectors", a: raytracer.NewVec3(1, 1, 1), b: raytracer.NewVec3(2, 2, 2), result: raytracer.NewVec3(0.5, 0.5, 0.5)},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				out, err := tc.a.DivideVector(tc.b)
				if tc.isError {
					assert.Nil(t, out)
					assert.EqualError(t, err, "cannot divide by 0")
				} else {
					assert.Nil(t, err)
					assert.Equal(t, tc.result, out)
				}
			})
		}
	})

	t.Run("Dividing vectors by a constant", func(t *testing.T) {
		for _, tc := range []struct {
			desc      string
			a, result *raytracer.Vec3
			b         float64
			isError   bool
		}{
			{desc: "dividing a vector by 0", a: raytracer.NewVec3(1, 1, 1), b: float64(0), isError: true},
			{desc: "dividing a vector by -3", a: raytracer.NewVec3(1, 1, 1), b: float64(-3), result: raytracer.NewVec3(-float64(1)/3, -float64(1)/3, -float64(1)/3)},
			{desc: "dividing negative unit vectory by a negative constant", a: raytracer.NewVec3(-1, -1, -1), b: float64(3), result: raytracer.NewVec3(-1.0/3, -1.0/3, -1.0/3)},
			{desc: "dividing negative unit vector by a positive constant", a: raytracer.NewVec3(-1, -1, -1), b: float64(10), result: raytracer.NewVec3(-0.1, -0.1, -0.1)},
			{desc: "dividing vector by 1", a: raytracer.NewVec3(1, 2, 3), b: float64(1), result: raytracer.NewVec3(1, 2, 3)},
		} {
			t.Run(tc.desc, func(t *testing.T) {
				out, err := tc.a.DivideFloat(tc.b)
				if tc.isError {
					assert.Nil(t, out)
					assert.EqualError(t, err, "cannot divide by 0")
				} else {
					assert.Equal(t, tc.result, out)
				}
			})
		}
	})
}
