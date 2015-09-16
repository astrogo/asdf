// Copyright 2015 The astrogo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asdf

var (
	blockMagicToken = []byte("\323BLK")
	blockIndex      = []byte("#ASDF BLOCK INDEX")
)

// File holds informations about an ASDF file
type File struct {
	Version  string // low-level file format version
	Comments []string
	Tree     *Tree
	Blocks   []Block
	Index    []uint64
}

// Tree holds structured informations and meta-data
type Tree struct {
	data map[string]interface{}
}

// Block represents a contiguous chunk of binary data in an ASDF file
type Block struct {
	Header Header
	Data   []byte
}

// Header holds informations a Block
type Header struct {
	Magic       [4]byte // block magic token ("\323BLK")
	Size        uint16  // size of the remainder of the header (not counting Magic nor Size)
	Flags       uint32
	Compression [4]byte
	AllocSize   uint64   // amount of space allocated for the block (not including the header), in bytes
	UsedSize    uint64   // amount of used space for the block on disk (not including the header), in bytes
	DataSize    uint64   // size of the block when decoded, in bytes
	Checksum    [16]byte // MD5 checksum of the used data in the block
}
