// +build !windows
// +build !linux

package phymem

import "errors"

// for test.
const providedCurrent = false

var notImpl = errors.New("Current() is not implemented for this platform")

// Current get physical memory which used by current process.
func Current() (uint, error) {
	return 0, notImpl
}
