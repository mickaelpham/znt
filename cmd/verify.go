package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

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

		// notifications := diff.FetchManagedNotifications()
		// fmt.Printf("Found %d notification\n", len(notifications))
		// fmt.Println(notifications)

		// profiles := diff.FetchProfiles()
		// fmt.Printf("Found %d communication profiles\n", len(profiles))
		// fmt.Println(profiles)

		f, err := os.Open(tplFile)
		if err != nil {
			log.Fatal(err)
		}

		tpl, err := diff.Parse(bufio.NewReader(f))
		if err != nil {
			log.Fatal(err)
		}

		triggerDiff := diff.NewTriggerDiff(tpl.Triggers(), diff.FetchManagedTriggers())
		fmt.Println(triggerDiff)

		profiles := diff.FetchProfiles()
		fmt.Println("--- Communication Profiles")
		for name, ID := range profiles {
			fmt.Printf("  * (%s) %s\n", ID, name)
		}
		fmt.Println()

		notificationDiff := diff.NewNotificationDiff(tpl.NotificationDefinitions(profiles), diff.FetchManagedNotifications())
		fmt.Println(notificationDiff)
	},
}
