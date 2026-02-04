package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cancelSubmissionCmd)
}

var cancelSubmissionCmd = &cobra.Command{
	Use:   "cancel-submission",
	Short: "Cancel a pending submission",
	Long:  `Cancel a pending submission for a Chrome Web Store item.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createClient()
		if err != nil {
			return err
		}

		itemName, err := getItemName()
		if err != nil {
			return err
		}

		result, err := client.Publishers.Items.CancelSubmission(itemName).Do()
		if err != nil {
			return fmt.Errorf("failed to cancel submission: %w", err)
		}

		if jsonOutput {
			output, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal JSON: %w", err)
			}
			fmt.Println(string(output))
		} else {
			fmt.Printf("Status: %s\n", result.StatusCode)
			if len(result.StatusDetail) > 0 {
				fmt.Println("Details:")
				for _, detail := range result.StatusDetail {
					fmt.Printf("  - %s\n", detail)
				}
			}
		}

		return nil
	},
}
