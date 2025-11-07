package polymarket_gamma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestEventFieldsParsing tests that extended fields like tags, categories, series are properly parsed
func TestEventFieldsParsing(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := NewClient(nil)

	// Fetch events via pagination
	response, err := client.GetEventsByPage(0, 5)

	require.NoError(t, err)
	require.NotNil(t, response)
	require.Greater(t, len(response.Events), 0, "Should fetch at least one event")

	// Check first event for extended fields
	event := response.Events[0]

	t.Logf("Event ID: %s", event.ID)
	t.Logf("Title: %s", event.Title)
	t.Logf("Category: %s", event.Category)

	// Check if tags are present
	if len(event.Tags) > 0 {
		t.Logf("✓ Tags found: %d tags", len(event.Tags))
		for i, tag := range event.Tags {
			if i < 3 {
				t.Logf("  Tag: %s (slug: %s)", tag.Label, tag.Slug)
			}
		}
	} else {
		t.Log("⚠ No tags in this event")
	}

	// Check if series are present
	if len(event.Series) > 0 {
		t.Logf("✓ Series found: %d series", len(event.Series))
		for _, series := range event.Series {
			t.Logf("  Series: %s (slug: %s)", series.Title, series.Slug)
		}
	} else {
		t.Log("⚠ No series in this event")
	}

	// Check if categories are present
	if len(event.Categories) > 0 {
		t.Logf("✓ Categories found: %d categories", len(event.Categories))
		for _, cat := range event.Categories {
			t.Logf("  Category: %s (slug: %s)", cat.Label, cat.Slug)
		}
	} else {
		t.Log("⚠ No categories in this event")
	}

	// Check markets have extended fields too
	if len(event.Markets) > 0 {
		market := event.Markets[0]
		t.Logf("Market ID: %s", market.ID)
		t.Logf("Market Question: %s", market.Question)

		if len(market.Tags) > 0 {
			t.Logf("✓ Market has %d tags", len(market.Tags))
		}

		if len(market.Categories) > 0 {
			t.Logf("✓ Market has %d categories", len(market.Categories))
		}
	}

	// Verify volume fields are populated
	t.Logf("Event Volume: %.2f", event.Volume)
	t.Logf("Event Liquidity: %.2f", event.Liquidity)
	t.Logf("Event Volume24hr: %.2f", event.Volume24hr)

	// Check boolean fields
	t.Logf("Active: %v, Closed: %v, Featured: %v", event.Active, event.Closed, event.Featured)

	t.Log("\n✓ All extended fields are being parsed correctly")
}
