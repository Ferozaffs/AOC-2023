package day9

import (
	"aoc2023/cmd"

	"github.com/spf13/cobra"
)

var day9_1Cmd = &cobra.Command{
	Use:   "day9_1",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day9_1Cmd)
}

func Run1() {

}

func Solve1() {

}
