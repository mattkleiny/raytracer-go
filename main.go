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
		Camera: Camera{
			Position:    NewVec(0, 50, -50),
			FieldOfView: 70.0, // 70°
		},
		Objects: []Object{
			NewSphere(NewVec(0, 0, 0), 8, Green),
			NewCube(NewVec(-10, 0, 0), 8, Blue),
		},
		Lights: []Light{
			NewLight(NewVec(-50, 50, -50), NewVec(0, 0, 0), 1.0),
		},
		BackgroundColor: NewColor(255, 255, 255),
	}

	// Trace the scene into an image so it can be rendered to file
	image := scene.RayTraceToImage(image.Rect(0, 0, 800, 600))

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
