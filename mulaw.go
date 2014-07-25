// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"math"
)

type (
	// MuLaw represents an uint8 mulaw encoded audio sample.
	MuLaw uint8

	// MuLawSamples represents an slice of MuLaw encoded audio samples.
	MuLawSamples []MuLaw
)

// Implements Slice interface.
func (p MuLawSamples) Len() int {
	return len(p)
}

// Implements Slice interface.
func (p MuLawSamples) Cap() int {
	return cap(p)
}

// Implements Slice interface.
func (p MuLawSamples) At(i int) F64 {
	p16 := MuLawToPCM16(p[i])
	return F64(p16) / F64(math.MaxInt16)
}

// Implements Slice interface.
func (p MuLawSamples) Set(i int, s F64) {
	p16 := F64ToPCM16(s)
	p[i] = PCM16ToMuLaw(p16)
}

// Implements Slice interface.
func (p MuLawSamples) Slice(low, high int) Slice {
	return p[low:high]
}

// Implements Slice interface.
func (p MuLawSamples) Make(length, capacity int) Slice {
	return make(MuLawSamples, length, capacity)
}

// Implements Slice interface.
func (p MuLawSamples) CopyTo(dst Slice) int {
	d, ok := dst.(MuLawSamples)
	if ok {
		return copy(d, p)
	}
	return sliceCopy(dst, p)
}

const (
	muLawCBias = 0x84
	muLawCClip = 32635
)

var (
	// For explanation see: http://www.threejacks.com/?q=node/176

	muLawCompressTable = [256]PCM8{
		0, 0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
		4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	}

	muLawDecompressTable = [256]PCM16{
		-32124, -31100, -30076, -29052, -28028, -27004, -25980, -24956,
		-23932, -22908, -21884, -20860, -19836, -18812, -17788, -16764,
		-15996, -15484, -14972, -14460, -13948, -13436, -12924, -12412,
		-11900, -11388, -10876, -10364, -9852, -9340, -8828, -8316,
		-7932, -7676, -7420, -7164, -6908, -6652, -6396, -6140,
		-5884, -5628, -5372, -5116, -4860, -4604, -4348, -4092,
		-3900, -3772, -3644, -3516, -3388, -3260, -3132, -3004,
		-2876, -2748, -2620, -2492, -2364, -2236, -2108, -1980,
		-1884, -1820, -1756, -1692, -1628, -1564, -1500, -1436,
		-1372, -1308, -1244, -1180, -1116, -1052, -988, -924,
		-876, -844, -812, -780, -748, -716, -684, -652,
		-620, -588, -556, -524, -492, -460, -428, -396,
		-372, -356, -340, -324, -308, -292, -276, -260,
		-244, -228, -212, -196, -180, -164, -148, -132,
		-120, -112, -104, -96, -88, -80, -72, -64,
		-56, -48, -40, -32, -24, -16, -8, -1,
		32124, 31100, 30076, 29052, 28028, 27004, 25980, 24956,
		23932, 22908, 21884, 20860, 19836, 18812, 17788, 16764,
		15996, 15484, 14972, 14460, 13948, 13436, 12924, 12412,
		11900, 11388, 10876, 10364, 9852, 9340, 8828, 8316,
		7932, 7676, 7420, 7164, 6908, 6652, 6396, 6140,
		5884, 5628, 5372, 5116, 4860, 4604, 4348, 4092,
		3900, 3772, 3644, 3516, 3388, 3260, 3132, 3004,
		2876, 2748, 2620, 2492, 2364, 2236, 2108, 1980,
		1884, 1820, 1756, 1692, 1628, 1564, 1500, 1436,
		1372, 1308, 1244, 1180, 1116, 1052, 988, 924,
		876, 844, 812, 780, 748, 716, 684, 652,
		620, 588, 556, 524, 492, 460, 428, 396,
		372, 356, 340, 324, 308, 292, 276, 260,
		244, 228, 212, 196, 180, 164, 148, 132,
		120, 112, 104, 96, 88, 80, 72, 64,
		56, 48, 40, 32, 24, 16, 8, 0,
	}
)

// PCM16ToMuLaw converts from a PCM16 encoded audio sample to an MuLaw encoded
// audio sample.
func PCM16ToMuLaw(s PCM16) MuLaw {
	sign := (s >> 8) & 0x80
	if sign != 0 {
		s = -s
	}
	if s > muLawCClip {
		s = muLawCClip
	}
	s = s + muLawCBias
	exponent := muLawCompressTable[(s>>7)&0xFF]
	mantissa := (s >> (exponent + 3)) & 0x0F
	compressedByte := ^(sign | (PCM16(exponent) << 4) | mantissa)
	return MuLaw(compressedByte)
}

// MuLawToPCM16 converts from an MuLaw encoded audio sample to an PCM16 encoded
// audio sample.
func MuLawToPCM16(s MuLaw) PCM16 {
	return muLawDecompressTable[s]
}
