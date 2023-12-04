package day4

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day4_1Cmd = &cobra.Command{
	Use:   "day4_1",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day4_1Cmd)
}

func Run1() {
	dat, _ := os.ReadFile("inputs/day4_data.txt")
	ans := Solve1(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve1(data string) int {
	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := scanner.Text()
		index := regexp.MustCompile(`\|`).FindStringIndex(s)

		winningNumbers := []int{}
		matches := 0
		for i, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(s, -1) {
			//Skip card ID
			if i == 0 {
				continue
			}

			n, _ := strconv.Atoi(s[m[0]:m[1]])
			if m[0] < index[0] {
				winningNumbers = append(winningNumbers, n)
			} else if slices.Contains(winningNumbers, n) {
				matches++
			}
		}

		sum += int(math.Pow(2, float64(matches-1)))

	}

	return sum
}
