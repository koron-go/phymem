package phymem

import "testing"

func TestCurrent(t *testing.T) {
	if !providedCurrent {
		t.Skip("Current() is not provided")
	}
	n, err := Current()
	if err != nil {
		t.Fatalf("Current() failed: %s", err)
	}
	if n == 0 {
		t.Fatal("Current() returns zero")
	}
}
