package day12

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var day12_1Cmd = &cobra.Command{
	Use:   "day12_1",
	Short: "Day 12, challenge 1",
	Long:  `Day 12, challenge 1`,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day12_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day12_data.txt")
	ans := Solve(string(dat), 1)

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve(data string, multiplier int) int {
	sum := 0

	reIndex := regexp.MustCompile("\\d+")

	scanner := bufio.NewScanner(strings.NewReader(data))
	row := 0
	for scanner.Scan() {
		fmt.Printf("Current row/sum: %d  |  %d\n", row, sum)
		row++

		neededSlots := 0
		index := []int{}

		indexScanned := reIndex.FindAllString(scanner.Text(), -1)
		for _, i := range indexScanned {
			n, _ := strconv.Atoi(i)
			index = append(index, n)
			neededSlots += n
		}

		split := strings.Fields(scanner.Text())
		springs := split[0]

		if multiplier > 1 {
			orgIndex := index
			orgSprings := springs

			for i := 0; i < multiplier-1; i++ {
				for _, idx := range orgIndex {
					index = append(index, idx)
					neededSlots += idx
				}

				springs += "?" + orgSprings
			}
		}

		slots := 0
		for i := 0; i < len(springs); i++ {
			if springs[i] == '#' || springs[i] == '?' {
				slots++
			}
		}

		s := time.Now()

		result := Count(0, slots, neededSlots, []byte(springs), index)

		fmt.Print(time.Since(s))
		fmt.Print("\n")

		sum += result
	}

	return sum
}

func Count(start int, slots int, neededSlots int, springs []byte, index []int) int {
	sum := 0

	//Pre check
	if slots < neededSlots {
		return sum
	}

	//Expand "?"
	for i := start; i < len(springs); i++ {
		if springs[i] == '?' {
			springs[i] = '.'
			sum += Count(i+1, slots-1, neededSlots, springs, index)

			springs[i] = '#'
			sum += Count(i+1, slots, neededSlots, springs, index)

			springs[i] = '?'
			return sum
		}
	}

	//Else check if valid
	if Check(springs, index) == true {
		sum++
	}

	return sum
}

func Check(springs []byte, index []int) bool {
	idx := 0
	for i := 0; i < len(springs); i++ {
		if springs[i] == '#' {
			start := i
			for ; i < len(springs); i++ {
				if springs[i] == '.' {
					break
				}
			}

			if idx < len(index) || index[idx] != i-start {
				return false
			} else {
				idx++
			}
		}
	}

	if len(index) == idx {
		return true
	}

	return false
}
