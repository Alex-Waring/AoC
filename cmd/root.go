package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "A simple CLI tool to import and prepare a day in AoC",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.Execute()
}
