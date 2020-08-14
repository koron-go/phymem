package phymem

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// for test.
const providedCurrent = true

func Current() (uint, error) {
	name := fmt.Sprintf("/proc/%d/status", os.Getpid())
	r, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	return readMemSize(r)
}

func readMemSize(r io.Reader) (uint, error) {
	/*
	 * proc/<n>/status contains process's status separated by a space.
	 * 1. [27]the process name
	 * 2. [27]the user name
	 * 3. [11]the process status
	 * 4. [11]the time current process has spent in user mode (ms)
	 * 5. [11]the time current process has spent in system calls (ms)
	 * 6. [11]the time current process has spent in real elapsed time (ms)
	 * 7. [11]the time children and descendants's; user mode (ms)
	 * 8. [11]the time children and descendants's; system calls (ms)
	 * 9. [11]the time children and descendants's; real elapsed time (ms)
	 * 10. [11]the amount of memory (kb)
	 * 11. [11]the base scheduling priority
	 * 12. [11]the current scheduling priority
	 */
	buf := make([]byte, (27+1)*2+(11+1)*10) // +1: a space
	n, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	if n != len(buf) {
		return 0, errors.New("insufficient process status")
	}
	p := (27+1)*2 + (11+1)*7
	s := strings.TrimSpace(string(buf[p : p+11]))
	msize, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(msize) * 1024, nil
}
