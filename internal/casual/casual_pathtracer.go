package casual

import (
	"fmt"

	"github.com/Lunarisnia/ray-tracer/internal/ppm"
	"github.com/Lunarisnia/ray-tracer/internal/rt"
	"github.com/Lunarisnia/vecm"
	"github.com/Lunarisnia/vecm/vector3f"
)

func GetRayColor(r rt.Ray) vecm.Vector3f {
	return vecm.Vector3f{}
}

func CasualRender() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 300
	imageHeight := int(float64(imageWidth) / aspectRatio)

	origin := vecm.Vector3f{}

	image := ppm.New(imageWidth, imageHeight)
	for j := range imageHeight {
		for i := range imageWidth {
			uv := vecm.Vector3f{
				X: float64(i) / float64(imageWidth),
				Y: 1.0 - (float64(j) / float64(imageHeight)),
				Z: 1.0,
			}
			uv = vector3f.Scale(uv, 2.0)
			uv = vector3f.Subtract(uv, vecm.Vector3f{X: 1.0, Y: 1.0, Z: 0})

			rayDirection := vector3f.Normalize(vector3f.Subtract(uv, origin))

			rayColor := GetRayColor(rt.Ray{
				Origin:    origin,
				Direction: rayDirection,
			})

			image.Insert(rayColor)
		}
	}
	image.Write("./test.ppm")
	fmt.Println("Finished")
}
