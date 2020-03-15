package phymem

import "github.com/koron-go/phymem/internal/procstatm"

// for test.
const providedCurrent = true

// Current get physical memory which used by current process.
func Current() (uint, error) {
	m, err := procstatm.Get("self")
	if err != nil {
		return 0, err
	}
	return m.Resident * 4096, nil
}
