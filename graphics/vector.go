// Copyright 2017, the project authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package graphics

import "math"

// Represents a vector in 3d space
type Vector struct {
	X, Y, Z float64
}

// Shorthand to create a new vector with the given components
func NewVec(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// Calculates the length of this vector
func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Calculates the square length of this vector
func (a Vector) LengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

// Calculates the dot product of this vector and the given vector
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Calculates the cross product of this vector and the given vector
func (a Vector) Cross(b Vector) Vector {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return NewVec(x, y, z)
}

// Normalizes the vector
func (a Vector) Normalize() Vector {
	d := a.Length()
	return NewVec(a.X/d, a.Y/d, a.Z/d)
}

// Negates the vector
func (a Vector) Negate() Vector {
	return NewVec(-a.X, -a.Y, -a.Z)
}

// Calculates the absolute value of the vector piece-wise
func (a Vector) Abs() Vector {
	return NewVec(math.Abs(a.X), math.Abs(a.Y), math.Abs(a.Z))
}

// Adds the vector piece-wise to the given vector
func (a Vector) Add(b Vector) Vector {
	return NewVec(a.X+b.X, a.Y+b.Y, a.Z+b.Z)
}

// Subtracts the vector piece-wise from the given vector
func (a Vector) Sub(b Vector) Vector {
	return NewVec(a.X-b.X, a.Y-b.Y, a.Z-b.Z)
}

// Multiplies the vector piece-wise by the given vector
func (a Vector) Mul(b Vector) Vector {
	return NewVec(a.X*b.X, a.Y*b.Y, a.Z*b.Z)
}

// Divides the vector piece-wise by the given vector
func (a Vector) Div(b Vector) Vector {
	return NewVec(a.X/b.X, a.Y/b.Y, a.Z/b.Z)
}

// Adds the given scalar amount to the vector
func (a Vector) AddS(b float64) Vector {
	return NewVec(a.X+b, a.Y+b, a.Z+b)
}

// Subtracts the given scalar amount from the vector
func (a Vector) SubS(b float64) Vector {
	return NewVec(a.X-b, a.Y-b, a.Z-b)
}

// Multiplies the vector by the given scalar amount
func (a Vector) MulS(b float64) Vector {
	return NewVec(a.X*b, a.Y*b, a.Z*b)
}

// Divides the vector by the given scalar amount
func (a Vector) DivS(b float64) Vector {
	return NewVec(a.X/b, a.Y/b, a.Z/b)
}
