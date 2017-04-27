package graphics

import (
	"image/color"
	"image"
)

// Represents a scene configuration for the ray-tracer; provides objects and materials which compose
// the scene structure and layout
type Scene struct {
}

// Traces an RGBA image of the given dimensions using the given scene configuration
func RayTrace(dimensions image.Rectangle, scene *Scene) (*image.RGBA) {
	result := image.NewRGBA(dimensions)

	width := result.Rect.Size().X
	height := result.Rect.Size().Y

	// for every pixel in the resultant image:
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// project a ray into the image and compute it's final color
			ray := computeRay(x, y)
			color := rayTrace(ray, 0, scene)

			result.Set(x, y, color)
		}
	}

	return result
}

// Concurrently traces an RGBA image of the given dimensions using the given scene configuration
func RayTraceConcurrent(dimensions image.Rectangle, scene *Scene) (*image.RGBA) {
	type ColorAndPoint struct {
		X, Y  int
		Color color.RGBA
	}

	result := image.NewRGBA(dimensions)

	width := result.Rect.Size().X
	height := result.Rect.Size().Y

	colors := make(chan ColorAndPoint)

	// for every pixel in the resultant image:
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// concurrently compute colors over the image dimensions
			go func() {
				// project a ray into the image and compute it's final color;
				// capture it's input point as well so we can compose in parallel
				ray := computeRay(x, y)
				colors <- ColorAndPoint{
					X:     x,
					Y:     y,
					Color: rayTrace(ray, 0, scene),
				}
			}()
		}
	}

	// TODO: compose the color values into the image

	return result
}

// Projects a ray into the screen from the given (x, y) coordinates.
func computeRay(x, y int) Ray {
	panic("Not yet implemented")
}

// Recursively traces a color from the given ray into the given scene configuration
func rayTrace(ray Ray, depth int, scene *Scene) color.RGBA {
	panic("Not yet implemented")
}
