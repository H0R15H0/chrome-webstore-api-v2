// Example usage of the Chrome Web Store API client.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/H0R15H0/chrome-webstore-api-v2/chromewebstore"
)

func main() {
	// Get credentials from environment variables
	clientID := os.Getenv("CHROME_WEBSTORE_CLIENT_ID")
	clientSecret := os.Getenv("CHROME_WEBSTORE_CLIENT_SECRET")
	refreshToken := os.Getenv("CHROME_WEBSTORE_REFRESH_TOKEN")
	publisherID := os.Getenv("CHROME_WEBSTORE_PUBLISHER_ID")
	itemID := os.Getenv("CHROME_WEBSTORE_ITEM_ID")

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		log.Fatal("Please set CHROME_WEBSTORE_CLIENT_ID, CHROME_WEBSTORE_CLIENT_SECRET, and CHROME_WEBSTORE_REFRESH_TOKEN environment variables")
	}

	if publisherID == "" || itemID == "" {
		log.Fatal("Please set CHROME_WEBSTORE_PUBLISHER_ID and CHROME_WEBSTORE_ITEM_ID environment variables")
	}

	ctx := context.Background()

	// Create an authenticated client
	client := chromewebstore.NewClientFromCredentials(ctx, chromewebstore.AuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	})

	// Create item name
	itemName := chromewebstore.NewItemName(publisherID, itemID)

	// Example 1: Fetch item status
	fmt.Println("Fetching item status...")
	status, err := client.Publishers.Items.FetchStatus(itemName).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Failed to fetch status: %v", err)
	}
	fmt.Printf("Item: %s\n", status.Name)
	fmt.Printf("State: %s\n", status.State)
	fmt.Printf("Version: %s\n", status.Version)
	fmt.Println()

	// Example 2: Upload a new version (uncomment to use)
	/*
		fmt.Println("Uploading new version...")
		file, err := os.Open("extension.zip")
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()

		uploadResp, err := client.Media.Upload(itemName).Context(ctx).Media(file, "application/zip").Do()
		if err != nil {
			log.Fatalf("Failed to upload: %v", err)
		}
		fmt.Printf("Upload status: %s\n", uploadResp.StatusCode)
		fmt.Println()
	*/

	// Example 3: Publish the item (uncomment to use)
	/*
		fmt.Println("Publishing item...")
		publishResp, err := client.Publishers.Items.Publish(itemName).Context(ctx).Do()
		if err != nil {
			log.Fatalf("Failed to publish: %v", err)
		}
		fmt.Printf("Publish status: %s\n", publishResp.StatusCode)
		fmt.Println()
	*/

	// Example 4: Publish to trusted testers only (uncomment to use)
	/*
		fmt.Println("Publishing to trusted testers...")
		publishResp, err := client.Publishers.Items.Publish(itemName).
			Context(ctx).
			PublishTarget(chromewebstore.PublishTargetTrustedTesters).
			Do()
		if err != nil {
			log.Fatalf("Failed to publish: %v", err)
		}
		fmt.Printf("Publish status: %s\n", publishResp.StatusCode)
		fmt.Println()
	*/

	// Example 5: Set deploy percentage (uncomment to use)
	/*
		fmt.Println("Setting deploy percentage to 50%...")
		deployResp, err := client.Publishers.Items.SetPublishedDeployPercentage(itemName).
			Context(ctx).
			DeployPercentage(50).
			Do()
		if err != nil {
			log.Fatalf("Failed to set deploy percentage: %v", err)
		}
		fmt.Printf("Deploy percentage status: %s\n", deployResp.StatusCode)
		fmt.Println()
	*/

	// Example 6: Cancel submission (uncomment to use)
	/*
		fmt.Println("Canceling submission...")
		cancelResp, err := client.Publishers.Items.CancelSubmission(itemName).Context(ctx).Do()
		if err != nil {
			log.Fatalf("Failed to cancel submission: %v", err)
		}
		fmt.Printf("Cancel status: %s\n", cancelResp.StatusCode)
		fmt.Println()
	*/

	fmt.Println("Done!")
}
