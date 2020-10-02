package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mickaelpham/znt/diff"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the diff",
	Long: `
Apply the triggers diff and notification diff to
the targeted Zuora environment`,
	Run: func(cmd *cobra.Command, args []string) {
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

		prompt := promptui.Prompt{
			Label:     "Apply changes to Zuora",
			IsConfirm: true,
		}

		proceed, _ := prompt.Run()
		if proceed != "y" {
			return
		}

		triggerDiff.Apply()
	},
}
