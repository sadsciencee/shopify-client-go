package cmd

import (
	"fmt"
	"github.com/Khan/genqlient/generate"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var gen = &cobra.Command{
	Use:   "gen",
	Short: "gen - Generate the client",
	Long:  `Generate the client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generate")
		config, err := generate.ReadAndValidateConfig("genqlient.yaml")
		if err != nil {
			fmt.Printf("Error reading config: %v\n", err)
			os.Exit(1)
		}
		res, err := generate.Generate(config)
		if err != nil {
			fmt.Printf("Error generating client: %v\n", err)
			os.Exit(1)
		}
		// Write each generated file to disk
		for filename, content := range res {
			// Ensure the directory exists
			dir := filepath.Dir(filename)
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
				os.Exit(1)
			}

			// Write the file
			if err := os.WriteFile(filename, content, 0644); err != nil {
				fmt.Printf("Error writing file %s: %v\n", filename, err)
				os.Exit(1)
			}
			fmt.Printf("Generated: %s\n", filename)
		}
	},
}
