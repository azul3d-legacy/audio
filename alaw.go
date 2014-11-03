// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"math"
)

// ALaw represents a slice of ALaw encoded audio samples.
type ALaw []uint8

// Len implements the Slice interface.
func (p ALaw) Len() int {
	return len(p)
}

// Cap implements the Slice interface.
func (p ALaw) Cap() int {
	return cap(p)
}

// At implements the Slice interface.
func (p ALaw) At(i int) float64 {
	p16 := ALawToInt16(p[i])
	return float64(p16) / float64(math.MaxInt16)
}

// Set implements the Slice interface.
func (p ALaw) Set(i int, s float64) {
	p16 := Float64ToInt16(s)
	p[i] = Int16ToALaw(p16)
}

// Slice implements the Slice interface.
func (p ALaw) Slice(low, high int) Slice {
	return p[low:high]
}

// Make implements the Slice interface.
func (p ALaw) Make(length, capacity int) Slice {
	return make(ALaw, length, capacity)
}

// CopyTo implements the Slice interface.
func (p ALaw) CopyTo(dst Slice) int {
	d, ok := dst.(ALaw)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}

const (
	alawCBias = 0x84
	alawCClip = 32635
)

var (
	// For explanation see: http://www.threejacks.com/?q=node/176

	aLawCompressTable = [256]uint8{
		1, 1, 2, 2, 3, 3, 3, 3,
		4, 4, 4, 4, 4, 4, 4, 4,
		5, 5, 5, 5, 5, 5, 5, 5,
		5, 5, 5, 5, 5, 5, 5, 5,
		6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6,
		7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7,
	}

	aLawDecompressTable = [256]int16{
		-5504, -5248, -6016, -5760, -4480, -4224, -4992, -4736,
		-7552, -7296, -8064, -7808, -6528, -6272, -7040, -6784,
		-2752, -2624, -3008, -2880, -2240, -2112, -2496, -2368,
		-3776, -3648, -4032, -3904, -3264, -3136, -3520, -3392,
		-22016, -20992, -24064, -23040, -17920, -16896, -19968, -18944,
		-30208, -29184, -32256, -31232, -26112, -25088, -28160, -27136,
		-11008, -10496, -12032, -11520, -8960, -8448, -9984, -9472,
		-15104, -14592, -16128, -15616, -13056, -12544, -14080, -13568,
		-344, -328, -376, -360, -280, -264, -312, -296,
		-472, -456, -504, -488, -408, -392, -440, -424,
		-88, -72, -120, -104, -24, -8, -56, -40,
		-216, -200, -248, -232, -152, -136, -184, -168,
		-1376, -1312, -1504, -1440, -1120, -1056, -1248, -1184,
		-1888, -1824, -2016, -1952, -1632, -1568, -1760, -1696,
		-688, -656, -752, -720, -560, -528, -624, -592,
		-944, -912, -1008, -976, -816, -784, -880, -848,
		5504, 5248, 6016, 5760, 4480, 4224, 4992, 4736,
		7552, 7296, 8064, 7808, 6528, 6272, 7040, 6784,
		2752, 2624, 3008, 2880, 2240, 2112, 2496, 2368,
		3776, 3648, 4032, 3904, 3264, 3136, 3520, 3392,
		22016, 20992, 24064, 23040, 17920, 16896, 19968, 18944,
		30208, 29184, 32256, 31232, 26112, 25088, 28160, 27136,
		11008, 10496, 12032, 11520, 8960, 8448, 9984, 9472,
		15104, 14592, 16128, 15616, 13056, 12544, 14080, 13568,
		344, 328, 376, 360, 280, 264, 312, 296,
		472, 456, 504, 488, 408, 392, 440, 424,
		88, 72, 120, 104, 24, 8, 56, 40,
		216, 200, 248, 232, 152, 136, 184, 168,
		1376, 1312, 1504, 1440, 1120, 1056, 1248, 1184,
		1888, 1824, 2016, 1952, 1632, 1568, 1760, 1696,
		688, 656, 752, 720, 560, 528, 624, 592,
		944, 912, 1008, 976, 816, 784, 880, 848,
	}
)

// Int16ToALaw converts an Int16 encoded audio sample to an ALaw encoded audio
// sample.
func Int16ToALaw(s int16) uint8 {
	sign := ((^s) >> 8) & 0x80
	if sign == 0 {
		s = -s
	}
	if s > alawCClip {
		s = alawCClip
	}
	var compressedByte uint8
	if s >= 256 {
		exponent := aLawCompressTable[(s>>8)&0x7F]
		mantissa := (s >> (exponent + 3)) & 0x0F
		compressedByte = uint8(((int16(exponent) << 4) | mantissa))
	} else {
		compressedByte = uint8(s >> 4)
	}
	compressedByte ^= uint8(sign ^ 0x55)
	return compressedByte
}

// ALawToInt16 converts an ALaw encoded audio sample to an Int16 encoded audio
// sample.
func ALawToInt16(s uint8) int16 {
	return aLawDecompressTable[s]
}
