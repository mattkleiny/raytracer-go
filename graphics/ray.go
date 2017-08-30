// Copyright 2017, the project authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package graphics

import "math"

// Represents a ray projected in 3d space
type Ray struct {
	Origin, Direction Vector
}

// Shorthand to create a new ray with the given components
func R(origin, direction Vector) Ray {
	return Ray{origin, direction.Normalize()}
}

// Calculates a point along the ray with the given distance
func (r Ray) Position(distance float64) Vector {
	return r.Origin.Add(r.Direction.MulS(distance))
}

// Reflects the ray about the given point and normal
func (r Ray) Reflect(point, normal Vector) Ray {
	origin := point.Add(normal.MulS(ε))
	direction := r.Direction.Sub(normal.MulS(2).MulS(r.Direction.Dot(normal)))

	return R(origin, direction)
}

// Refracts the ray about the given point and normal
func (r Ray) Refract(point, normal Vector, inside bool) Ray {
	ior := 1.1
	eta := ior

	if !inside {
		eta = 1 / ior
	}

	cosi := -normal.Dot(r.Direction)
	k := 1 - eta*eta*(1-cosi*cosi)

	origin := point.Sub(normal.MulS(ε))
	direction := r.Direction.MulS(eta).Add(normal.MulS(eta*cosi - math.Sqrt(k)))

	return R(origin, direction)
}
