//go:build !debug

package atomic

// BoundsCheck is helper function to make sure buffer writes and reads to
// not go out of bounds on stated buffer capacity
//
//go:norace
func BoundsCheck(index int32, length int32, myLength int32) {
}
