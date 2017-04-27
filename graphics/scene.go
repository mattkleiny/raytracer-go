package graphics

import "image/color"

// Represents a scene configuration for the ray-tracer; provides objects and materials which compose
// the scene structure and layout
type Scene struct {
	Camera          Camera     // The main scene camera
	Objects         []Object   // The objects composing the scene itself (cubes, spheres, etc)
	Lights          []Light    // The lights within the world
	BackgroundColor color.RGBA // The background color of the scene
}

// Represents a camera within the world
type Camera struct {
	Position    Vector  // The position of the camera
	FieldOfView float64 // The field of view, in degrees
}

// Represents a light in the world
type Light struct {
	Direction Ray        // The direction of the light, as a ray
	Color     color.RGBA // The color of the light
}

// Represents a specific object within a scene
type Object interface {
	// Determines if the object intersects with the given ray, and computes the hit and normal
	Intersects(ray Ray) (intersects bool, hit, normal Vector)
}

// Represents a sphere that can be placed in a scene
type Sphere struct {
	Position Vector
	Radius   float64
	Material Material
}

// Represents a cube that can be placed in a scene
type Cube struct {
	Position Vector
	Size     float64
	Material Material
}

// Describes the material surface of an object within a scene
type Material struct {
	Diffuse color.RGBA // The diffuse color of the material
	IsGlass bool       // If the material is glass, it possesses unique reflection properties
}

// Some commonly used materials
var (
	Green = Material{Diffuse: NewColor(0, 255, 0)}
	Blue  = Material{Diffuse: NewColor(0, 0, 255)}
)

// Creates a new sphere at the given position with the given radius and material
func NewSphere(position Vector, radius float64, material Material) *Sphere {
	return &Sphere{Position: position, Radius: radius, Material: material}
}

// Creates a new cube at the given position with the given cubed-size and material
func NewCube(position Vector, size float64, material Material) *Cube {
	return &Cube{Position: position, Size: size, Material: material}
}

// Calculates whether the sphere intersects with the given ray, and computes the hit and normal
func (sphere *Sphere) Intersects(ray Ray) (intersects bool, hit, normal Vector) {
	panic("Not yet implemented")
}

// Calculates whether the cube intersects with the given ray, and computes the hit and normal
func (cube *Cube) Intersects(ray Ray) (intersects bool, hit, normal Vector) {
	panic("Not yet implemented")
}

// Creates a new RGBA color with the given channel values
func NewColor(r, g, b uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: 255 }
}
