package graphics

import (
	"math"
	"image"
	"github.com/go-gl/mathgl/mgl64"
)

// Ray-traces an RGBA image of the given dimensions using the given scene configuration
func (scene *Scene) RayTraceToImage(dimensions image.Rectangle) (*image.RGBA) {
	const MaxDepth = 3 // The maximum depth for trace recursion

	result := image.NewRGBA(dimensions)

	width := result.Rect.Dx()
	height := result.Rect.Dy()

	// compute aspect ratio of the image
	fwidth := float64(width)
	fheight := float64(height)
	aspectRatio := fwidth / fheight

	// pre-compute field of view and camera angle
	fov := scene.Camera.FieldOfView
	angle := math.Tan(fov / 2 * math.Pi / 180)

	// compute color information for the given (x, y) coordinates
	traceColorAt := func(x, y int, image *image.RGBA) {
		// projects a ray into the scene at the given (x, y) image coordinates
		projectRay := func(x, y int) Ray {
			// manually cast to floating point; because go-lang
			fx := float64(x)
			fy := float64(y)

			// compute pixel camera (x, y) coordinates
			pX := (2*((fx+0.5)/fwidth) - 1) * angle * aspectRatio
			pY := 1 - 2*((fy+0.5)/fheight)*angle

			origin := V(0, 0, 0)
			direction := V(pX, pY, -1).Normalize()

			return Ray{origin, direction}
		}

		// project a ray into the scene and compute it's final color
		ray := projectRay(x, y)
		color := scene.trace(ray, 0, MaxDepth)

		image.Set(x, y, ConvertToRGBA(color))
	}

	// for every pixel in the resultant image
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// compute it's color information
			traceColorAt(x, y, result)
		}
	}

	return result
}

// Recursively traces a ray from the given the scene and computes it's resultant color
func (scene *Scene) trace(ray Ray, depth int, maxDepth int) (mgl64.Vec3) {
	// determines the closest sphere to the ray origin and computes the intersect hit and normal
	findIntersectingSphere := func(ray Ray) (result *Sphere, hit, normal mgl64.Vec3) {
		tnear := math.MaxFloat64 // the nearest intersection

		for _, sphere := range scene.Spheres {
			t0, t1 := math.MaxFloat64, math.MaxFloat64

			// determine if the ray projected from the camera intersected with the object
			if sphere.Intersects(ray, &t0, &t1) {
				if t0 < 0 {
					t0 = t1
				}

				if t0 < tnear {
					tnear = t0
					result = &sphere
				}
			}
		}

		if result != nil {
			// calculate hit and normal vectors
			hit = ray.Origin.Add(ray.Direction.Mul(tnear))
			normal = hit.Sub(result.Center).Normalize()
		}

		return
	}

	// Computes the fresnel lens constants for the given point and direction
	// See (https://en.wikipedia.org/wiki/Fresnel_lens) for more information.
	fresnel := func(hit, direction mgl64.Vec3) (float64, float64) {
		panic("Not yet implemented")
	}

	// find the first intersecting object within the scene
	sphere, hit, normal := findIntersectingSphere(ray)
	if sphere == nil {
		return scene.BackgroundColor
	}

	material := sphere.Material

	// compute reflection/refraction up to a certain depth
	if material.Transparency > 0 || material.Reflectivity > 0 && depth < maxDepth {
		// compute reflection and refraction colors
		reflection := scene.trace(ray.Reflect(normal), depth+1, maxDepth)
		refraction := scene.trace(ray.Refract(normal), depth+1, maxDepth)

		Kr, Kt := fresnel(normal, ray.Direction)

		return reflection.Mul(Kr).Add(refraction.Mul(1 - Kt))
	}

	// compute diffuse illumination, accounting for light sources and shadows
	for _, light := range scene.Lights {
		// project a ray from the hit point, accounting for a small bias in direction, toward
		// the light position; we then determine whether another sphere occludes the light source and
		// project a shadow if it does
		const Bias = 1e-4 // suitable small directional bias
		lightRay := Ray{hit.Add(normal.Mul(Bias)), light.Position.Sub(hit).Normalize()}

		for _, other := range scene.Spheres {
			t0, t1 := math.MaxFloat64, math.MaxFloat64
			if other.Intersects(lightRay, &t0, &t1) {
				return V(0, 0, 0) // covered in shadow
			}
		}
	}

	return material.Diffuse
}
