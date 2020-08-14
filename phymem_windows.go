package phymem

import (
	"syscall"
	"unsafe"

	"github.com/koron-go/phymem/internal/psapi"
)

// for test.
const providedCurrent = true

// Current get physical memory which used by current process.
func Current() (uint, error) {
	m, err := currentPMC()
	if err != nil {
		return 0, err
	}
	return uint(m.WorkingSetSize), nil
}

var zeroPMC = psapi.ProcessMemoryCounters{}

func currentPMC() (psapi.ProcessMemoryCounters, error) {
	p, err := syscall.GetCurrentProcess()
	if err != nil {
		return zeroPMC, err
	}
	var m psapi.ProcessMemoryCounters
	err = psapi.GetProcessMemoryInfo(p, &m, uint32(unsafe.Sizeof(m)))
	if err != nil {
		return zeroPMC, err
	}
	return m, nil
}

// ProcessMemoryCounters is set of counters related memory of the process.
// This type is provided for Windows only.
type ProcessMemoryCounters struct {
	PageFaultCount             uint32
	PeakWorkingSetSize         uint
	WorkingSetSize             uint
	QuotaPeakPagedPoolUsage    uint
	QuotaPagedPoolUsage        uint
	QuotaPeakNonPagedPoolUsage uint
	QuotaNonPagedPoolUsage     uint
	PagefileUsage              uint
	PeakPagefileUsage          uint
}

func toProcessMemoryCounters(src psapi.ProcessMemoryCounters) ProcessMemoryCounters {
	return ProcessMemoryCounters{
		PageFaultCount:             src.PageFaultCount,
		PeakWorkingSetSize:         uint(src.PeakWorkingSetSize),
		WorkingSetSize:             uint(src.WorkingSetSize),
		QuotaPeakPagedPoolUsage:    uint(src.QuotaPeakPagedPoolUsage),
		QuotaPagedPoolUsage:        uint(src.QuotaPagedPoolUsage),
		QuotaPeakNonPagedPoolUsage: uint(src.QuotaPeakNonPagedPoolUsage),
		QuotaNonPagedPoolUsage:     uint(src.QuotaNonPagedPoolUsage),
		PagefileUsage:              uint(src.PagefileUsage),
		PeakPagefileUsage:          uint(src.PeakPagefileUsage),
	}
}

// CurrentProcessMemoryCounters get ProcessMemoryCounters for current process.
// This function is provided for Windows only.
func CurrentProcessMemoryCounters() (ProcessMemoryCounters, error) {
	m, err := currentPMC()
	if err != nil {
		return ProcessMemoryCounters{}, err
	}
	return toProcessMemoryCounters(m), nil
}
