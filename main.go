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
	// command line flags and arguments
	filenameFlag = flag.String("filename", "output.jpg", "filename of the resultant .jpg")
	widthFlag    = flag.Int("width", 1000, "The width of the image to create")
	heightFlag   = flag.Int("height", 1000, "The height of the image to create")

	// the scene that is being rendered
	scene = Scene{
		Camera: Camera{
			FieldOfView: 75.0, // 75°
		},
		Objects: []Object{
			&Sphere{
				Center: V(5.0, -1, -15),
				Radius: 2,
				Material: Material{
					Diffuse:      Blue,
					Reflectivity: 1.0,
					Transparency: 0.5,
				},
			},
			&Sphere{
				Center: V(3.0, 0, -35),
				Radius: 2,
				Material: Material{
					Diffuse:      Green,
					Reflectivity: 1.0,
					Transparency: 0.5,
				},
			},
			&Sphere{
				Center: V(-5.5, 0, -15),
				Radius: 3,
				Material: Material{
					Diffuse:      Red,
					Reflectivity: 1.0,
					Transparency: 0.5,
				},
			},
		},
		Lights: []Light{
			{
				Position: V(-20, 30, 20),
				Emissive: White,
			},
		},
		BackgroundColor: White,
	}
)

// Entry point for the application
func main() {
	parseCommandLine()
	image := scene.Render(image.Rect(0, 0, *widthFlag, *heightFlag))
	encodeToJpg(image, *filenameFlag)
}

// Parses the command line and reports any errors
func parseCommandLine() {
	flag.Parse()

	if *filenameFlag == "" {
		flag.Usage()
		log.Fatal("No filename specified")
	}

	if *widthFlag == 0 {
		flag.Usage()
		log.Fatal("No width specified")
	}

	if *heightFlag == 0 {
		flag.Usage()
		log.Fatal("No height specified")
	}
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
