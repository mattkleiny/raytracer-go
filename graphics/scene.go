// Copyright 2017, the project authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package graphics

import (
	"math"
)

// Represents a scene configuration for the ray-tracer; provides objects and materials which compose
// the scene structure and layout
type Scene struct {
	Camera          Camera   // The scene camera
	Lights          []Light  // The world light sources
	Objects         []Object // The objects composing the scene itself
	BackgroundColor Color    // The background color of the scene
}

// Represents a camera within the scene
type Camera struct {
	FieldOfView float64 // The field of view, in degrees
}

// Represents a light in the scene
type Light struct {
	Position Vector // The position of the light
	Emissive Color  // The emissive color of the light
}

// Describes the material surface of an object within a scene
type Material struct {
	Diffuse      Color   // The diffuse color of the material
	Transparency float64 // The transparency level of the material (between 0.0 and 1.0 inclusive)
	Reflectivity float64 // The reflectivity level of the material (between 0.0 and 1.0 inclusive)
}

// Represents an object in the scene
type Object interface {
	// Calculates whether the object Intersects the given ray and returns the distance.
	Intersects(ray Ray, distance *float64) bool
	GetMaterial() Material // The material of the object
	GetPosition() Vector   // The center position of the object
}

// Represents a sphere that can be placed in a scene
type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
}

// Calculates whether the sphere Intersects with the given ray and returns the hit distance.
func (sphere Sphere) Intersects(ray Ray, distance *float64) bool {
	to := ray.Origin.Sub(sphere.Center)

	b := to.Dot(ray.Direction)
	c := to.Dot(to) - sphere.Radius*sphere.Radius

	d := b*b - c

	if d > 0 {
		d = math.Sqrt(d)

		t1 := -b - d
		if t1 > ε {
			*distance = t1
			return true
		}

		t2 := -b + d
		if t2 > ε {
			*distance = t2
			return true
		}
	}

	return false
}

func (sphere Sphere) GetMaterial() Material {
	return sphere.Material
}

func (sphere Sphere) GetPosition() Vector {
	return sphere.Center
}

func (sphere Sphere) GetRadius() float64 {
	return sphere.Radius
}
