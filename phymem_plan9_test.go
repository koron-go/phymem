package phymem

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadMemSize(t *testing.T) {
	s := fmt.Sprintf("%-27s %-27s %-11s %11d %11d %11d %11d %11d %11d %11d %11d %11d ",
		"name", "user", "Await", 0, 0, 39170, 0, 10, 0, 236, 10, 10)
	r := strings.NewReader(s)
	msize, err := readMemSize(r)
	if err != nil {
		t.Fatal("readMemSize:", err)
	}
	const want = 236 * 1024
	if msize != want {
		t.Errorf("readMemSize(%q) = %d; want %d", s, msize, want)
	}
}
