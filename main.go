package main

import (
	"github.com/Lunarisnia/ray-tracer/internal/ppm"
	"github.com/Lunarisnia/ray-tracer/internal/rt"
	"github.com/Lunarisnia/vecm"
	"github.com/Lunarisnia/vecm/vector3f"
)

func GetRayColor(r rt.Ray) vecm.Vector3f {
	return vecm.Vector3f{
		X: 1,
	}
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	focalLength := 1.0
	viewportWidth := 4.0
	viewportHeight := viewportWidth / (float64(imageWidth) / float64(imageHeight))

	viewportU := vecm.Vector3f{X: viewportWidth}
	viewportV := vecm.Vector3f{Y: -viewportHeight}

	// get size of pixel
	pixelDeltaU := vector3f.Divide(viewportU, float64(imageWidth))
	pixelDeltaV := vector3f.Divide(viewportV, float64(imageHeight))

	origin := vecm.Vector3f{}

	topLeftPixel := vector3f.Subtract(origin, vecm.Vector3f{Z: focalLength})
	topLeftPixel = vector3f.Subtract(topLeftPixel, vector3f.Divide(pixelDeltaU, 2.0))
	topLeftPixel = vector3f.Subtract(topLeftPixel, vector3f.Divide(pixelDeltaV, 2.0))

	pixelOffset := vector3f.Add(pixelDeltaU, pixelDeltaV)
	pixelOffset = vector3f.Scale(pixelOffset, 0.5)
	pixel00 := vector3f.Add(topLeftPixel, pixelOffset)

	image := ppm.New(imageWidth, imageHeight)
	for j := range imageHeight {
		for i := range imageWidth {
			pixelCenterU := vector3f.Scale(pixelDeltaU, float64(i))
			pixelCenterV := vector3f.Scale(pixelDeltaV, float64(j))
			pixelCenter := vector3f.Add(pixel00, pixelCenterU)

			pixelCenter = vector3f.Add(pixel00, pixelCenterV)
			rayDirection := vector3f.Normalize(vector3f.Subtract(pixelCenter, origin))

			// TODO: The raycast
			color := GetRayColor(rt.Ray{
				Origin:    origin,
				Direction: rayDirection,
			})

			image.Insert(color)
		}
	}

	image.Write("./test.ppm")
}
