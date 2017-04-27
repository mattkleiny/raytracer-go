package main

import (
	"github.com/xeusalmighty/raytracer/graphics"
	"image"
	"os"
	"image/jpeg"
)

func main() {
	// build a simple scene for configuration
	scene := graphics.Scene{}

	// trace an image out of the scene configuration
	image := graphics.RayTrace(image.Rect(0, 0, 800, 600), &scene)

	// and render the image to file
	encodeToJpg(image, "output.jpg")
}

// Encodes the given image to a .JPG file with the given name
func encodeToJpg(image *image.RGBA, name string) {
	output, err := os.Create(name)

	if err != nil {
		panic(err)
	}
	defer output.Close()

	jpeg.Encode(output, image, &jpeg.Options{jpeg.DefaultQuality})
}
