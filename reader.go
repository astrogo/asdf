// Copyright 2015 The astrogo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asdf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Open open an ASDF file for reading.
func Open(r io.Reader) (File, error) {
	var f File
	var err error
	rbuf := bufio.NewReader(r)
	bline, err := rbuf.ReadBytes('\n')
	if err != nil {
		return f, err
	}
	line := string(bytes.TrimSpace(bline))
	if !strings.HasPrefix(line, "#ASDF") {
		return f, fmt.Errorf("asdf: missing version header (got=%q)", line)
	}
	f.Version = line

comment_loop:
	for {
		c, err := rbuf.Peek(1)
		if err != nil {
			return f, err
		}
		switch c[0] {
		case '#':
			bline, err = rbuf.ReadBytes('\n')
			if err != nil {
				return f, err
			}
			f.Comments = append(
				f.Comments,
				string(bytes.TrimSpace(bline)),
			)
		default:
			break comment_loop
		}
	}

	c, err := rbuf.Peek(1)
	if err != nil {
		if err == io.EOF {
			// asdf file with only a header (and comments)
			return f, nil
		}
		return f, err
	}
	if c[0] == '%' {
		ybuf := []byte{}
		for {
			bline, err = rbuf.ReadBytes('\n')
			if err != nil {
				return f, err
			}
			// FIXME(sbinet)
			// see: https://github.com/go-yaml/yaml/issues/74
			if bytes.HasPrefix(bline, []byte("%TAG ")) {
				bline = append(
					[]byte("%TAG !asdf"),
					bline[len([]byte("%TAG ")):]...,
				)
			}
			ybuf = append(ybuf, bline...)
			if bytes.HasPrefix(bline, []byte("...")) {
				break
			}
		}
		fmt.Printf("=== yaml ===\n%s\n===\n", string(ybuf))

		var ymap map[string]interface{}
		err = yaml.Unmarshal(ybuf, &ymap)
		if err != nil {
			return f, err
		}
		f.Tree = &Tree{data: ymap}
	}

	err = decodeBlocks(rbuf, &f)
	if err != nil {
		return f, err
	}

	if len(f.Blocks) > 0 {
		// only asdf files with at least a block may have an index
		err = decodeIndex(rbuf, &f)
		if err != nil {
			return f, err
		}
	}

	return f, err
}

func decodeBlocks(rbuf *bufio.Reader, f *File) error {

	for {
		c, err := rbuf.Peek(len(blockMagicToken))
		if err != nil {
			fmt.Printf("peek failed: %v\n", err)
			if err == io.EOF {
				return nil
			}
			return err
		}
		if !bytes.Equal(c, blockMagicToken) {
			return nil
		}
		var block Block
		err = decodeBlock(rbuf, &block)
		if err != nil {
			return err
		}
		f.Blocks = append(f.Blocks, block)
	}
	return nil
}

func decodeBlock(r io.Reader, blk *Block) error {
	var err error

	read := func(v interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, v)
	}

	read(blk.Header.Magic[:])
	read(&blk.Header.Size)
	read(&blk.Header.Flags)
	read(blk.Header.Compression[:])
	read(&blk.Header.AllocSize)
	read(&blk.Header.UsedSize)
	read(&blk.Header.DataSize)
	read(blk.Header.Checksum[:])

	if err != nil {
		return err
	}
	blk.Data = make([]byte, int(blk.Header.AllocSize))
	read(blk.Data)
	return err
}

func decodeIndex(rbuf *bufio.Reader, f *File) error {
	c, err := rbuf.Peek(len(blockIndex))
	if err != nil {
		if err == io.EOF {
			return nil
		}
		fmt.Printf("index-peek failed: %v\n", err)
		return err
	}
	if !bytes.Equal(c, blockIndex) {
		return nil
	}
	var ybuf []byte
	for {
		bline, err := rbuf.ReadBytes('\n')
		if err != nil {
			return err
		}
		ybuf = append(ybuf, bline...)
		if bytes.HasPrefix(bline, []byte("...")) {
			break
		}
	}
	err = yaml.Unmarshal(ybuf, &f.Index)
	return err
}
