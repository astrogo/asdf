// Copyright 2015 The astrogo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import "time"

// ASDF is the top-level attributes for every ASDF file.
//
// tag:stsci.edu:asdf/core/asdf-0.1.0
type ASDF struct {
	Library Software       `yaml:"asdf_library"` // Describes the ASDF library that produced the file
	History HistoryEntries `yaml:"history"`      // Log of transformations that have happened to the file.
	Data    NdArray        `yaml:"data"`         // main science data array in the file
	FITS    interface{}    `yaml:"fits"`         // recipe to convert this ASDF file to a FITS file
	WCS     interface{}    `yaml:"wcs"`          // location of the main WCS for the main data
	Extra   NdArrays       `yaml:"?"`
}

/*
func (asdf *ASDF) UnmarshalYAML(unmarshal func(interface{}) error) error {
	for _, v := range []interface{}{
		&asdf.Library,
		&asdf.History,
		&asdf.Data,
		&asdf.FITS,
		&asdf.WCS,
		&asdf.Extra,
	} {
		//fmt.Printf(">>> i=%d v=%T...\n", i, v)
		err := unmarshal(v)
		if err != nil {
			return err
		}
	}
	return nil
}
*/

// Software describes a software package.
//
// tag:stsci.edu:asdf/core/software-0.1.0
type Software struct {
	Name     string // name of the application or library.
	Author   string // author (or institution) that produced the software package.
	Homepage string // URI to the homepage of the software.
	Version  string // version of the software used.
}

type SoftwareList []Software

func (s *SoftwareList) UnmarshalYAML(unmarshal func(interface{}) error) error {
	for i := range *s {
		v := &(*s)[i]
		err := unmarshal(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// HistoryEntry is an entry in the file history.
//
// tag:stsci.edu:asdf/core/history_entry-0.1.0
type HistoryEntry struct {
	Description string       // description of the transformation performed.
	Time        time.Time    // timestamp for the operation, in UTC.
	Software    SoftwareList // list of softwares that performed the operation.
}

type HistoryEntries []HistoryEntry

func (h *HistoryEntries) UnmarshalYAML(unmarshal func(interface{}) error) error {
	for i := range *h {
		v := &(*h)[i]
		err := unmarshal(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// NdArray is an n-dimensional array.
//
// tag:stsci.edu:asdf/core/ndarray-0.1.0
type NdArray struct {
	Source    interface{}   `yaml:"source"`    // source of the data.
	Data      interface{}   `yaml:"data"`      // data for the array inline
	Shape     []interface{} `yaml:"shape"`     // shape of the array
	DataType  interface{}   `yaml:"datatype"`  // data format of the array elements.
	ByteOrder ByteOrder     `yaml:"byteorder"` // byte order (big- or little-endian) of the array data.
	Offset    int           `yaml:"offset"`    // offset, in bytes, within the data for the start of this view
	Strides   []int         `yaml:"strides"`   // number of bytes to skip in each dimension
	Mask      interface{}   `yaml:"mask"`      // describes how missing values in the array are store.
}

type NdArrays []struct {
	Key   string
	Value NdArray
}

func (arr *NdArrays) UnmarshalYAML(unmarshal func(interface{}) error) error {
	for i := range *arr {
		v := &(*arr)[i]
		err := unmarshal(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// ByteOrder represents the byte order (big- or little-endian) of some byte
// data.
type ByteOrder string

const (
	Big    ByteOrder = "big"
	Little ByteOrder = "little"
)
