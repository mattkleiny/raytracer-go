package graphics

import (
	"testing"
	"math"
)

func TestVector_Clone(t *testing.T) {
	vector := NewVector(1, 2, 3)
	clone := vector.Clone()

	if clone.X != 1 || clone.Y != 2 || clone.Z != 3 {
		t.Fail()
	}
}

func TestVector_Add(t *testing.T) {
	vector := NewVector(0, 1, 0)
	result := vector.Add(UnitY)

	if result.X != 0 && result.Y != 2 && result.Z != 0 {
		t.Fail()
	}
}

func TestVector_Subtract(t *testing.T) {
	vector := NewVector(0, 1, 0)
	result := vector.Subtract(UnitY)

	if result.X != 0 && result.Y != 0 && result.Z != 0 {
		t.Fail()
	}
}

func TestVector_MultiplyScalar(t *testing.T) {
	vector := Identity
	result := vector.Multiply(2)

	if result.X != 2 && result.Y != 2 && result.Z != 2 {
		t.Fail()
	}
}

func TestVector_DivideScalar(t *testing.T) {
	vector := NewVector(2, 2, 2)
	result := vector.Divide(2)

	if result.X != 1 && result.Y != 1 && result.Z != 1 {
		t.Fail()
	}
}

func TestVector_Magnitude(t *testing.T) {
	vector := NewVector(2, 2, 2)
	magnitude := vector.Magnitude()

	if math.Abs(magnitude-4.4) <= Epsilon {
		t.Fail()
	}
}

func TestVector_Distance(t *testing.T) {
	vector := NewVector(2, 2, 2)
	distance := vector.Distance(Zero)

	if math.Abs(distance-4.4) <= Epsilon {
		t.Fail()
	}
}

func TestVector_Normalize(t *testing.T) {
	vector := NewVector(0, 2, 0)
	normalized := vector.Normalize()

	if math.Abs(normalized.Y-1) <= Epsilon {
		t.Fail()
	}
}
