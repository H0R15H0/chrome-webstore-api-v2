package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uploadCmd)
}

var uploadCmd = &cobra.Command{
	Use:   "upload <file.zip>",
	Short: "Upload an extension package",
	Long:  `Upload a ZIP file containing the extension package to Chrome Web Store.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]

		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		client, err := createClient()
		if err != nil {
			return err
		}

		itemName, err := getItemName()
		if err != nil {
			return err
		}

		result, err := client.Media.Upload(itemName).Media(file, "application/zip").Do()
		if err != nil {
			return fmt.Errorf("failed to upload: %w", err)
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
			fmt.Printf("Status:  %s\n", result.UploadState)
			if result.CrxVersion != "" {
				fmt.Printf("Version: %s\n", result.CrxVersion)
			}
		}

		return nil
	},
}
