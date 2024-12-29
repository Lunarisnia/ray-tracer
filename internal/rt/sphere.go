package rt

import (
	"math"

	"github.com/Lunarisnia/vecm"
	"github.com/Lunarisnia/vecm/vector3f"
)

type Sphere struct {
	Center vecm.Vector3f
	Radius float64
}

func (s *Sphere) CasualIntersect(r Ray) bool {
	center := vector3f.Subtract(s.Center, r.Origin)
	a := vector3f.Dot(r.Direction, r.Direction)
	b := vector3f.Dot(r.Direction, center)
	c := vector3f.Dot(center, center) - (s.Radius * s.Radius)

	discriminant := b*b - a*c
	if discriminant < 0.0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)
	dist := (b - sqrtd) / a

	nearestDistance := 0.001
	// TODO: Need hit info to set the max dist
	if dist <= nearestDistance || dist >= 1988888888.0 {
		dist = (b + sqrtd) / a
		if dist <= nearestDistance || dist >= 198788888.0 {
			return false
		}
	}

	return true
}

func (s *Sphere) Intersect(r *Ray) bool {
	oc := vector3f.Subtract(s.Center, r.Origin)
	a := vector3f.Dot(r.Direction, r.Direction)
	b := -2.0 * vector3f.Dot(r.Direction, oc)
	c := vector3f.Dot(oc, oc) - (s.Radius * s.Radius)
	discriminant := b*b - 4*a*c

	return discriminant >= 0.0
}
