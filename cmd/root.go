package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thomaszub/advent-of-code-2023/day1"
	"github.com/thomaszub/advent-of-code-2023/day2"
)

var rootCmd = &cobra.Command{
	Use:   "advent-of-code-2023",
	Short: "Command line tool for the Advent of Code 2023",
	Long:  "Command line tool for the Advent of Code 2023",
}

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Execute the solution for day 1",
	Long:  "Execute the solution for day 1",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day1.Execute()
	},
}

var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Execute the solution for day 2",
	Long:  "Execute the solution for day 2",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day2.Execute()
	},
}

func Execute() {
	rootCmd.AddCommand(day1Cmd)
	rootCmd.AddCommand(day2Cmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
