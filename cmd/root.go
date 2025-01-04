package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "shopify-client-go",
	Short: "Initialize and introspect Shopify client",
	Long:  `Tooling to automate schema introspection and client generation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To introspect a Shopify schema, run `shopify-client-go introspect`")
		fmt.Println("To generate a Shopify client, run `shopify-client-go generate`")
	},
}

func init() {
	rootCmd.AddCommand(introspect)
	rootCmd.AddCommand(gen)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
