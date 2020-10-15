package vector

import (
	"math"
	"testing"
)

const tolerance = 1e-15

func eq(a float64, b float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestNewVector(t *testing.T) {
	v := NewVector(9)
	if len(v) != 9 {
		t.Errorf("len(NewVector(9)) = %d", len(v))
	}
}

func TestAdd(t *testing.T) {
	v1 := Vector([]float64{1, 2})
	v2 := Vector([]float64{3, 4})

	v := v1.Add(v2)
	if !eq(v[0], 4) || !eq(v[1], 6) {
		t.Errorf("%v + %v = %v", v1, v2, v)
	}
}

func TestSub(t *testing.T) {
	v1 := Vector([]float64{1, 2})
	v2 := Vector([]float64{3, 4})

	v := v1.Sub(v2)
	if !eq(v[0], -2) || !eq(v[1], -2) {
		t.Errorf("%v - %v = %v", v1, v2, v)
	}
}

func TestMul(t *testing.T) {
	v1 := Vector([]float64{1, 2})

	v := v1.Mul(2)
	if !eq(v[0], 2) || !eq(v[1], 4) {
		t.Errorf("%v * %v = %v", v1, 2, v)
	}
}

func TestDot(t *testing.T) {
	v1 := Vector([]float64{1, 2})
	v2 := Vector([]float64{3, 4})

	v := v1.Dot(v2)
	if !eq(v, 11) {
		t.Errorf("(%v, %v) = %v", v1, v2, v)
	}
}

func TestCross(t *testing.T) {
	v1 := Vector([]float64{1, 2, 3})
	v2 := Vector([]float64{3, 4, 5})

	v := v1.Cross(v2)
	if !eq(v[0], -2) || !eq(v[1], 4) || !eq(v[2], -2) {
		t.Errorf("%v x %v = %v", v1, v2, v)
	}
}

func TestNorm(t *testing.T) {
	v1 := Vector([]float64{3, 4})

	v := v1.Norm()
	if !eq(v, 5) {
		t.Errorf("||%v|| = %v", v1, v)
	}
}

func TestChain(t *testing.T) {
	v1 := Vector([]float64{1, 2})
	v2 := Vector([]float64{3, 4})
	v3 := Vector([]float64{5, 6})

	v := v1.Add(v2).Dot(v3)

	if !eq(v, 56) {
		t.Errorf("(%v + %v, %v) = %v", v1, v2, v3, v)
	}
}
