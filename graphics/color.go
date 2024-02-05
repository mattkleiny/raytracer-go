// Copyright 2017, the project authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package graphics

import (
	"image/color"
	"math"
)

// Represents a color composed of red, green and blue
type Color struct {
	R, G, B float64
}

// Shorthand to create a new color with the given components
func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

// Adds the color piece-wise to the given color
func (a Color) Add(b Color) Color {
	return NewColor(a.R+b.R, a.G+b.G, a.B+b.B)
}

// Subtracts the color piece-wise from the given color
func (a Color) Sub(b Color) Color {
	return NewColor(a.R-b.R, a.G-b.G, a.B-b.B)
}

// Multiplies the color piece-wise by the given color
func (a Color) Mul(b Color) Color {
	return NewColor(a.R*b.R, a.G*b.G, a.B*b.B)
}

// Divides the color piece-wise by the given color
func (a Color) Div(b Color) Color {
	return NewColor(a.R/b.R, a.G/b.G, a.B/b.B)
}

// Adds the given scalar amount to the color
func (a Color) AddS(b float64) Color {
	return NewColor(a.R+b, a.G+b, a.B+b)
}

// Subtracts the given scalar amount from the color
func (a Color) SubS(b float64) Color {
	return NewColor(a.R-b, a.G-b, a.B-b)
}

// Multiplies the color by the given scalar amount
func (a Color) MulS(b float64) Color {
	return NewColor(a.R*b, a.G*b, a.B*b)
}

// Divides the color by the given scalar amount
func (a Color) DivS(b float64) Color {
	return NewColor(a.R/b, a.G/b, a.B/b)
}

// Converts a color in floating point range (0.0 to 1.0) to a color with the given channel values
func (c Color) ConvertToRGBA() color.RGBA {
	// Clamps a floating point (0..1) into a (0..255) uint8
	clamp := func(value float64) uint8 {
		if value == 1.0 {
			return 255
		}
		return uint8(math.Floor(value * 256.0))
	}

	return color.RGBA{
		R: clamp(c.R),
		G: clamp(c.G),
		B: clamp(c.B),
		A: 255,
	}
}
