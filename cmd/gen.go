package cmd

import (
	"fmt"
	"github.com/Khan/genqlient/generate"
	"github.com/spf13/cobra"
)

var gen = &cobra.Command{
	Use:   "gen",
	Short: "gen - Generate the client",
	Long:  `Generate the client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generate")
		/*dir, _ := os.Getwd()
		fmt.Printf("Current working directory: %s\n", dir)
		config := generate.Config{}
		config.ValidateAndFillDefaults(dir)
		_, err := generate.Generate(&config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}*/
		generate.Main()
	},
}
