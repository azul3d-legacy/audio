// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"math"
)

// Int16 represents a signed 16-bit linear PCM encoded audio sample.
type Int16 []int16

// Int16ToFloat64 converts a Int16 encoded audio sample to Float64.
func Int16ToFloat64(s int16) float64 {
	return float64(s) / float64(math.MaxInt16)
}

// Float64ToInt16 converts a Float64 encoded audio sample to Int16.
func Float64ToInt16(s float64) int16 {
	return int16(math.Floor(float64((s * float64(math.MaxInt16)) + 0.5)))
}

// Len implements the Slice interface.
func (p Int16) Len() int {
	return len(p)
}

// Cap implements the Slice interface.
func (p Int16) Cap() int {
	return cap(p)
}

// At implements the Slice interface.
func (p Int16) At(i int) float64 {
	return Int16ToFloat64(p[i])
}

// Set implements the Slice interface.
func (p Int16) Set(i int, s float64) {
	p[i] = Float64ToInt16(s)
}

// Slice implements the Slice interface.
func (p Int16) Slice(low, high int) Slice {
	return p[low:high]
}

// Make implements the Slice interface.
func (p Int16) Make(length, capacity int) Slice {
	return make(Int16, length, capacity)
}

// CopyTo implements the Slice interface.
func (p Int16) CopyTo(dst Slice) int {
	d, ok := dst.(Int16)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}
