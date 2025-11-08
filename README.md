# Polymarket Gamma API Go Client

A Go client for interacting with the Polymarket Gamma structure API.

# Notes

This package:

- Does not error when new fields are added
- Will validate known fields
- May become out of date (please make a PR!)
- Supports querying by event, which is Polymarket's recommended method for market & event discovery 


## Installation

```bash
go get github.com/CalderWhite/polymarket-gamma-go
```

## Usage

## Fetching events by Id

```go
package main

import (
    "fmt"
    "log"
    
    polymarket_gamma "github.com/polymarket-gamma-go"
)

func main() {
    // Create a client with default settings
    client := polymarket_gamma.NewClient(nil)
    
    // Fetch events by IDs (query uses integers, response contains string IDs)
    response, err := client.GetEventsByIDs([]int{2890, 2891, 2892})
    if err != nil {
        log.Fatal(err)
    }
    
    for _, event := range response.Events {
        fmt.Printf("Event: %s\n", event.Title)
        for _, market := range event.Markets {
            fmt.Printf("  Market: %s\n", market.Question)
        }
    }
}
```

## Pagination

This is primarily for event/market discovery.

```go
// Fetch the latest 10 events (by ID)
response, err := client.GetEventsByPage(0, 10, false)

// Fetch events in order
response, err := client.GetEventsByPage(0, 10, true)
```

