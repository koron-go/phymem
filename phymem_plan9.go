package phymem

import (
	"os"

	"github.com/koron-go/phymem/internal/procstatus"
)

// for test.
const providedCurrent = true

// Current get physical memory which used by current process.
func Current() (uint, error) {
	return procstatus.GetRSS(os.Getpid())
}
