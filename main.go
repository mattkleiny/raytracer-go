// Copyright 2017, the project authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	. "github.com/mattkleiny/raytracer-go/graphics"
)

var (
	// command line flags and arguments
	filenameFlag = flag.String("filename", "output.jpg", "filename of the resultant .jpg")
	widthFlag    = flag.Int("width", 1000, "The width of the image to create")
	heightFlag   = flag.Int("height", 1000, "The height of the image to create")
	encoder      = encodeToJpg

	// the scene that is being rendered
	scene = Scene{
		Camera: Camera{
			FieldOfView: 75.0, // 75°
		},
		Objects: []Object{
			&Sphere{
				Center: NewVec(5.0, -1, -15),
				Radius: 2,
				Material: Material{
					Diffuse:      Blue,
					Reflectivity: 1.0,
					Transparency: 0.5,
				},
			},
			&Sphere{
				Center: NewVec(3.0, 0, -35),
				Radius: 2,
				Material: Material{
					Diffuse:      Green,
					Reflectivity: 1.0,
					Transparency: 0.5,
				},
			},
			&Sphere{
				Center: NewVec(-5.5, 0, -15),
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
				Position: NewVec(-20, 30, 20),
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
	encoder(image, *filenameFlag)
}

// Parses the command line and reports any errors
func parseCommandLine() {
	flag.Parse()

	if *filenameFlag == "" {
		flag.Usage()
		log.Fatal("No filename specified")
	}

	if strings.HasSuffix(*filenameFlag, "png") {
		encoder = encodeToPng
	} else if strings.HasSuffix(*filenameFlag, "jpg") {
		encoder = encodeToJpg
	} else {
		flag.Usage()
		log.Fatal("Expected either a .jpg or .png file format")
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
		log.Fatal("Failed to create output image: ", name, err)
	}
	defer output.Close()

	jpeg.Encode(output, image, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

// Encodes the given image to a .PNG file with the given name
func encodeToPng(image *image.RGBA, name string) {
	output, err := os.Create(name)

	if err != nil {
		log.Fatal("Failed to create output image: ", name, err)
	}
	defer output.Close()

	png.Encode(output, image)
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[V int64 | float64](values []V) V {
	var s V

	for _, value := range values {
		s += value
	}

	return s
}
