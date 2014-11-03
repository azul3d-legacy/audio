// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

// A Buffer is a variable-sized buffer of audio samples with Read and Write
// methods. Buffers must be allocated via the NewBuffer function.
type Buffer struct {
	buf Slice // contents are the samples buf[off : len(buf)]
	off int   // read at buf[off], write at buf[len(buf)]
}

// Samples returns a slice of the unread portion of the buffer. If the caller
// changes the contents of the returned slice, the contents of the buffer will
// change, provided there are no intervening method calls on the Buffer.
func (b *Buffer) Samples() Slice {
	return b.buf.Slice(b.off, b.buf.Len())
}

// Len returns the number of samples of the unread portion of the buffer;
// b.Len() == b.Samples().Len().
func (b *Buffer) Len() int {
	return b.buf.Len() - b.off
}

// Truncate discards all but the first n unread samples from the buffer.
// It panics if n is negative or greater than the length of the buffer.
func (b *Buffer) Truncate(n int) {
	switch {
	case n < 0 || n > b.Len():
		panic("audio.Buffer: truncation out of range")
	case n == 0:
		// Reuse buffer space.
		b.off = 0
	}
	b.buf = b.buf.Slice(0, b.off+n)
}

// Reset resets the buffer so it has no content.
// b.Reset() is the same as b.Truncate(0).
func (b *Buffer) Reset() { b.Truncate(0) }

// grow grows the buffer to guarantee space for n more samples.
// It returns the index where samples should be written.
func (b *Buffer) grow(n int) int {
	m := b.Len()
	// If buffer is empty, reset to recover space.
	if m == 0 && b.off != 0 {
		b.Truncate(0)
	}
	if b.buf.Len()+n > b.buf.Cap() {
		newLength := 2*b.buf.Cap() + n
		buf := b.buf.Make(newLength, newLength)
		b.buf.Slice(b.off, b.buf.Len()).CopyTo(buf)
		b.buf = buf
		b.off = 0
	}
	b.buf = b.buf.Slice(0, b.off+m+n)
	return b.off + m
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for
// another n samples. After Grow(n), at least n samples can be written to the
// buffer without another allocation.
// If n is negative, Grow will panic.
func (b *Buffer) Grow(n int) {
	if n < 0 {
		panic("audio.Buffer.Grow: negative count")
	}
	m := b.grow(n)
	b.buf = b.buf.Slice(0, m)
}

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil.
func (b *Buffer) Write(p Slice) (n int, err error) {
	m := b.grow(p.Len())
	return p.CopyTo(b.buf.Slice(m, b.buf.Len())), nil
}

// MinRead is the minimum slice size passed to a Read call by
// Buffer.ReadFrom.  As long as the Buffer has at least MinRead samples beyond
// what is required to hold the contents of r, ReadFrom will not grow the
// underlying buffer.
const minRead = 512

// ReadFrom reads data from r until EOS and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of samples read. Any
// error except EOS encountered during the read is also returned.
func (b *Buffer) ReadFrom(r Reader) (n int64, err error) {
	// If buffer is empty, reset to recover space.
	if b.off >= b.buf.Len() {
		b.Truncate(0)
	}
	for {
		if free := b.buf.Cap() - b.buf.Len(); free < minRead {
			// not enough space at end
			newBuf := b.buf
			if b.off+free < minRead {
				// not enough space using beginning of buffer;
				// double buffer capacity
				newLen := 2*b.buf.Cap() + minRead
				newBuf = b.buf.Make(newLen, newLen)
			}
			b.buf.Slice(b.off, b.buf.Len()).CopyTo(newBuf)
			b.buf = newBuf.Slice(0, b.buf.Len()-b.off)
			b.off = 0
		}
		m, e := r.Read(b.buf.Slice(b.buf.Len(), b.buf.Cap()))
		b.buf = b.buf.Slice(0, b.buf.Len()+m)
		n += int64(m)
		if e == EOS {
			break
		}
		if e != nil {
			return n, e
		}
	}
	return n, nil // err is EOS, so return nil explicitly
}

// WriteTo writes data to w until the buffer is drained or an error occurs.
// The return value n is the number of samples written; it always fits into an
// int, but it is int64 to match the WriterTo interface. Any error
// encountered during the write is also returned.
func (b *Buffer) WriteTo(w Writer) (n int64, err error) {
	if b.off < b.buf.Len() {
		nSamples := b.Len()
		m, e := w.Write(b.buf.Slice(b.off, b.buf.Len()))
		if m > nSamples {
			panic("audio.Buffer.WriteTo: invalid Write count")
		}
		b.off += m
		n = int64(m)
		if e != nil {
			return n, e
		}
		// all samples should have been written, by definition of
		// Write method in Writer
		if m != nSamples {
			return n, ErrShortWrite
		}
	}
	// Buffer is now empty; reset.
	b.Truncate(0)
	return
}

// WriteSample appends the sample c to the buffer, growing the buffer as needed.
func (b *Buffer) WriteSample(c float64) {
	m := b.grow(1)
	b.buf.Set(m, c)
}

// Read reads the next p.Len() samples from the buffer or until the buffer
// is drained.  The return value n is the number of samples read.  If the
// buffer has no data to return, err is EOS (unless len(p) is zero);
// otherwise it is nil.
func (b *Buffer) Read(p Slice) (n int, err error) {
	if b.off >= b.buf.Len() {
		// Buffer is empty, reset to recover space.
		b.Truncate(0)
		if p.Len() == 0 {
			return
		}
		return 0, EOS
	}
	n = b.buf.Slice(b.off, b.buf.Len()).CopyTo(p)
	b.off += n
	return
}

// Next returns a slice containing the next n samples from the buffer,
// advancing the buffer as if the samples had been returned by Read.
// If there are fewer than n samples in the buffer, Next returns the entire buffer.
// The slice is only valid until the next call to a read or write method.
func (b *Buffer) Next(n int) Slice {
	m := b.Len()
	if n > m {
		n = m
	}
	data := b.buf.Slice(b.off, b.off+n)
	b.off += n
	return data
}

// ReadSample reads and returns the next sample from the buffer.
// If no sample is available, it returns error EOS.
func (b *Buffer) ReadSample() (c float64, err error) {
	if b.off >= b.buf.Len() {
		// Buffer is empty, reset to recover space.
		b.Truncate(0)
		return 0, EOS
	}
	c = b.buf.At(b.off)
	b.off++
	return c, nil
}

// Seek seeks to the specified sample number, relative to the start of the
// stream. As such, subsequent Read() calls on the Buffer, begin reading at
// the specified sample.
//
// If offset > b.Len(), then the offset is unchanged and the seek operation
// fails returning error == EOS.
func (b *Buffer) Seek(offset uint64) error {
	if int(offset) > b.Len() {
		return EOS
	}
	b.off = int(offset)
	return nil
}

// NewBuffer creates and initializes a new Buffer using buf as its initial
// contents. The buffer will internally use the given slice, buf, which also
// defines the sample storage type. NewBuffer is intended to prepare a Buffer
// to read existing data. It can also be used to size the internal buffer for
// writing. To do that, buf should have the desired capacity but a length of
// zero.
func NewBuffer(buf Slice) *Buffer {
	return &Buffer{buf: buf}
}
