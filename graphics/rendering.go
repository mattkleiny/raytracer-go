package graphics

import (
	"math"
	"image"
	"image/color"
)

const MaxDepth = 3     // The maximum depth for trace recursion
const LightBias = 1e-4 // suitable small directional bias for light calculations

// Ray-traces an RGBA image of the given dimensions using the given scene configuration
func (scene *Scene) RayTraceToImage(dimensions image.Rectangle) (*image.RGBA) {
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

	// projects a ray into the scene at the given (x, y) image coordinates
	projectRay := func(x, y int) Ray {
		// manually cast to floating point; because go-lang
		fx := float64(x)
		fy := float64(y)

		// compute pixel camera (x, y) coordinates
		pX := (2*((fx+0.5)/fwidth) - 1) * angle * aspectRatio
		pY := 1 - 2*((fy+0.5)/fheight)*angle

		origin := V(0, 0, 0)
		direction := V(pX, pY, -1)

		return R(origin, direction)
	}

	// for every pixel in the resultant image
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// project a ray into the scene and compute it's final color
			ray := projectRay(x, y)
			color := scene.trace(ray, 0, MaxDepth)

			result.Set(x, y, convertToRGBA(color))
		}
	}

	return result
}

// Recursively traces a ray from the given the scene and computes it's resultant color
func (scene *Scene) trace(ray Ray, depth int, maxDepth int) (Vector) {
	// determines the closest object to the ray origin and computes the intersect hit and normal
	findIntersectingObject := func(ray Ray) (result Object, hit, normal Vector) {
		nearest := math.MaxFloat64 // the nearest intersection

		for _, object := range scene.Objects {
			distance := math.MaxFloat32

			// determine if the ray projected from the camera intersected with the object
			if object.Intersects(ray, &distance) {
				if distance < nearest {
					nearest = distance
					result = object
				}
			}
		}

		if result != nil {
			// calculate hit and normal vectors
			hit = ray.Origin.Add(ray.Position(nearest))
			normal = hit.Sub(result.GetPosition()).Normalize()
		}

		return
	}

	// find the first intersecting object within the scene
	object, hit, normal := findIntersectingObject(ray)
	if object == nil {
		return scene.BackgroundColor
	}

	material := object.GetMaterial()
	sampledColor := V(0, 0, 0) // the resultant color

	// compute reflection/refraction up to a certain depth
	if material.Transparency > 0 || material.Reflectivity > 0 && depth < maxDepth {
		panic("Not yet implemented")
	}

	// compute diffuse illumination, accounting for light sources and shadows
	for _, light := range scene.Lights {
		// project a ray from the hit point, accounting for a small bias in direction, toward
		// the light position; we then determine whether another object occludes the light source and
		// project a shadow if it does
		transmission := V(1, 1, 1)
		lightRay := R(hit.Add(normal.MulS(LightBias)), light.Position.Sub(hit))

		for _, other := range scene.Objects {
			distance := math.MaxFloat64

			if other.Intersects(lightRay, &distance) {
				transmission = V(0, 0, 0)
				break
			}
		}

		sampledColor = sampledColor.Add(material.Diffuse).Mul(transmission).MulS(math.Max(0, normal.Dot(lightRay.Direction))).Mul(light.Emission)
	}

	return sampledColor
}

// Converts a vector 3 in floating point range (0.0 to 1.0) to a color with the given channel values
func convertToRGBA(vec Vector) color.RGBA {
	// Clamps a floating point (0..1) into a (0..255) uint8
	clamp := func(value float64) uint8 {
		if value == 1.0 {
			return 255
		}
		return uint8(math.Floor(value * 256.0))
	}

	return color.RGBA{
		R: clamp(vec.X),
		G: clamp(vec.Y),
		B: clamp(vec.Z),
		A: 255,
	}
}
