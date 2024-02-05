// Copyright 2017, the project authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package graphics

import (
	"image"
	"math"
)

// Renders an RGBA image of the given dimensions using the given scene configuration via ray-tracing
func (scene *Scene) Render(dimensions image.Rectangle) *image.RGBA {
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

		direction := NewVec(pX, pY, -1)

		return R(Zero, direction)
	}

	// for every pixel in the resultant image
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// project a ray into the scene and compute it's final color
			ray := projectRay(x, y)
			color := scene.sample(ray, 0, MaxTraceDepth)

			result.Set(x, y, color.ConvertToRGBA())
		}
	}

	return result
}

// Samples the scene by projecting a ray and computes it's resultant color
func (scene *Scene) sample(ray Ray, depth, maxDepth int) Color {
	// determines the closest object to the ray origin and computes the intersect hit and a
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
			// calculate hit and a vectors
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
	sampledColor := Black // the resultant color

	// compute reflection/refraction illumination up to a certain depth
	if (material.Transparency > 0 || material.Reflectivity > 0) && depth < maxDepth {
		// computes the fresnel lens effect for reflective/transparent surfaces
		// see https://en.wikipedia.org/wiki/Fresnel_lens
		computeFresnel := func(direction, normal Vector) float64 {
			mix := func(a, b, mix float64) float64 {
				return b*mix + a*(1-mix)
			}

			return mix(math.Pow(1+direction.Dot(normal), 3), 1, 0.1)
		}

		// determine whether we're inside or outside the surface
		inside := false
		if ray.Direction.Dot(normal) > 0 {
			normal = normal.Negate()
			inside = true
		}

		fresnel := computeFresnel(ray.Direction, normal)

		// compute reflective and refractive color by recursively tracing light along reflective and
		// refractive angles; then combine the resultant colours
		reflectionColor := scene.sample(ray.Reflect(hit, normal), depth+1, maxDepth)
		refractionColor := Black

		if material.Transparency > 0 {
			refractionColor = scene.sample(ray.Refract(hit, normal, inside), depth+1, maxDepth)
		}

		sampledColor = reflectionColor.MulS(fresnel).Add(refractionColor.MulS(1 - fresnel).MulS(material.Transparency)).Mul(material.Diffuse)
	} else {
		// compute diffuse illumination, accounting for light sources and shadows
		for _, light := range scene.Lights {
			// project a ray from the hit point, accounting for a small bias in direction, toward
			// the light position; we then determine whether another object occludes the light source and
			// project a shadow if it does
			transmission := White
			lightRay := R(hit.Add(normal.MulS(ε)), light.Position.Sub(hit))

			for _, other := range scene.Objects {
				distance := math.MaxFloat64

				if other.Intersects(lightRay, &distance) {
					transmission = Black // prevent color transmission; this object is in shadow
					break
				}
			}

			// compute shaded color
			sampledColor = sampledColor.Add(material.Diffuse).Mul(transmission).MulS(math.Max(0, normal.Dot(lightRay.Direction))).Mul(light.Emissive)
		}
	}

	return sampledColor
}
