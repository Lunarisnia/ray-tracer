package ppm

import (
	"fmt"
	"math"
	"os"

	"github.com/Lunarisnia/vecm"
	"github.com/Lunarisnia/vecm/vector3f"
)

type PPM struct {
	buffer []string
}

func New(width int, height int) *PPM {
	buffer := make([]string, 0)
	buffer = append(buffer, "P3\n")
	buffer = append(buffer, fmt.Sprintf("%v %v\n", width, height))
	buffer = append(buffer, "255\n")
	return &PPM{
		buffer: buffer,
	}
}

func (p *PPM) Insert(color vecm.Vector3f) {
	color = vector3f.Scale(color, float64(255))

	color.X = math.Round(color.X)
	color.Y = math.Round(color.Y)
	color.Z = math.Round(color.Z)

	p.buffer = append(p.buffer, fmt.Sprintf("%v %v %v\n", color.X, color.Y, color.Z))
}

func (p *PPM) Write(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, str := range p.buffer {
		_, err := file.WriteString(str)
		if err != nil {
			return err
		}
	}

	return nil
}
