package graphics

import (
	"image/color"
	"image"
)

func RayTrace(dimensions image.Rectangle) (*image.RGBA) {
	result := image.NewRGBA(dimensions)

	width := result.Rect.Size().X
	height := result.Rect.Size().Y

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			ray := computeRay(x, y)
			color := rayTrace(ray, 0)

			result.Set(x, y, color)
		}
	}

	return result
}

func computeRay(x, y int) Ray {
	panic("Not yet implemented")
}

func rayTrace(ray Ray, depth int) color.RGBA { // recursive
	panic("Not yet implemented")
}
