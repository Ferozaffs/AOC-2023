package day4

import (
	"aoc2023/cmd"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day4_2Cmd = &cobra.Command{
	Use:   "day4_2",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run2()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day4_2Cmd)
}

func Run2() {
	dat, _ := os.ReadFile("inputs/day4_data.txt")
	ans := Solve2(string(dat))

	fmt.Printf("ANSWER: %d\n", ans)
}

func Solve2(data string) int {
	sum := 0

	lines := []string{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		index := regexp.MustCompile(`\|`).FindStringIndex(lines[i])

		winningNumbers := []int{}
		matches := 0
		card := 0
		for j, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(lines[i], -1) {
			//Skip card ID
			if j == 0 {
				card, _ = strconv.Atoi(lines[i][m[0]:m[1]])
				continue
			}

			n, _ := strconv.Atoi(lines[i][m[0]:m[1]])
			if m[0] < index[0] {
				winningNumbers = append(winningNumbers, n)
			} else if slices.Contains(winningNumbers, n) {
				matches++
			}
		}
		for j := 0; j < matches; j++ {
			lines = append(lines, lines[card+j])
		}
		sum++
	}

	return sum
}
