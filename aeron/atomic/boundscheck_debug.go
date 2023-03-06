//go:build debug

package atomic

import (
	"log"
)

// BoundsCheck is helper function to make sure buffer writes and reads to
// not go out of bounds on stated buffer capacity
//
//go:norace
func BoundsCheck(index int32, length int32, myLength int32) {
	if (index + length) > myLength {
		log.Fatalf("Out of Bounds. int32: %d + %d Capacity: %d", index, length, myLength)
	}
}
