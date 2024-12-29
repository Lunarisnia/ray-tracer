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
	maxColor := 255.0
	if color.X > maxColor {
		color.X = maxColor
	}
	if color.Y > maxColor {
		color.Y = maxColor
	}
	if color.Z > maxColor {
		color.Z = maxColor
	}

	minColor := 0.0
	if color.X < minColor {
		color.X = minColor
	}
	if color.Y < minColor {
		color.Y = minColor
	}
	if color.Z < minColor {
		color.Z = minColor
	}

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
