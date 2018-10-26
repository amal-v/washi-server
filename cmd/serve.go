package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func init() {
	var region int16
	rootCmd.AddCommand(serve)
	serve.Flags().Int16VarP(&region, "PORT", "P", 3000, "Server PORT (Optional)")
	viper.BindPFlag("PORT", serve.Flags().Lookup("PORT"))

}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start Server",
	Long:  `Can be used to start server at a particular PORT Eg serve --PORT 3000`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("PORT"))

	},
}
