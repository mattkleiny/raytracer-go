package graphics

import (
	"image"
	"image/color"
	"math"
	"sync"
	"github.com/go-gl/mathgl/mgl64"
)

// Concurrently traces an RGBA image of the given dimensions using the given scene configuration
func (scene *Scene) TraceImage(dimensions image.Rectangle) (*image.RGBA) {
	var barrier sync.WaitGroup

	result := image.NewRGBA(dimensions)

	width := result.Rect.Dx()
	height := result.Rect.Dy()

	// Represents a pixel in 2d space; for use in our pixel channel
	type Pixel struct {
		X, Y  int
		Color color.RGBA
	}

	pixels := make(chan Pixel)

	// concurrently compute color information for the given (x, y) coordinates
	traceColorAt := func(x, y int, pixels chan Pixel, barrier sync.WaitGroup) {
		// projects a ray into the scene at the given (x, y) image coordinates
		projectRay := func(x, y int) Ray {
			fov := scene.Camera.FieldOfView

			// manually cast to floating point; because go-lang
			fx := float64(x)
			fy := float64(y)

			fwidth := float64(width)
			fheight := float64(height)

			aspectRatio := fwidth / fheight

			// compute pixel camera (x, y)
			pX := (2*((fx+0.5)/fwidth) - 1) * math.Tan(fov/2*math.Pi/180) * aspectRatio
			pY := 1 - 2*((fy+0.5)/fheight)*math.Tan(fov/2*math.Pi/180)

			origin := scene.Camera.Position
			direction := NewVec(pX, pY, -1).Sub(origin).Normalize()

			return NewRay(origin, direction)
		}

		// project a ray into the image and compute it's final color
		ray := projectRay(x, y)
		color := scene.traceRay(ray, 0, 10)

		// push pixels out via channel
		pixels <- Pixel {x, y, color}
	}

	// for every pixel in the resultant image
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// concurrently compute it's color information
			barrier.Add(1)
			go traceColorAt(x, y, pixels, barrier)
		}
	}

	barrier.Wait() // wait until all pixels are computed

	// compose all pixels into the resultant image
	for pixel := range pixels {
		result.Set(pixel.X, pixel.Y, pixel.Color)
	}

	return result
}

// Recursively traces a color from the given ray into the given scene configuration
func (scene *Scene) traceRay(ray Ray, depth int, maxDepth int) (color color.RGBA) {
	// Determines the closest object to the ray origin and computes the TSect hit and normal
	findClosestObject := func(ray Ray) (distance float64, object Object, hit, normal mgl64.Vec3) {
		// TODO: see if you can clean this up
		for _, o := range scene.Objects {
			// determine if the ray projected from the camera intersected with the object
			i, h, n := o.Intersects(ray)

			if i {
				// if it did, determine if it was the closest object that we intersected with
				Δ := hit.Sub(ray.Origin).Len()

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
	_, object, _, _ := findClosestObject(ray)

	if object == nil {
		return scene.BackgroundColor // no object; project background color
	}

	panic("Not yet implemented")
}
