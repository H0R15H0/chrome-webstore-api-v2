package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setPublishedDeployPercentageCmd)
}

var setPublishedDeployPercentageCmd = &cobra.Command{
	Use:   "set-published-deploy-percentage <percentage>",
	Short: "Set the deploy percentage for a published item",
	Long:  `Set the deploy percentage (0-100) for a published Chrome Web Store item.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		percentage, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid percentage: %w", err)
		}

		if percentage < 0 || percentage > 100 {
			return fmt.Errorf("percentage must be between 0 and 100")
		}

		client, err := createClient()
		if err != nil {
			return err
		}

		itemName, err := getItemName()
		if err != nil {
			return err
		}

		result, err := client.Publishers.Items.SetPublishedDeployPercentage(itemName).
			DeployPercentage(percentage).
			Do()
		if err != nil {
			return fmt.Errorf("failed to set deploy percentage: %w", err)
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
