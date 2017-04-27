package graphics

import (
	"image/color"
	"image"
	"math"
)

// Serially traces an RGBA image of the given dimensions using the given scene configuration
func (scene *Scene) TraceToImage(dimensions image.Rectangle) (*image.RGBA) {
	result := image.NewRGBA(dimensions)

	width := result.Rect.Size().X
	height := result.Rect.Size().Y

	// for every pixel in the resultant image:
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// project a ray into the image and compute it's final color
			ray := ProjectRay(scene.Camera, x, y)
			color := scene.traceRay(ray, 0)

			result.Set(x, y, color)
		}
	}

	return result
}

// Concurrently traces an RGBA image of the given dimensions using the given scene configuration
func (scene *Scene) TraceToImageConcurrent(dimensions image.Rectangle) (*image.RGBA) {
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
				ray := ProjectRay(scene.Camera, x, y)
				colors <- ColorAndPoint{
					X:     x,
					Y:     y,
					Color: scene.traceRay(ray, 0),
				}
			}()
		}
	}

	// TODO: compose the color values into the image

	return result
}

// Recursively traces a color from the given ray into the given scene configuration
func (scene *Scene) traceRay(ray Ray, depth int) color.RGBA {
	minDistance := math.MaxFloat64 // the minimum distance between the ray and the intersected object

	hitPoint := NewVector(0, 0, 0)
	hitNormal := NewVector(0, 0, 0)

	var intersectedObject *Object = nil

	// for each of the objects within the scene
	for _, object := range scene.Objects {
		intersects, hit, normal := object.Intersects(ray)
		if intersects {
			// determine if this object is in-front of other objects
			distance := ray.Origin.DistanceSqr(hit)
			if distance < minDistance {
				hitPoint = hit
				hitNormal = normal
				minDistance = distance
				intersectedObject = object
			}
		}
	}

	if intersectedObject == nil {
		return scene.BackgroundColor // no object; project background color
	}

	panic("Not yet implemented")
}
