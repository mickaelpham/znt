package cmd

import (
	"fmt"

	"github.com/mickaelpham/znt/remote"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify notifications exist",
	Long: `
Query all notification definitions for the given Zuora
environment, and verify they match the template.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := remote.NewToken()
		fmt.Println(token)
	},
}
