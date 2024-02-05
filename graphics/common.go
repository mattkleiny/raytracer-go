// Copyright 2017, the project authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package graphics

const ε = 1e-9          // Minimum floating point precision
const MaxTraceDepth = 3 // The maximum depth for trace recursion

// Commonly used vectors
var (
	Zero     = NewVec(0, 0, 0)
	UnitX    = NewVec(1, 0, 0)
	UnitY    = NewVec(0, 1, 0)
	UnitZ    = NewVec(0, 0, 1)
	Identity = NewVec(1, 1, 1)
)

// Commonly used colors
var (
	Black = NewColor(0, 0, 0)
	Red   = NewColor(1, 0, 0)
	Green = NewColor(0, 1, 0)
	Blue  = NewColor(0, 0, 1)
	White = NewColor(1, 1, 1)
)
