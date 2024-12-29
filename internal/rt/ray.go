package rt

import (
	"github.com/Lunarisnia/vecm"
	"github.com/Lunarisnia/vecm/vector3f"
)

type Ray struct {
	Origin    vecm.Vector3f
	Direction vecm.Vector3f
}

func (r *Ray) At(t float64) vecm.Vector3f {
	dir := vector3f.Scale(r.Direction, t)
	return vector3f.Add(r.Origin, dir)
}
