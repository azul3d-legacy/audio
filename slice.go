// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

// Slice is a generic audio slice, it can conceptually be thought of as a
// slice of some audio encoding type.
//
// Conversion between two encoded audio slices is as simple as:
//  dst, ok := src.(MuLaw)
//  if !ok {
//      // Create a new slice of the target encoding and copy the samples over
//      // because src is not MuLaw encoded.
//      dst = make(MuLaw, src.Len())
//      src.CopyTo(dst)
//  }
type Slice interface {
	// Len returns the number of elements in the slice.
	//
	// Equivalent slice syntax:
	//
	//  len(b)
	Len() int

	// Cap returns the number of elements in the slice.
	//
	// Equivalent slice syntax:
	//
	//  cap(b)
	Cap() int

	// Set sets the specified index in the slice to the specified Float64
	// encoded audio sample, s.
	//
	// If the slice's audio samples are not stored in Float64 encoding, then
	// the sample should be converted to the slice's internal format and then
	// stored.
	//
	// Just like slices, slice indices must be non-negative; and no greater
	// than (Len() - 1), or else a panic may occur.
	//
	// Equivalent slice syntax:
	//
	//  b[index] = s
	//   -> b.Set(index, s)
	//
	Set(index int, s float64)

	// At returns the Float64 encoded audio sample at the specified index in
	// the slice.
	//
	// If the slice's audio samples are not stored in Float64 encoding, then
	// the sample should be converted to Float64 encoding, and subsequently
	// returned.
	//
	// Just like slices, slice indices must be non-negative; and no greater
	// than (Len() - 1), or else a panic may occur.
	//
	// Equivalent slice syntax:
	//
	//  b[index]
	//   -> b.At(index)
	//
	At(index int) float64

	// Slice returns a new slice of the slice, using the low and high
	// parameters.
	//
	// Equivalent slice syntax:
	//
	//  b[low:high]
	//   -> b.Slice(low, high)
	//
	//  b[2:]
	//   -> b.Slice(2, a.Len())
	//
	//  b[:3]
	//   -> b.Slice(0, 3)
	//
	//  b[:]
	//   -> b.Slice(0, a.Len())
	//
	Slice(low, high int) Slice

	// Make creates and returns a new slice of this slices type. This allows
	// allocating a new slice of exactly the same type for lossless copying of
	// data without knowing about the underlying type.
	//
	// It is exactly the same syntax as the make builtin:
	//
	//  make(MuLaw, len, cap)
	//
	// Where cap cannot be less than len.
	Make(length, capacity int) Slice

	// CopyTo operates exactly like the copy builtin, but this slice is always
	// the source operand. Equivalent slice syntax:
	//
	//  copy(dst, src)
	//   -> src.CopyTo(dst)
	//
	CopyTo(dst Slice) int
}

// sliceCopy copies copies audio samples from the source slice to the
// destination slice. Returns the number of elements copied, which is the
// minimum of the dst.Len() and src.Len() values.
func sliceCopy(dst, src Slice) int {
	var i int
	for i = 0; i < src.Len() && i < dst.Len(); i++ {
		dst.Set(i, src.At(i))
	}
	return i
}
