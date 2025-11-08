package main

import (
	"fmt"
	"log"

	polymarket_gamma "github.com/CalderWhite/polymarket-gamma-go"
)

func main() {
	// Example 1: Basic client with default settings
	fmt.Println("=== Example 1: Fetch events by IDs ===")
	client := polymarket_gamma.NewClient(nil)

	response, err := client.GetEventsByIDs([]int{2890, 2891})
	if err != nil {
		log.Printf("Error fetching events: %v\n", err)
	} else {
		fmt.Printf("Fetched %d events\n", len(response.Events))
		for _, event := range response.Events {
			fmt.Printf("  Event #%s: %s\n", event.ID, event.Title)
			fmt.Printf("    Markets: %d\n", len(event.Markets))
		}
	}

	// Example 2: Fetch with pagination
	fmt.Println("\n=== Example 2: Fetch latest events with pagination ===")
	response, err = client.GetEventsByPage(0, 5, false)
	if err != nil {
		log.Printf("Error fetching events: %v\n", err)
	} else {
		fmt.Printf("Fetched %d events\n", len(response.Events))
		for _, event := range response.Events {
			fmt.Printf("  Event #%s: %s\n", event.ID, event.Title)
		}
	}

}
