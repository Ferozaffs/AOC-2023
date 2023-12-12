package day12

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `???#...#.?#??#?#? 1,1,1,8`

	ans := Solve(sampleData, 5)
	if ans != 16384 {
		t.Fatalf("Mismatch! Expected 16384 got %d", ans)
	}
}
