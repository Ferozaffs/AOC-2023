package day19

import (
	"aoc2023/cmd"

	"github.com/spf13/cobra"
)

var day19_1Cmd = &cobra.Command{
	Use:   "day19_1",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Run1()
	},
}

func init() {
	cmd.RootCmd.AddCommand(day19_1Cmd)
}

func Run1() {

}

func Solve1() {

}
