package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var region string
	rootCmd.AddCommand(serve)
	rootCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region (required)")
	rootCmd.MarkFlagRequired("region")

}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start Server",
	Long:  `Can be used to start server at a particular PORT Eg serve --PORT 3000`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
