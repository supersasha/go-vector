// Package vector provides basic mathematical vector operations
// NOTE: The following decisions were made:
// 1. []float64 was chosen as the type of vector since it:
//	  - gives some universality (compared to array [N]float64).
//		However it might make sense to limit it to, say, [3]float64
//		if we know that we are limited to 3-dimensional space.
//	  - float64 is most popular floating point type
// 2. Operations are implemented as methods to make it possible
//	  chaining like v1.Add(v2).Mul(5.0).Sub(v3).Cross(v4).Dot(v5)
// 3. For the purpose of chaining and assuming different-length
//	  vectors for binary operations as a serious bug, operations
//	  don't return errors. They still check lengths and panic
//	  since dimension errors is more informative compared to
//	  index range errors.
// 4. All operations return new vectors instead of modifying them in-place.
//	  This is made for these reasons:
//	  - immutability gives more control and causes less errors
//	  -	chaining
// 5. Assumed right-hand orthonormal basis for the cross product
package vector

import (
	"math"
)

type Vector []float64

// NewVector creates a vector
func NewVector(n int) Vector {
	return Vector(make([]float64, n))
}

// Add returns sum of two vectors
func (v1 Vector) Add(v2 Vector) Vector {
	if len(v1) != len(v2) {
		panic("dimensions don't match")
	}
	v := NewVector(len(v1))
	for i := 0; i < len(v1); i++ {
		v[i] = v1[i] + v2[i]
	}
	return v
}

// Add returns difference of two vectors
func (v1 Vector) Sub(v2 Vector) Vector {
	if len(v1) != len(v2) {
		panic("dimensions don't match")
	}
	v := NewVector(len(v1))
	for i := 0; i < len(v1); i++ {
		v[i] = v1[i] - v2[i]
	}
	return v
}

// Mul multilies vector by a scalar
func (v1 Vector) Mul(f float64) Vector {
	v := NewVector(len(v1))
	for i := 0; i < len(v1); i++ {
		v[i] = v1[i] * f
	}
	return v
}

// Dot returns dot product of two vectors
func (v1 Vector) Dot(v2 Vector) float64 {
	if len(v1) != len(v2) {
		panic("dimensions don't match")
	}

	var s float64
	for i := 0; i < len(v1); i++ {
		s += v1[i] * v2[i]
	}
	return s
}

// Cross returns cross product of two vectors in right-hand orthonormal basis
func (v1 Vector) Cross(v2 Vector) Vector {
	if len(v1) != 3 || len(v2) != 3 {
		panic("cross product: both vectors must be 3-dimentional")
	}
	v := Vector([]float64{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	})
	return v
}

// Norm return the norm(2) of a vector
func (v1 Vector) Norm() float64 {
	var s float64
	for i := 0; i < len(v1); i++ {
		s += v1[i] * v1[i]
	}
	return math.Sqrt(s)
}
