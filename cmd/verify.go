package cmd

import (
	"fmt"

	"github.com/mickaelpham/znt/diff"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify notifications exist",
	Long: `
Query all notification definitions for the given Zuora
environment, and verify they match the template.`,
	Run: func(cmd *cobra.Command, args []string) {
		// triggers := diff.FetchManagedTriggers()
		// fmt.Printf("Found %d event triggers\n", len(triggers))

		notifications := diff.FetchManagedNotifications()
		fmt.Printf("Found %d notification\n", len(notifications))
		fmt.Println(notifications)
	},
}
