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

// Implements Slice interface.
func (p Int16) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p Int16) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p Int16) At(i int) float64 {
	return Int16ToFloat64(p[i])
}

// Implements Slice interface.
func (p Int16) Set(i int, s float64) {
	p[i] = Float64ToInt16(s)
}

// Implements Slice interface.
func (p Int16) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p Int16) Make(length, capacity int) Slice {
	return make(Int16, length, capacity)
}

// Implements Slice interface.
func (p Int16) CopyTo(dst Slice) int {
	d, ok := dst.(Int16)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}

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

// Implements Slice interface.
func (p Int32) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p Int32) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p Int32) At(i int) float64 {
	return Int32ToFloat64(p[i])
}

// Implements Slice interface.
func (p Int32) Set(i int, s float64) {
	p[i] = Float64ToInt32(s)
}

// Implements Slice interface.
func (p Int32) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p Int32) Make(length, capacity int) Slice {
	return make(Int32, length, capacity)
}

// Implements Slice interface.
func (p Int32) CopyTo(dst Slice) int {
	d, ok := dst.(Int32)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}
