package graphics

type Camera struct {
	position Vector
	facing   Vector
}

func NewCamera() *Camera {
	return &Camera{
		position: Zero,
		facing:   UnitZ,
	}
}