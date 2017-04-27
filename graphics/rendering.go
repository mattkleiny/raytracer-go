package graphics

import (
	"image"
	"image/color"
)

// Concurrently traces an RGBA image of the given dimensions using the given scene configuration
func (scene *Scene) TraceImage(dimensions image.Rectangle) (*image.RGBA) {
	// encapsulates color and position for our concurrent step algorithm
	type ColorAndPoint struct {
		X, Y  int
		Color color.RGBA
	}

	// concurrently compute color information for the given (x, y) coordinates
	traceColorAt := func(x, y int, image *image.RGBA) {
		// project a ray into the image and compute it's final color
		ray := NewRay(scene.Camera, x, y)
		color := scene.trace(ray, 0, 10)

		image.Set(x, y, color) // TODO: determine if this is thread-safe
	}

	result := image.NewRGBA(dimensions)

	width := result.Rect.Dx()
	height := result.Rect.Dy()

	// for every pixel in the resultant image
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// concurrently compute it's color information
			go traceColorAt(x, y, result)
		}
	}

	return result
}

// Recursively traces a color from the given ray into the given scene configuration
func (scene *Scene) trace(ray Ray, depth int, maxDepth int) color.RGBA {
	// Determines the closest object to the ray origin and computes the TSect hit and normal
	findClosestObject := func(ray Ray) (distance float64, object Object, hit, normal Vector) {
		// TODO: see if you can clean this up
		for _, o := range scene.Objects {
			// determine if the ray projected from the camera intersected with the object
			i, h, n := o.Intersects(ray)

			if i {
				// if it did, determine if it was the closest object that we intersected with
				Δ := ray.Origin.DistanceSqr(hit)

				if distance < Δ {
					// if it is, retain the hit point and normal information
					distance = Δ

					hit = h
					normal = n
					object = o
				}
			}
		}
		return
	}

	// for each of the objects within the scene
	distance, object, hit, normal := findClosestObject(ray)

	if object == nil {
		return scene.BackgroundColor // no object; project background color
	}

	panic("Not yet implemented")
}
