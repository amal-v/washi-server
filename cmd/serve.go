package cmd

import (
	"net/http"

	"washi/api"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serve)

}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start Server",
	Long:  `Can be used to start server at a particular PORT Eg serve --PORT 3000`,
	Run: func(cmd *cobra.Command, args []string) {
		http.ListenAndServe(":3000", api.Router())
	},
}
