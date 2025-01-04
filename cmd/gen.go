package cmd

import (
	"fmt"
	"github.com/Khan/genqlient/generate"
	"github.com/spf13/cobra"
	"os"
)

var gen = &cobra.Command{
	Use:   "gen",
	Short: "gen - Generate the client",
	Long:  `Generate the client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generate")
		dir, _ := os.Getwd()
		fmt.Printf("Current working directory: %s\n", dir)

		// Add debug output for files in current directory
		files, _ := os.ReadDir(".")
		for _, f := range files {
			fmt.Printf("Found file: %s\n", f.Name())
		}

		generate.Main()
	},
}
