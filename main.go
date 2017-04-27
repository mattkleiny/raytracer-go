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
	widthFlag    = flag.Int("width", 1000, "The width of the image to create")
	heightFlag   = flag.Int("height", 1000, "The height of the image to create")
)

func main() {
	parseCommandLine()

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
	image := scene.RayTraceToImage(image.Rect(0, 0, *widthFlag, *heightFlag))

	// and render the image to file
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
