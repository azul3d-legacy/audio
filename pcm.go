// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"math"
)

type (
	// PCM8 represents an unsigned 8-bit linear PCM audio sample.
	PCM8 uint8

	// PCM8Samples represents a slice of PCM8 encoded audio samples.
	PCM8Samples []PCM8
)

// PCM8ToFloat64 converts a PCM8 encoded audio sample to Float64.
func PCM8ToFloat64(s PCM8) float64 {
	// In 0 to 1 range
	f := float64(s) / float64(math.MaxUint8)

	// Switch to -1 to +1 range
	f *= 2
	f -= 1
	return f
}

// Float64ToPCM8 converts a Float64 encoded audio sample to PCM8.
func Float64ToPCM8(s float64) PCM8 {
	// In -1 to +1 range, switch to 0 to 1
	s += 1
	s /= 2
	return PCM8(math.Floor(float64((s * float64(math.MaxUint8)) + 0.5)))
}

// Implements Slice interface.
func (p PCM8Samples) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p PCM8Samples) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p PCM8Samples) At(i int) float64 {
	return PCM8ToFloat64(p[i])
}

// Implements Slice interface.
func (p PCM8Samples) Set(i int, s float64) {
	p[i] = Float64ToPCM8(s)
}

// Implements Slice interface.
func (p PCM8Samples) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p PCM8Samples) Make(length, capacity int) Slice {
	return make(PCM8Samples, length, capacity)
}

// Implements Slice interface.
func (p PCM8Samples) CopyTo(dst Slice) int {
	d, ok := dst.(PCM8Samples)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}

type (
	// PCM16 represents a signed 16-bit linear PCM audio sample.
	PCM16 int16

	// PCM16Samples represents a slice of PCM16 encoded audio samples.
	PCM16Samples []PCM16
)

// PCM16ToFloat64 converts a PCM16 encoded audio sample to Float64.
func PCM16ToFloat64(s PCM16) float64 {
	return float64(s) / float64(math.MaxInt16)
}

// Float64ToPCM16 converts a Float64 encoded audio sample to PCM16.
func Float64ToPCM16(s float64) PCM16 {
	return PCM16(math.Floor(float64((s * float64(math.MaxInt16)) + 0.5)))
}

// Implements Slice interface.
func (p PCM16Samples) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p PCM16Samples) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p PCM16Samples) At(i int) float64 {
	return PCM16ToFloat64(p[i])
}

// Implements Slice interface.
func (p PCM16Samples) Set(i int, s float64) {
	p[i] = Float64ToPCM16(s)
}

// Implements Slice interface.
func (p PCM16Samples) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p PCM16Samples) Make(length, capacity int) Slice {
	return make(PCM16Samples, length, capacity)
}

// Implements Slice interface.
func (p PCM16Samples) CopyTo(dst Slice) int {
	d, ok := dst.(PCM16Samples)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}

type (
	// PCM32 represents a signed 32-bit linear PCM audio sample.
	PCM32 int32

	// PCM32Samples represents a slice of PCM32 encoded audio samples.
	PCM32Samples []PCM32
)

// PCM32ToFloat64 converts a PCM32 encoded audio sample to Float64.
func PCM32ToFloat64(s PCM32) float64 {
	return float64(s) / float64(math.MaxInt32)
}

// Float64ToPCM32 converts a Float64 encoded audio sample to PCM32.
func Float64ToPCM32(s float64) PCM32 {
	return PCM32(math.Floor(float64((s * float64(math.MaxInt32)) + 0.5)))
}

// Implements Slice interface.
func (p PCM32Samples) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p PCM32Samples) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p PCM32Samples) At(i int) float64 {
	return PCM32ToFloat64(p[i])
}

// Implements Slice interface.
func (p PCM32Samples) Set(i int, s float64) {
	p[i] = Float64ToPCM32(s)
}

// Implements Slice interface.
func (p PCM32Samples) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p PCM32Samples) Make(length, capacity int) Slice {
	return make(PCM32Samples, length, capacity)
}

// Implements Slice interface.
func (p PCM32Samples) CopyTo(dst Slice) int {
	d, ok := dst.(PCM32Samples)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}
