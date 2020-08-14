package phymem

import "testing"

func TestCurrentProcessMemoryCounters(t *testing.T) {
	c, err := CurrentProcessMemoryCounters()
	if err != nil {
		t.Fatalf("CurrentProcessMemoryCounters() failed: %s", err)
	}
	if c.WorkingSetSize == 0 {
		t.Fatal("WorkingSetSize is zero")
	}
	t.Logf("ProcessMemoryCounters=%#v", c)
}
