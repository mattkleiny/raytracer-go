package graphics

import "math"

// Minimum floating point precision
const ε = 1e-9

// Represents a vector in 3d space
type Vector struct {
	X, Y, Z float64
}

// Shorthand to create a new vector with the given components
func V(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Cross(b Vector) Vector {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return V(x, y, z)
}

func (a Vector) Normalize() Vector {
	d := a.Length()
	return V(a.X/d, a.Y/d, a.Z/d)
}

func (a Vector) Negate() Vector {
	return V(-a.X, -a.Y, -a.Z)
}

func (a Vector) Abs() Vector {
	return V(math.Abs(a.X), math.Abs(a.Y), math.Abs(a.Z))
}

func (a Vector) Add(b Vector) Vector {
	return V(a.X+b.X, a.Y+b.Y, a.Z+b.Z)
}

func (a Vector) Sub(b Vector) Vector {
	return V(a.X-b.X, a.Y-b.Y, a.Z-b.Z)
}

func (a Vector) Mul(b Vector) Vector {
	return V(a.X*b.X, a.Y*b.Y, a.Z*b.Z)
}

func (a Vector) Div(b Vector) Vector {
	return V(a.X/b.X, a.Y/b.Y, a.Z/b.Z)
}

func (a Vector) Mod(b Vector) Vector {
	x := a.X - b.X*math.Floor(a.X/b.X)
	y := a.Y - b.Y*math.Floor(a.Y/b.Y)
	z := a.Z - b.Z*math.Floor(a.Z/b.Z)

	return V(x, y, z)
}

func (a Vector) AddScalar(b float64) Vector {
	return V(a.X+b, a.Y+b, a.Z+b)
}

func (a Vector) SubScalar(b float64) Vector {
	return V(a.X-b, a.Y-b, a.Z-b)
}

func (a Vector) MulScalar(b float64) Vector {
	return V(a.X*b, a.Y*b, a.Z*b)
}

func (a Vector) DivScalar(b float64) Vector {
	return V(a.X/b, a.Y/b, a.Z/b)
}

// Represents a ray projected in 3d space
type Ray struct {
	Origin, Direction Vector
}

// Shorthand to create a new ray with the given components
func R(origin, direction Vector) Ray {
	return Ray{origin, direction.Normalize()}
}

func (r Ray) Position(t float64) Vector {
	return r.Origin.Add(r.Direction.MulScalar(t))
}
