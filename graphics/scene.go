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
	Camera          Camera     // The main scene camera
	Objects         []Object   // The objects composing the scene itself (cubes, spheres, etc)
	Lights          []Light    // The lights within the world
	BackgroundColor color.RGBA // The background color of the scene
}

// Represents a camera within the world
type Camera struct {
	Position    mgl64.Vec3 // The position of the camera
	FieldOfView float64    // The field of view, in degrees
}

// Represents a light in the world
type Light struct {
	Position   mgl64.Vec3 // The position of the light
	Direction  mgl64.Vec3 // The direction of the light
	Brightness float64    // The brightness of the light
}

// Represents a specific object within a scene
type Object interface {
	// Determines if the object intersects with the given ray, and computes the hit and normal
	Intersects(ray Ray) (intersects bool, hit, normal mgl64.Vec3)
	GetMaterial() *Material // The object's material
}

// Represents a sphere that can be placed in a scene
type Sphere struct {
	Position mgl64.Vec3
	Radius   float64
	Material Material
}

// Represents a cube that can be placed in a scene
type Cube struct {
	Position mgl64.Vec3
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
	Origin = NewVec(0, 0, 0) // Origin in world-space
	Green  = Material{Diffuse: NewColor(0, 255, 0)}
	Blue   = Material{Diffuse: NewColor(0, 0, 255)}
)

// Creates a new vector with the given components
func NewVec(x, y, z float64) mgl64.Vec3 {
	return mgl64.Vec3{x, y, z}
}

// Creates a new ray with the given components
func NewRay(origin, direction mgl64.Vec3) Ray {
	return Ray{Origin: origin, Direction: direction}
}

// Creates a new sphere at the given position with +the given radius and material
func NewSphere(position mgl64.Vec3, radius float64, material Material) *Sphere {
	return &Sphere{Position: position, Radius: radius, Material: material}
}

// Creates a new cube at the given position with the given cubed-size and material
func NewCube(position mgl64.Vec3, size float64, material Material) *Cube {
	return &Cube{Position: position, Size: size, Material: material}
}

// Creates a new light at the given position with the given direction and brightness
func NewLight(position, direction mgl64.Vec3, brightness float64) Light {
	return Light{Position: position, Direction: direction, Brightness: brightness}
}

// Creates a new RGBA color with the given channel values
func NewColor(r, g, b uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: 255 }
}

// Creates a new RGBA color with the given channel values
func ConvertToRGBA(vec mgl64.Vec4) color.RGBA {
	// Clamps a floating point representation of an RGBA color into an RGBA struct
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
		A: clamp(vec[3]),
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

// Calculates whether the sphere intersects with the given ray, and computes the hit and normal
func (sphere *Sphere) Intersects(ray Ray) (intersects bool, hit, normal mgl64.Vec3) {
	panic("Not yet implemented")
}

// Calculates whether the cube intersects with the given ray, and computes the hit and normal
func (cube *Cube) Intersects(ray Ray) (intersects bool, hit, normal mgl64.Vec3) {
	panic("Not yet implemented")
}

// Retrieves the sphere's material
func (sphere *Sphere) GetMaterial() *Material {
	return &sphere.Material
}

// Retrieves the cube's material
func (cube *Cube) GetMaterial() *Material {
	return &cube.Material
}
