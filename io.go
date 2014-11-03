// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"errors"
)

// EOS is the error returned by Read when no more input is available. Functions
// should return EOS only to signal a graceful end of input. If the EOS occurs
// unexpectedly in a structured data stream, the appropriate error is either
// ErrUnexpectedEOS or some other error giving more detail.
var EOS = errors.New("end of stream")

// ErrUnexpectedEOS means that EOS was encountered in the middle of reading a
// fixed-size block or data structure.
var ErrUnexpectedEOS = errors.New("unexpected end of stream")

// ErrShortWrite means that a write accepted fewer bytes than requested but
// failed to return an explicit error.
var ErrShortWrite = errors.New("short write")

// Reader is a generic interface which describes any type who can have audio
// samples read from it into an audio slice.
type Reader interface {
	// Read tries to read into the audio slice, b, filling it with at max
	// b.Len() audio samples.
	//
	// Returned is the number of samples that where read into the slice, and
	// an error if any occurred.
	//
	// It is possible for the number of samples read to be non-zero; and for an
	// error to be returned at the same time (E.g. read 300 audio samples, but
	// also encountered EOS).
	Read(b Slice) (read int, e error)
}

// ReadSeeker is the generic seekable audio reader interface.
type ReadSeeker interface {
	Reader

	// Seek seeks to the specified sample number, relative to the start of the
	// stream. As such, subsequent Read() calls on the Reader, begin reading at
	// the specified sample.
	//
	// If any error is returned, it means it was impossible to seek to the
	// specified audio sample for some reason, and that the current playhead is
	// unchanged.
	Seek(sample uint64) error
}

// Writer is a generic interface which describes any type who can have audio
// samples written from an audio slice into it.
type Writer interface {
	// Write attempts to write all, b.Len(), samples in the slice to the
	// writer.
	//
	// Returned is the number of samples from the slice that where wrote to
	// the writer, and an error if any occured.
	//
	// If the number of samples wrote is less than buf.Len() then the returned
	// error must be non-nil. If any error occurs it should be considered fatal
	// with regards to the writer: no more data can be subsequently wrote after
	// an error.
	Write(b Slice) (wrote int, err error)
}

// WriterTo is the interface that wraps the WriteTo method.
//
// WriteTo writes data to w until there's no more data to write or when an
// error occurs. The return value n is the number of samples written. Any error
// encountered during the write is also returned.
//
// The Copy function uses WriterTo if available.
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}

// ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from r until EOS or error. The return value n is the
// number of bytes read. Any error except EOS encountered during the read is
// also returned.
//
// The Copy function uses ReaderFrom if available.
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

// Copy copies from src to dst until either EOS is reached on src or an
// error occurs.  It returns the number of samples copied and the first error
// encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOS. Because Copy is
// defined to read from src until EOS, it does not treat an EOS from Read as
// an error to be reported.
//
// If src implements the WriterTo interface, the copy is implemented by calling
// src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom interface, the
// copy is implemented by calling dst.ReadFrom(src).
func Copy(dst Writer, src Reader) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy. Avoids an
	// allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	buf := make(Float64, (32*1024)/8)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er == EOS {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}
