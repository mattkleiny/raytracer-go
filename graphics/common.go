package graphics

const ε = 1e-9          // Minimum floating point precision
const MaxTraceDepth = 3 // The maximum depth for trace recursion

// Commonly used vectors
var (
	Zero     = V(0, 0, 0)
	UnitX    = V(1, 0, 0)
	UnitY    = V(0, 1, 0)
	UnitZ    = V(0, 0, 1)
	Identity = V(1, 1, 1)
)

// Commonly used colors
var (
	Black = C(0, 0, 0)
	Red   = C(1, 0, 0)
	Green = C(0, 1, 0)
	Blue  = C(0, 0, 1)
	White = C(1, 1, 1)
)
