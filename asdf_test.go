// Copyright 2015 The astrogo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asdf

import (
	"fmt"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	r, err := os.Open("testdata/ascii.asdf")
	if err != nil {
		t.Fatalf("could not open file: %v\n", err)
	}
	defer r.Close()

	f, err := Open(r)
	if err != nil {
		t.Fatalf("could not open asdf file: %v\n", err)
	}

	fmt.Printf("f.version=%q\n", f.Version)
	fmt.Printf("f.comments=%v\n", f.Comments)
	fmt.Printf("f.tree=%v\n", f.Tree)
	fmt.Printf("f.blocks=%v\n", f.Blocks)
	fmt.Printf("f.index=%v\n", f.Index)
}
