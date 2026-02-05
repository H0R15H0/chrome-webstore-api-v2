package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var projection string

func init() {
	fetchStatusCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
	fetchStatusCmd.Flags().StringVar(&projection, "projection", "", "Projection type: DRAFT or PUBLISHED")
	rootCmd.AddCommand(fetchStatusCmd)
}

var fetchStatusCmd = &cobra.Command{
	Use:   "fetch-status",
	Short: "Fetch the status of an item",
	Long:  `Fetch the current status of a Chrome Web Store item.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createClient()
		if err != nil {
			return err
		}

		itemName, err := getItemName()
		if err != nil {
			return err
		}

		call := client.Publishers.Items.FetchStatus(itemName)
		if projection != "" {
			call = call.Projection(projection)
		}
		status, err := call.Do()
		if err != nil {
			return fmt.Errorf("failed to fetch status: %w", err)
		}

		if jsonOutput {
			output, err := json.MarshalIndent(status, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal JSON: %w", err)
			}
			fmt.Println(string(output))
		} else {
			fmt.Printf("Name:    %s\n", status.Name)
			fmt.Printf("Item ID: %s\n", status.ItemID)

			if status.SubmittedItemRevisionStatus != nil {
				fmt.Println("Submitted:")
				fmt.Printf("  State: %s\n", status.SubmittedItemRevisionStatus.State)
				for _, ch := range status.SubmittedItemRevisionStatus.DistributionChannels {
					fmt.Printf("  Version: %s (Deploy: %d%%)\n", ch.CrxVersion, ch.DeployPercentage)
				}
			}

			if status.PublishedItemRevisionStatus != nil {
				fmt.Println("Published:")
				fmt.Printf("  State: %s\n", status.PublishedItemRevisionStatus.State)
				for _, ch := range status.PublishedItemRevisionStatus.DistributionChannels {
					fmt.Printf("  Version: %s (Deploy: %d%%)\n", ch.CrxVersion, ch.DeployPercentage)
				}
			}
		}

		return nil
	},
}
