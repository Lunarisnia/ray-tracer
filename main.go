package main

import (
	"fmt"

	"github.com/Lunarisnia/ray-tracer/internal/ppm"
	"github.com/Lunarisnia/vecm"
)

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 300
	imageHeight := int(float64(imageWidth) / aspectRatio)

	image := ppm.New(imageWidth, imageHeight)
	for j := range imageHeight {
		for i := range imageWidth {
			image.Insert(vecm.Vector3f{
				X: float64(i) / float64(imageWidth),
				Y: float64(j) / float64(imageHeight),
				Z: 0,
			})
		}
	}
	image.Write("./test.ppm")
	fmt.Println("Hello, World")
}
