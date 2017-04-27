package graphics

import "image/color"

// Represents a scene configuration for the ray-tracer; provides objects and materials which compose
// the scene structure and layout
type Scene struct {
	Camera          Vector     // The position of the 'camera' within the scene
	Objects         []Object   // The objects composing the scene itself (cubes, spheres, etc)
	BackgroundColor color.RGBA // The background color of the scene
}

// Represents a specific object within a scene
type Object struct {
	Material Material // The material surface of the object
}

// Creates a new sphere at the given position with the given radius and material
func NewSphere(position Vector, radius float64, material Material) Object {
	panic("Not yet implemented")
}

// Creates a new cube at the given position with the given cubed-size and material
func NewCube(position Vector, size float64, material Material) Object {
	panic("Not yet implemented")
}

// Calculates whether the object intersects with the given ray, and computes the hit and normal
func (object *Object) Intersects(ray Ray) (intersects bool, hit, normal Vector) {
	panic("Not yet implemented")
}

// Describes the material surface of an object within a scene
type Material struct {
	DiffuseColor color.RGBA
}

// Some commonly used materials
var (
	Green = Material{DiffuseColor: NewColor(0, 255, 0)}
	Blue  = Material{DiffuseColor: NewColor(0, 0, 255)}
)

// Creates a new RGBA color with the given color channel values
func NewColor(r, g, b uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: 255 }
}
