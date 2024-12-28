package main

import (
	"fmt"

	"github.com/Lunarisnia/ray-tracer/internal/ppm"
	"github.com/Lunarisnia/vecm"
)

func main() {
	image := ppm.New(100, 100)
	for range 100 * 100 {
		image.Insert(vecm.Vector3f{
			X: 100.0 / 255.0,
			Y: 100.0 / 255.0,
			Z: 0,
		})
	}
	image.Write("./test.ppm")
	fmt.Println("Hello, World")
}
