package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	fetchStatusCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
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

		status, err := client.Publishers.Items.FetchStatus(itemName).Do()
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
			fmt.Printf("State:   %s\n", status.State)
			fmt.Printf("Version: %s\n", status.Version)
			if status.DownloadURL != "" {
				fmt.Printf("Download URL: %s\n", status.DownloadURL)
			}
			if len(status.DetailedStatus) > 0 {
				fmt.Println("Detailed Status:")
				for _, detail := range status.DetailedStatus {
					fmt.Printf("  - %s: %s\n", detail.StatusCode, detail.Description)
				}
			}
		}

		return nil
	},
}
