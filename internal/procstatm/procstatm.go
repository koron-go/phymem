package procstatm

import (
	"fmt"
	"os"
)

// Statm provides statistics of memory.
type Statm struct {
	Size     uint
	Resident uint
	Share    uint
	Text     uint
	Lib      uint
	Data     uint
	Darty    uint
}

var zero = Statm{}

// Get gets Statm for process (pid).
// pid should be process id (integer) or string ("self" or so).
func Get(pid interface{}) (Statm, error) {
	name := fmt.Sprintf("/proc/%s/statm", pid)
	f, err := os.Open(name)
	if err != nil {
		return zero, err
	}
	defer f.Close()
	m := Statm{}
	_, err = fmt.Fscanf(f, "%d %d %d %d %d %d %d",
		&m.Size, &m.Resident, &m.Share, &m.Text, &m.Lib, &m.Data, &m.Darty)
	if err != nil {
		return zero, fmt.Errorf("failed to scan %s: %w", name, err)
	}
	return m, nil
}
