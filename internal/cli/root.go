package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/H0R15H0/chrome-webstore-api-v2/chromewebstore"
	"github.com/spf13/cobra"
)

var (
	publisherID string
	itemID      string
	jsonOutput  bool
)

var rootCmd = &cobra.Command{
	Use:          "cws",
	Short:        "Chrome Web Store API CLI",
	Long:         `A command-line interface for the Chrome Web Store API v2.`,
	SilenceUsage: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&publisherID, "publisher-id", "", "Publisher ID (or set CHROME_WEBSTORE_PUBLISHER_ID)")
	rootCmd.PersistentFlags().StringVar(&itemID, "item-id", "", "Item ID (or set CHROME_WEBSTORE_ITEM_ID)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func getPublisherID() string {
	if publisherID != "" {
		return publisherID
	}
	return os.Getenv("CHROME_WEBSTORE_PUBLISHER_ID")
}

func getItemID() string {
	if itemID != "" {
		return itemID
	}
	return os.Getenv("CHROME_WEBSTORE_ITEM_ID")
}

func createClient() (*chromewebstore.Client, error) {
	clientID := os.Getenv("CHROME_WEBSTORE_CLIENT_ID")
	clientSecret := os.Getenv("CHROME_WEBSTORE_CLIENT_SECRET")
	refreshToken := os.Getenv("CHROME_WEBSTORE_REFRESH_TOKEN")

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		return nil, fmt.Errorf("missing required environment variables: CHROME_WEBSTORE_CLIENT_ID, CHROME_WEBSTORE_CLIENT_SECRET, CHROME_WEBSTORE_REFRESH_TOKEN")
	}

	config := chromewebstore.AuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}

	return chromewebstore.NewClientFromCredentials(context.Background(), config), nil
}

func getItemName() (chromewebstore.ItemName, error) {
	pubID := getPublisherID()
	itmID := getItemID()

	if pubID == "" {
		return "", fmt.Errorf("publisher-id is required (use --publisher-id flag or CHROME_WEBSTORE_PUBLISHER_ID environment variable)")
	}
	if itmID == "" {
		return "", fmt.Errorf("item-id is required (use --item-id flag or CHROME_WEBSTORE_ITEM_ID environment variable)")
	}

	return chromewebstore.NewItemName(pubID, itmID), nil
}
