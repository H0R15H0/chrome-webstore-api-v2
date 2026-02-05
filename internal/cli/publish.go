package cli

import (
	"encoding/json"
	"fmt"

	"github.com/H0R15H0/chrome-webstore-api-v2/chromewebstore"
	"github.com/spf13/cobra"
)

var (
	publishType      string
	deployPercentage int
)

func init() {
	publishCmd.Flags().StringVar(&publishType, "type", "", "Publish type: 'immediate' or 'staged'")
	publishCmd.Flags().IntVar(&deployPercentage, "deploy-percentage", 0, "Deploy percentage for staged rollout (0-100)")
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish an item",
	Long:  `Publish a Chrome Web Store item.`,
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

		if publishType == "immediate" {
			call.PublishType(chromewebstore.PublishTypeImmediate)
		} else if publishType == "staged" {
			call.PublishType(chromewebstore.PublishTypeStaged)
		}

		if deployPercentage > 0 {
			call.DeployPercentage(deployPercentage)
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
			fmt.Printf("Name:    %s\n", result.Name)
			fmt.Printf("Item ID: %s\n", result.ItemID)
			fmt.Printf("State:   %s\n", result.State)
		}

		return nil
	},
}
