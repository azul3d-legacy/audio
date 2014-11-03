// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"math"
)

// Int32 represents a signed 32-bit linear PCM encoded audio sample.
type Int32 []int32

// Int32ToFloat64 converts a Int32 encoded audio sample to Float64.
func Int32ToFloat64(s int32) float64 {
	return float64(s) / float64(math.MaxInt32)
}

// Float64ToInt32 converts a Float64 encoded audio sample to Int32.
func Float64ToInt32(s float64) int32 {
	return int32(math.Floor(float64((s * float64(math.MaxInt32)) + 0.5)))
}

// Len implements the Slice interface.
func (p Int32) Len() int {
	return len(p)
}

// Cap implements the Slice interface.
func (p Int32) Cap() int {
	return cap(p)
}

// At implements the Slice interface.
func (p Int32) At(i int) float64 {
	return Int32ToFloat64(p[i])
}

// Set implements the Slice interface.
func (p Int32) Set(i int, s float64) {
	p[i] = Float64ToInt32(s)
}

// Slice implements the Slice interface.
func (p Int32) Slice(low, high int) Slice {
	return p[low:high]
}

// Make implements the Slice interface.
func (p Int32) Make(length, capacity int) Slice {
	return make(Int32, length, capacity)
}

// CopyTo implements the Slice interface.
func (p Int32) CopyTo(dst Slice) int {
	d, ok := dst.(Int32)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}
