package graphics

import (
	"fmt"
	"math"
)

// Precision for our floating-point comparisons
const Epsilon float64 = 0.00001

type Ray struct{ Vector }
type Vector struct{ X, Y, Z float64 }

var (
	Zero     = Vector{0, 0, 0}
	Identity = Vector{1, 1, 1}
	UnitX    = Vector{1, 0, 0}
	UnitY    = Vector{0, 1, 0}
	UnitZ    = Vector{0, 0, 1}
)

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func (v Vector) Clone() Vector {
	return Vector{v.X, v.Y, v.Z }
}

func (v Vector) Add(o Vector) Vector {
	return NewVector(v.X+o.X, v.Y+o.Y, v.Z+o.Z)
}

func (v Vector) Subtract(o Vector) Vector {
	return NewVector(v.X-o.X, v.Y-o.Y, v.Z-o.Z)
}

func (v Vector) Multiply(s float64) Vector {
	return NewVector(v.X*s, v.Y*s, v.Z*s)
}

func (v Vector) Divide(s float64) Vector {
	return NewVector(v.X/s, v.Y/s, v.Z/s)
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.MagnitudeSqr())
}

func (v Vector) MagnitudeSqr() float64 {
	return v.X*v.X + v.Y*v.Y*v.Z*v.Z
}

func (v Vector) Distance(o Vector) float64 {
	return math.Sqrt(v.DistanceSqr(o))
}

func (v Vector) DistanceSqr(o Vector) float64 {
	x := o.X - v.X
	y := o.Y - v.Y
	z := o.Z - v.Z

	return x*x + y*y + z*z
}

func (v Vector) Normalize() Vector {
	return v.Divide(v.Magnitude())
}

func (v Vector) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.X, v.Y, v.Z)
}
