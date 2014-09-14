// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"bufio"
	"errors"
	"io"
)

// ErrFormat specifies an error where the format of the audio data is unknown
// of by the registered formats of this package.
var ErrFormat = errors.New("audio: unknown format")

// A format holds an audio format's name, magic header and how to decode it.
type format struct {
	name, magic string
	newDecoder  func(r interface{}) (Decoder, error)
}

// Formats is the list of registered formats.
var formats []format

// RegisterFormat registers an image format for use by NewDecoder().
//
// Name is the name of the format, like "wav" or "ogg".
//
// Magic is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
//
// newDecoder is the function that returns either [Decoder, nil] or
// [nil, ErrInvalidData] upon being called where the returned decoder is used
// to decode the io.Reader or io.ReadSeeker's encoded audio data.
func RegisterFormat(name, magic string, newDecoder func(r interface{}) (Decoder, error)) {
	formats = append(formats, format{name, magic, newDecoder})
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match returns whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) format {
	for _, f := range formats {
		b, err := r.Peek(len(f.magic))
		if err == nil && match(f.magic, b) {
			return f
		}
	}
	return format{}
}

// NewDecoder returns a decoder which can be used to decode the encoded audio
// data stored in the io.Reader or io.ReadSeeker, 'r'.
//
// The string returned is the format name used during format registration.
//
// Format registration is typically done by the init method of the codec-
// specific package.
func NewDecoder(r interface{}) (Decoder, string, error) {
	var rr reader
	switch t := r.(type) {
	case io.Reader:
		rr = asReader(t)
	case io.ReadSeeker:
		rr = asReader(t)
	default:
		panic("NewDecoder(): Invalid reader type; must be io.Reader or io.ReadSeeker!")
	}
	f := sniff(rr)
	if f.newDecoder == nil {
		return nil, "", ErrFormat
	}
	decoder, err := f.newDecoder(rr)
	return decoder, f.name, err
}
