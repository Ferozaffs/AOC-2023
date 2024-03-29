package day21

import (
	"testing"
)

func TestSample2(t *testing.T) {
	sampleData := `.................................
	.....###.#......###.#......###.#.
	.###.##..#..###.##..#..###.##..#.
	..#.#...#....#.#...#....#.#...#..
	....#.#........#.#........#.#....
	.##...####..##...####..##...####.
	.##..#...#..##..#...#..##..#...#.
	.......##.........##.........##..
	.##.#.####..##.#.####..##.#.####.
	.##..##.##..##..##.##..##..##.##.
	.................................
	.................................
	.....###.#......###.#......###.#.
	.###.##..#..###.##..#..###.##..#.
	..#.#...#....#.#...#....#.#...#..
	....#.#........#.#........#.#....
	.##...####..##..S####..##...####.
	.##..#...#..##..#...#..##..#...#.
	.......##.........##.........##..
	.##.#.####..##.#.####..##.#.####.
	.##..##.##..##..##.##..##..##.##.
	.................................
	.................................
	.....###.#......###.#......###.#.
	.###.##..#..###.##..#..###.##..#.
	..#.#...#....#.#...#....#.#...#..
	....#.#........#.#........#.#....
	.##...####..##...####..##...####.
	.##..#...#..##..#...#..##..#...#.
	.......##.........##.........##..
	.##.#.####..##.#.####..##.#.####.
	.##..##.##..##..##.##..##..##.##.
	.................................`

	ans := Solve1(sampleData, 6, true)
	if ans != 16 {
		t.Fatalf("Mismatch! Expected 16 got %d", ans)
	}

	ans = Solve1(sampleData, 10, true)
	if ans != 50 {
		t.Fatalf("Mismatch! Expected 50 got %d", ans)
	}

	ans = Solve1(sampleData, 50, true)
	if ans != 1594 {
		t.Fatalf("Mismatch! Expected 1594 got %d", ans)
	}

	ans = Solve1(sampleData, 100, true)
	if ans != 6536 {
		t.Fatalf("Mismatch! Expected 6536 got %d", ans)
	}

	ans = Solve1(sampleData, 500, true)
	if ans != 167004 {
		t.Fatalf("Mismatch! Expected 167004 got %d", ans)
	}

	ans = Solve1(sampleData, 1000, true)
	if ans != 668697 {
		t.Fatalf("Mismatch! Expected 668697 got %d", ans)
	}

	ans = Solve1(sampleData, 5000, true)
	if ans != 16733044 {
		t.Fatalf("Mismatch! Expected 16733044 got %d", ans)
	}
}
