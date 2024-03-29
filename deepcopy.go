/*
deepcopy is a package that provides a simple way to make deep copies of objects.
It uses the encoding/gob package to encode and decode objects with a bytes.Buffer to reuse memory for multiple calls.

Example:
package main

import (
	"github.com/renshep/deepcopy"
	"fmt"
)

type TestStruct struct {
	Data []string
}

func main() {
	orig := TestStruct{
		Data: []string{"a", "b", "c"},
	}
	copyBuffer := deepcopy.NewCopyBuffer[TestStruct]()
	copy, copy_err := copyBuffer.DeepCopy(&orig)
	if copy_err != nil {
		panic(copy_err)
	}
	fmt.Println(orig)
	fmt.Println(copy)
}

Copyright (C) 2024 Naomi Shepard

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

package deepcopy

import (
	"bytes"
	"encoding/gob"
)

// CopyBuffer is a struct that holds a bytes.Buffer and an encoding/gob Encoder and Decoder.
// Use NewCopyBuffer to create a new CopyBuffer.
type CopyBuffer[T any] struct {
	buf bytes.Buffer
	enc *gob.Encoder
	dec *gob.Decoder
}

// NewCopyBuffer creates a new CopyBuffer.
// Use this function to create a new CopyBuffer and then call DeepCopy to make deep copies of objects.
func NewCopyBuffer[T any]() *CopyBuffer[T] {
	copyBuffer := new(CopyBuffer[T])
	copyBuffer.enc = gob.NewEncoder(&copyBuffer.buf)
	copyBuffer.dec = gob.NewDecoder(&copyBuffer.buf)
	return copyBuffer
}

// DeepCopy makes a deep copy of an object.
// Multiple calls to DeepCopy will reuse the memory in the CopyBuffer for encoding and decoding.
func (b *CopyBuffer[T]) DeepCopy(orig *T) (copy *T, err error) {
	b.buf.Reset()
	err = b.enc.Encode(orig)
	if err != nil {
		return
	}
	copy = new(T)
	err = b.dec.Decode(copy)
	return
}
