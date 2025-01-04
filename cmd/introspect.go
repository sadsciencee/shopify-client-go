package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Aoc - Pew Pew",
	Long:  `Aoc is a command line utility to super charge your Advent of Code experience`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(year2024.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
