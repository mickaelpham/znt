package cmd

import (
	"fmt"

	"github.com/mickaelpham/znt/auth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify notifications exist",
	Long: `
Query all notification definitions for the given Zuora
environment, and verify they match the template.`,
	Run: func(cmd *cobra.Command, args []string) {

		token := auth.NewToken(
			viper.GetString("baseurl"),
			viper.GetString("client"),
			viper.GetString("secret"),
		)

		fmt.Println(token)
	},
}
