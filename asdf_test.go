// Copyright 2015 The astrogo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asdf

import (
	"fmt"
	"os"
	"testing"
)

var files = []string{
	"testdata/ascii.asdf",
	"testdata/basic.asdf",
	"testdata/complex.asdf",
	"testdata/compressed.asdf",
	"testdata/exploded0000.asdf",
	"testdata/exploded.asdf",
	"testdata/float.asdf",
	"testdata/int.asdf",
	"testdata/shared.asdf",
	"testdata/stream.asdf",
	"testdata/unicode_bmp.asdf",
	"testdata/unicode_spp.asdf",
}

func TestReader(t *testing.T) {
	for _, fname := range files {
		r, err := os.Open(fname)
		if err != nil {
			t.Fatalf("could not open file [%s]: %v\n", fname, err)
		}
		defer r.Close()

		f, err := Open(r)
		if err != nil {
			t.Fatalf("could not open asdf file [%s]: %v\n", fname, err)
		}

		fmt.Printf("file=%q\n", fname)
		fmt.Printf("f.version=%q\n", f.Version)
		fmt.Printf("f.comments=%v\n", f.Comments)
		fmt.Printf("f.tree=%v\n", f.Tree)
		fmt.Printf("f.blocks=%v\n", f.Blocks)
		fmt.Printf("f.index=%v\n", f.Index)
	}
}
