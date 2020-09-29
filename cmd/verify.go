package cmd

import "github.com/spf13/cobra"

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify notifications exist",
	Long: `
Query all notification definitions for the given Zuora
environment, and verify they match the template.`,
}
