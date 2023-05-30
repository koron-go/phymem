//go:build !windows && !linux && !darwin && !freebsd && !plan9 && !netbsd
// +build !windows,!linux,!darwin,!freebsd,!plan9,!netbsd

package phymem

import "errors"

// for test.
const providedCurrent = false

var errNotImpl = errors.New("Current() is not implemented for this platform")

// Current get physical memory which used by current process.
func Current() (uint, error) {
	return 0, errNotImpl
}
