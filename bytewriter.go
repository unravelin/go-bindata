// Copyright 2018 The go-bindata Authors. All rights reserved.
// Use of this source code is governed by a CC0 1.0 Universal (CC0 1.0)
// Public Domain Dedication license that can be found in the LICENSE file.

package bindata

import (
	"fmt"
	"io"
)

//
// ByteWriter define a writer to write content of file.
//
type ByteWriter struct {
	io.Writer
	c int
}

func (w *ByteWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	for n = range p {
		_, err = fmt.Fprintf(w.Writer, "0x%02x,", p[n])
		if err != nil {
			return n, err
		}
		w.c++
	}

	n++

	return n, nil
}
