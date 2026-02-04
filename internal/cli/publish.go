package cli

import (
	"encoding/json"
	"fmt"

	"github.com/H0R15H0/chrome-webstore-api-v2/chromewebstore"
	"github.com/spf13/cobra"
)

var publishTarget string

func init() {
	publishCmd.Flags().StringVar(&publishTarget, "target", "", "Publish target: 'testers' for trusted testers only")
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish an item",
	Long:  `Publish a Chrome Web Store item to the specified audience.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createClient()
		if err != nil {
			return err
		}

		itemName, err := getItemName()
		if err != nil {
			return err
		}

		call := client.Publishers.Items.Publish(itemName)

		if publishTarget == "testers" {
			call.PublishTarget(chromewebstore.PublishTargetTrustedTesters)
		}

		result, err := call.Do()
		if err != nil {
			return fmt.Errorf("failed to publish: %w", err)
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
