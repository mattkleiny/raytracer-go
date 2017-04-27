package main

import (
	"os"
	"image"
	"image/jpeg"
	. "github.com/xeusalmighty/raytracer/graphics"
)

func main() {
	// Create a simple scene with a few objects
	scene := Scene{
		Camera: NewVector(0, 50, -50),
		Objects: []Object{
			NewSphere(NewVector(0, 0, 0), 8, Green),
			NewCube(NewVector(-10, 0, 0), 8, Blue),
		},
		BackgroundColor: NewColor(255, 255, 255),
	}

	// Trace the scene into an image so it can be rendered to file
	image := scene.TraceImage(image.Rect(0, 0, 800, 600))

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
