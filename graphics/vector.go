package graphics

type Vector struct{ x, y, z float64 }

var (
	Zero     = Vector{0, 0, 0}
	Up       = Vector{0, 1, 0}
	Identity = Vector{1, 1, 1}
	UnitX    = Vector{1, 0, 0}
	UnitY    = Vector{0, 1, 0}
	UnitZ    = Vector{0, 0, 1}
)

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func (v Vector) Clone() Vector {
	return Vector{v.x, v.y, v.z }
}
