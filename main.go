package main

import (
	"github.com/xeusalmighty/raytracer/graphics"
	"image"
	"os"
	"image/jpeg"
)

func main() {
	image := graphics.RayTrace(image.Rect(0, 0, 800, 600))

	output, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	jpeg.Encode(output, image, &jpeg.Options{jpeg.DefaultQuality})
}
