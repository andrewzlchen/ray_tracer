package raytracer_test

import (
	"testing"

	rt "github.com/andrewzlchen/raytracer/src"

	"github.com/stretchr/testify/assert"
)

func TestSphere_Hit(t *testing.T) {
	type args struct {
		ray  *rt.Ray
		tMin float64
		tMax float64
	}
	tests := []struct {
		name         string
		s            *rt.Sphere
		args         args
		wantedDidHit bool
		wantedError  error
	}{
		{
			name: "zero direction ray should not intersect sphere",
			s:    rt.NewSphere(rt.NewVec3(0, 0, 2), 1),
			args: args{
				ray:  rt.NewRay(rt.NewVec3(0, 0, 0), rt.NewVec3(0, 0, 0)),
				tMin: 0,
				tMax: 100,
			},
			wantedDidHit: false,
		},
		{
			name: "1,1,1 ray should intersect sphere with center = 1,1,1",
			s:    rt.NewSphere(rt.NewVec3(1, 1, 1), 1),
			args: args{
				ray:  rt.NewRay(rt.NewVec3(0, 0, 0), rt.NewVec3(1, 1, 1)),
				tMin: 0,
				tMax: 100,
			},
			wantedDidHit: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, didHit, err := tt.s.Hit(tt.args.ray, tt.args.tMin, tt.args.tMax)
			if tt.wantedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantedError, err)
			} else {
				assert.Equal(t, tt.wantedDidHit, didHit)
			}
		})
	}
}
