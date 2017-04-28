package graphics

import (
	"image/color"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// Represents a ray projected into 3d space
type Ray struct {
	Origin    mgl64.Vec3
	Direction mgl64.Vec3
}

// Represents a scene configuration for the ray-tracer; provides objects and materials which compose
// the scene structure and layout
type Scene struct {
	Camera          Camera     // The scene camera
	Lights          []Light    // The world light sources
	Spheres         []Sphere   // The objects composing the scene itself (currently only spheres)
	BackgroundColor mgl64.Vec3 // The background color of the scene
}

// Represents a camera within the scene
type Camera struct {
	FieldOfView float64 // The field of view, in degrees
}

// Represents a light in the scene
type Light struct {
	Position   mgl64.Vec3 // The position of the light
	Emission   mgl64.Vec3 // The emissive color of the light
}

// Describes the material surface of an object within a scene
type Material struct {
	Diffuse      mgl64.Vec3 // The diffuse color of the material
	Transparency float64    // The transparency level of the material (between 0.0 and 1.0 inclusive)
	Reflectivity float64    // The reflectivity level of the material (between 0.0 and 1.0 inclusive)
}

// Represents a sphere that can be placed in a scene
type Sphere struct {
	Center   mgl64.Vec3
	Radius   float64
	Material Material
}

// Shorthand to create a new vector with the given components
func V(x, y, z float64) mgl64.Vec3 {
	return mgl64.Vec3{x, y, z}
}

// Converts a vector 3 in floating point range (0.0 to 1.0) to a color with the given channel values
func ConvertToRGBA(vec mgl64.Vec3) color.RGBA {
	// Clamps a floating point (0..1) into a (0..255) uint8
	clamp := func(value float64) uint8 {
		if value == 1.0 {
			return 255
		}
		return uint8(math.Floor(value * 256.0))
	}

	return color.RGBA{
		R: clamp(vec[0]),
		G: clamp(vec[1]),
		B: clamp(vec[2]),
		A: 255,
	}
}

// Computes a reflected ray at the given hit point
func (ray Ray) Reflect(hit mgl64.Vec3) Ray {
	panic("Not yet implemented")
}

// Computes a refracted ray at the given hit point
func (ray Ray) Refract(hit mgl64.Vec3) Ray {
	panic("Not yet implemented")
}

// Calculates whether the sphere intersects with the given ray and returns the two potential hit
// points on either side of the ray (which may be degenerate and form a single point).
func (sphere *Sphere) Intersects(ray Ray, t0 *float64, t1 *float64) bool {
	l := sphere.Center.Sub(ray.Origin)
	radius2 := sphere.Radius * sphere.Radius

	tca := l.Dot(ray.Direction)

	if tca < 0 {
		return false
	}

	d2 := l.Dot(l) - tca*tca
	if d2 > radius2 {
		return false
	}

	thc := math.Sqrt(radius2 - d2)

	*t0 = tca - thc
	*t1 = tca + thc

	return true
}
