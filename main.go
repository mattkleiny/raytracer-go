package main

import (
	"flag"
	"image"
	"os"
	"log"
	"image/jpeg"
	. "github.com/xeusalmighty/raytracer/graphics"
)

var (
	filenameFlag = flag.String("filename", "output.jpg", "filename of the resultant .jpg")
)

func main() {
	// parse command line
	flag.Parse()

	if *filenameFlag == "" {
		flag.Usage()
		log.Fatal("No filename specified")
	}

	// create a simple scene with a few objects
	scene := Scene{
		Camera: Camera{
			Position:    V(0, 50, -50),
			FieldOfView: 75.0, // 75°
		},
		Objects: []Object{
			&Sphere{
				Position: V(0, 0, 0),
				Radius:   8.0,
				Material: Material{
					Diffuse: V(0, 1, 0),
					IsGlass: true,
				},
			},
			&Cube{
				Position: V(-10, 0, 0),
				Size:     8.0,
				Material: Material{
					Diffuse: V(0, 0, 1),
					IsGlass: false,
				},
			},
		},
		Light: Light{
			Position:   V(-50, 50, -50),
			Direction:  V(0, 0, 0),
			Brightness: 1.0,
		},
		Color: V(1.0, 1.0, 1.0),
	}

	// trace the scene into an image so it can be rendered to file
	image := scene.RayTraceToImage(image.Rect(0, 0, 800, 600))

	// and render the image to file
	encodeToJpg(image, *filenameFlag)
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
