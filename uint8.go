// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"math"
)

// Uint8 represents a unsigned 8-bit linear PCM encoded audio sample.
type Uint8 []uint8

// Uint8ToFloat64 converts a Uint8 encoded audio sample to Float64.
func Uint8ToFloat64(s uint8) float64 {
	// In 0 to 1 range
	f := float64(s) / float64(math.MaxUint8)

	// Switch to -1 to +1 range
	f *= 2
	f -= 1
	return f
}

// Float64ToUint8 converts a Float64 encoded audio sample to Uint8.
func Float64ToUint8(s float64) uint8 {
	// In -1 to +1 range, switch to 0 to 1
	s += 1
	s /= 2
	return uint8(math.Floor(float64((s * float64(math.MaxUint8)) + 0.5)))
}

// Implements Slice interface.
func (p Uint8) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p Uint8) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p Uint8) At(i int) float64 {
	return Uint8ToFloat64(p[i])
}

// Implements Slice interface.
func (p Uint8) Set(i int, s float64) {
	p[i] = Float64ToUint8(s)
}

// Implements Slice interface.
func (p Uint8) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p Uint8) Make(length, capacity int) Slice {
	return make(Uint8, length, capacity)
}

// Implements Slice interface.
func (p Uint8) CopyTo(dst Slice) int {
	d, ok := dst.(Uint8)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}
