package graphics

import (
	"math"
	"fmt"
)

// Represents a floating-point vector in 3d space
type Vector struct {
	X, Y, Z float64
}

// Creates a new vector with the given components
func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// Clones the given vector piece-wise
func (v Vector) Clone() Vector {
	return Vector{v.X, v.Y, v.Z }
}

// Immutably adds the given vector to this vector
func (v Vector) Add(o Vector) Vector {
	return NewVector(v.X+o.X, v.Y+o.Y, v.Z+o.Z)
}

// Immutably subtracts the given vector from this vector
func (v Vector) Subtract(o Vector) Vector {
	return NewVector(v.X-o.X, v.Y-o.Y, v.Z-o.Z)
}

// Immutably multiplies the given scalar to this vector
func (v Vector) Multiply(s float64) Vector {
	return NewVector(v.X*s, v.Y*s, v.Z*s)
}

// Immutably divides the given scalar from this vector
func (v Vector) Divide(s float64) Vector {
	return NewVector(v.X/s, v.Y/s, v.Z/s)
}

// Computes the magnitude of the vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.MagnitudeSqr())
}

// Computes the square magnitude of the vector
func (v Vector) MagnitudeSqr() float64 {
	return v.X*v.X + v.Y*v.Y*v.Z*v.Z
}

// Computes the distance between this vector and the given vector
func (v Vector) Distance(o Vector) float64 {
	return math.Sqrt(v.DistanceSqr(o))
}

// Computes the square distance between this vector and the given vector
func (v Vector) DistanceSqr(o Vector) float64 {
	x := o.X - v.X
	y := o.Y - v.Y
	z := o.Z - v.Z

	return x*x + y*y + z*z
}

// Normalizes the vector
func (v Vector) Normalize() Vector {
	return v.Divide(v.Magnitude())
}

// String-ifies the vector
func (v Vector) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.X, v.Y, v.Z)
}

// Represents a ray projected into 3d space
type Ray struct {
	Origin    Vector
	Direction Vector
}

// Projects a ray into the scene from the given (x, y) coordinates
func ProjectRay(origin Vector, x, y int) Ray {
	panic("Not yet implemented")
}
