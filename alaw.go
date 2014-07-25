// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"math"
)

type (
	// ALaw represents an uint8 alaw encoded audio sample.
	ALaw uint8

	// ALawSamples represents an slice of ALaw encoded audio samples.
	ALawSamples []ALaw
)

// Implements Slice interface.
func (p ALawSamples) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p ALawSamples) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p ALawSamples) At(i int) F64 {
	p16 := ALawToPCM16(p[i])
	return F64(p16) / F64(math.MaxInt16)
}

// Implements Slice interface.
func (p ALawSamples) Set(i int, s F64) {
	p16 := F64ToPCM16(s)
	p[i] = PCM16ToALaw(p16)
}

// Implements Slice interface.
func (p ALawSamples) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p ALawSamples) Make(length, capacity int) Slice {
	return make(ALawSamples, length, capacity)
}

// Implements Slice interface.
func (p ALawSamples) CopyTo(dst Slice) int {
	d, ok := dst.(ALawSamples)
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

	aLawCompressTable = [256]PCM8{
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

	aLawDecompressTable = [256]PCM16{
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

// PCM16ToALaw converts an PCM16 encoded audio sample to an ALaw encoded audio
// sample.
func PCM16ToALaw(s PCM16) ALaw {
	sign := ((^s) >> 8) & 0x80
	if sign == 0 {
		s = -s
	}
	if s > alawCClip {
		s = alawCClip
	}
	var compressedByte ALaw
	if s >= 256 {
		exponent := aLawCompressTable[(s>>8)&0x7F]
		mantissa := (s >> (exponent + 3)) & 0x0F
		compressedByte = ALaw(((PCM16(exponent) << 4) | mantissa))
	} else {
		compressedByte = ALaw(s >> 4)
	}
	compressedByte ^= ALaw(sign ^ 0x55)
	return compressedByte
}

// ALawToPCM16 converts an ALaw encoded audio sample to an PCM16 encoded audio
// sample.
func ALawToPCM16(s ALaw) PCM16 {
	return aLawDecompressTable[s]
}
