// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

// Float64 represents a slice of 64-bit floating-point linear PCM encoded
// audio samples, in the range of -1 to +1.
type Float64 []float64

// Implements Slice interface.
func (p Float64) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p Float64) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p Float64) At(i int) float64 {
	return p[i]
}

// Implements Slice interface.
func (p Float64) Set(i int, s float64) {
	p[i] = s
}

// Implements Slice interface.
func (p Float64) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p Float64) Make(length, capacity int) Slice {
	return make(Float64, length, capacity)
}

// Implements Slice interface.
func (p Float64) CopyTo(dst Slice) int {
	d, ok := dst.(Float64)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}
