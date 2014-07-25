// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

type (
	// F32 represents an 32-bit floating-point linear audio sample in the range
	// of -1 to +1.
	F32 float32

	// F32Samples represents an slice of F32 encoded audio samples.
	F32Samples []F32
)

// Implements Slice interface.
func (p F32Samples) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p F32Samples) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p F32Samples) At(i int) F64 {
	return F64(p[i])
}

// Implements Slice interface.
func (p F32Samples) Set(i int, s F64) {
	p[i] = F32(s)
}

// Implements Slice interface.
func (p F32Samples) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p F32Samples) Make(length, capacity int) Slice {
	return make(F32Samples, length, capacity)
}

// Implements Slice interface.
func (p F32Samples) CopyTo(dst Slice) int {
	d, ok := dst.(F32Samples)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}

type (
	// F64 represents an 64-bit floating-point linear audio sample in the range
	// of -1 to +1.
	F64 float64

	// F32Samples represents an slice of F32 encoded audio samples.
	F64Samples []F64
)

// Implements Slice interface.
func (p F64Samples) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p F64Samples) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p F64Samples) At(i int) F64 {
	return p[i]
}

// Implements Slice interface.
func (p F64Samples) Set(i int, s F64) {
	p[i] = s
}

// Implements Slice interface.
func (p F64Samples) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p F64Samples) Make(length, capacity int) Slice {
	return make(F64Samples, length, capacity)
}

// Implements Slice interface.
func (p F64Samples) CopyTo(dst Slice) int {
	d, ok := dst.(F64Samples)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}
