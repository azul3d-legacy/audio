// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import "testing"

func TestBufferIO(t *testing.T) {
	buf := NewBuffer(Int16{})
	_ = Reader(buf)
	_ = ReadSeeker(buf)
	_ = Writer(buf)
	_ = WriterTo(buf)
	_ = ReaderFrom(buf)
}
