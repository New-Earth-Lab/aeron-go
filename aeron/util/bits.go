/*
Copyright 2016 Stanislav Liberman
Copyright 2023 Rubus Technologies Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"unsafe"
)

var i32 int32
var i64 int64

const (
	// CacheLineLength is a constant for the size of a CPU cache line
	CacheLineLength int32 = 64

	// SizeOfInt32 is a constant for the size of int32. Ha. Just for Clarity
	SizeOfInt32 int32 = int32(unsafe.Sizeof(i32))

	// SizeOfInt64 is a constant for the size of int64
	SizeOfInt64 int32 = int32(unsafe.Sizeof(i64))
)

// AlignInt32 will return a number rounded up to the alignment boundary
func AlignInt32(value, alignment int32) int32 {
	return (value + (alignment - 1)) & ^(alignment - 1)
}

// IsPowerOfTwo checks that the argument number is a power of two
func IsPowerOfTwo(value int64) bool {
	return value > 0 && ((value & (^value + 1)) == value)
}

func MemPrint(ptr uintptr, len int) string {
	var output string

	for i := 0; i < len; i += 1 {
		ptr := unsafe.Pointer(ptr + uintptr(i))
		output += fmt.Sprintf("%02x ", *(*int8)(ptr))
	}

	return output
}

func Print(bytes []byte) {
	for i, b := range bytes {
		if i > 0 && i%16 == 0 && i%32 != 0 {
			fmt.Print(" :  ")
		}
		if i > 0 && i%32 == 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%02x ", b)
	}
	fmt.Print("\n")
}
