package polymarket_gamma

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockEvent(id string) Event {
	return Event{
		ID:           id,
		Slug:         "test-event",
		Title:        "Test Event",
		Subtitle:     "Test Subtitle",
		Description:  "Test Description",
		Category:     "Sports",
		Subcategory:  "Basketball",
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(24 * time.Hour),
		Active:       true,
		Closed:       false,
		Archived:     false,
		Featured:     false,
		New:          false,
		Volume:       12345.67,
		Liquidity:    5000.00,
		Volume24hr:   100.50,
		CommentCount: 42,
		Tags: []Tag{
			{
				ID:        "tag-1",
				Label:     "Test Tag",
				Slug:      "test-tag",
				ForceShow: true,
			},
		},
		Categories: []Category{
			{
				ID:    "cat-1",
				Label: "Sports",
				Slug:  "sports",
			},
		},
		Markets: []Market{
			mockMarket("market-1"),
		},
	}
}

func mockMarket(id string) Market {
	return Market{
		ID:            id,
		Question:      "Will this happen?",
		ConditionID:   "condition-1",
		Slug:          "test-market",
		Active:        true,
		Closed:        false,
		Archived:      false,
		Featured:      false,
		MarketType:    "binary",
		Outcomes:      `["Yes", "No"]`,
		OutcomePrices: `["0.5", "0.5"]`,
		Volume:        "10000",
		Liquidity:     "5000",
		VolumeNum:     10000.0,
		LiquidityNum:  5000.0,
		Volume24hr:    250.0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Tags: []Tag{
			{
				ID:    "tag-market-1",
				Label: "Market Tag",
				Slug:  "market-tag",
			},
		},
	}
}

func TestGetEventsByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request
		assert.Equal(t, "/events", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Check query parameters
		ids := r.URL.Query()["id"]
		assert.Contains(t, ids, "1")
		assert.Contains(t, ids, "2")

		// Return mock events (API returns string IDs)
		events := []Event{
			mockEvent("1"),
			mockEvent("2"),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}))
	defer server.Close()

	// Create client with mock server
	client := NewClient(&ClientConfig{
		BaseURL: server.URL,
	})

	// Test fetching events by ID
	response, err := client.GetEventsByIDs([]int{1, 2})

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Len(t, response.Events, 2)
	assert.Equal(t, "1", response.Events[0].ID)
	assert.Equal(t, "2", response.Events[1].ID)
	assert.Equal(t, "Test Event", response.Events[0].Title)

	// Verify new fields
	event := response.Events[0]
	assert.Equal(t, "Test Subtitle", event.Subtitle)
	assert.Equal(t, "Sports", event.Category)
	assert.Equal(t, "Basketball", event.Subcategory)
	assert.Equal(t, 12345.67, event.Volume)
	assert.Equal(t, 5000.00, event.Liquidity)
	assert.Equal(t, 100.50, event.Volume24hr)
	assert.Equal(t, 42, event.CommentCount)

	// Verify tags
	assert.Len(t, event.Tags, 1)
	assert.Equal(t, "tag-1", event.Tags[0].ID)
	assert.Equal(t, "Test Tag", event.Tags[0].Label)
	assert.Equal(t, "test-tag", event.Tags[0].Slug)

	// Verify categories
	assert.Len(t, event.Categories, 1)
	assert.Equal(t, "cat-1", event.Categories[0].ID)
	assert.Equal(t, "Sports", event.Categories[0].Label)

	// Verify markets
	assert.Len(t, event.Markets, 1)
	assert.Equal(t, "market-1", event.Markets[0].ID)

	// Verify market fields
	market := event.Markets[0]
	assert.Equal(t, 10000.0, market.VolumeNum)
	assert.Equal(t, 5000.0, market.LiquidityNum)
	assert.Equal(t, 250.0, market.Volume24hr)
	assert.Len(t, market.Tags, 1)
	assert.Equal(t, "Market Tag", market.Tags[0].Label)
}

func TestGetLatestEvents(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request
		assert.Equal(t, "/events", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Check query parameters
		query := r.URL.Query()
		assert.Equal(t, "0", query.Get("offset"))
		assert.Equal(t, "10", query.Get("limit"))
		assert.Equal(t, "true", query.Get("ascending"))
		assert.Equal(t, "id", query.Get("sortBy"))

		// Return mock events (API returns string IDs)
		events := []Event{
			mockEvent("1"),
			mockEvent("2"),
			mockEvent("3"),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}))
	defer server.Close()

	// Create client with mock server
	client := NewClient(&ClientConfig{
		BaseURL: server.URL,
	})

	// Test fetching latest events with pagination
	response, err := client.GetEventsByPage(0, 10)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Len(t, response.Events, 3)
	assert.Equal(t, "1", response.Events[0].ID)
	assert.Equal(t, "2", response.Events[1].ID)
	assert.Equal(t, "3", response.Events[2].ID)
}

func TestGetEventsWithExtraFields(t *testing.T) {
	// Create a mock server that returns events with extra fields
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return mock events with extra fields including tags, categories, etc.
		response := `[
			{
				"id": "1",
				"slug": "test-event",
				"title": "Test Event",
				"subtitle": "Test Subtitle",
				"description": "Test Description",
				"category": "Sports",
				"subcategory": "Basketball",
				"startDate": "2025-01-01T00:00:00Z",
				"endDate": "2025-01-02T00:00:00Z",
				"active": true,
				"closed": false,
				"archived": false,
				"featured": false,
				"volume": 12345.67,
				"liquidity": 5000.0,
				"volume24hr": 100.5,
				"commentCount": 42,
				"newFieldThatDidntExistBefore": "should not cause errors",
				"anotherExtraField": 12345,
				"tags": [
					{
						"id": "tag-1",
						"label": "Test Tag",
						"slug": "test-tag",
						"forceShow": true,
						"createdAt": "2025-01-01T00:00:00Z",
						"updatedAt": "2025-01-01T00:00:00Z"
					}
				],
				"categories": [
					{
						"id": "cat-1",
						"label": "Sports",
						"slug": "sports",
						"createdAt": "2025-01-01T00:00:00Z",
						"updatedAt": "2025-01-01T00:00:00Z"
					}
				],
				"series": [
					{
						"id": "series-1",
						"title": "Test Series",
						"slug": "test-series",
						"createdAt": "2025-01-01T00:00:00Z",
						"updatedAt": "2025-01-01T00:00:00Z"
					}
				],
				"markets": [
					{
						"id": "market-1",
						"question": "Will this happen?",
						"conditionId": "condition-1",
						"slug": "test-market",
						"active": true,
						"closed": false,
						"archived": false,
						"marketType": "binary",
						"outcomes": "[\"Yes\", \"No\"]",
						"outcomePrices": "[\"0.5\", \"0.5\"]",
						"volume": "10000",
						"liquidity": "5000",
						"volumeNum": 10000.0,
						"liquidityNum": 5000.0,
						"volume24hr": 250.0,
						"createdAt": "2025-01-01T00:00:00Z",
						"updatedAt": "2025-01-01T00:00:00Z",
						"tags": [
							{
								"id": "market-tag-1",
								"label": "Market Tag",
								"slug": "market-tag"
							}
						],
						"extraMarketField": "should also not cause errors"
					}
				]
			}
		]`

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with mock server
	client := NewClient(&ClientConfig{
		BaseURL: server.URL,
	})

	// Test fetching events - should not fail despite extra fields
	response, err := client.GetEventsByIDs([]int{1})

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Len(t, response.Events, 1)

	event := response.Events[0]
	assert.Equal(t, "1", event.ID)
	assert.Equal(t, "Test Event", event.Title)
	assert.Equal(t, "Test Subtitle", event.Subtitle)
	assert.Equal(t, "Sports", event.Category)
	assert.Equal(t, "Basketball", event.Subcategory)

	// Verify new fields are parsed
	assert.Equal(t, 12345.67, event.Volume)
	assert.Equal(t, 5000.0, event.Liquidity)
	assert.Equal(t, 100.5, event.Volume24hr)
	assert.Equal(t, 42, event.CommentCount)

	// Verify tags are parsed
	assert.Len(t, event.Tags, 1)
	assert.Equal(t, "tag-1", event.Tags[0].ID)
	assert.Equal(t, "Test Tag", event.Tags[0].Label)
	assert.Equal(t, "test-tag", event.Tags[0].Slug)
	assert.True(t, event.Tags[0].ForceShow)

	// Verify categories are parsed
	assert.Len(t, event.Categories, 1)
	assert.Equal(t, "cat-1", event.Categories[0].ID)
	assert.Equal(t, "Sports", event.Categories[0].Label)

	// Verify series are parsed
	assert.Len(t, event.Series, 1)
	assert.Equal(t, "series-1", event.Series[0].ID)
	assert.Equal(t, "Test Series", event.Series[0].Title)

	// Verify markets
	assert.Len(t, event.Markets, 1)
	market := event.Markets[0]
	assert.Equal(t, "market-1", market.ID)
	assert.Equal(t, 10000.0, market.VolumeNum)
	assert.Equal(t, 5000.0, market.LiquidityNum)
	assert.Equal(t, 250.0, market.Volume24hr)

	// Verify market tags
	assert.Len(t, market.Tags, 1)
	assert.Equal(t, "market-tag-1", market.Tags[0].ID)
	assert.Equal(t, "Market Tag", market.Tags[0].Label)
}

func TestGetEventsAPIError(t *testing.T) {
	// Create a mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	// Create client with mock server
	client := NewClient(&ClientConfig{
		BaseURL: server.URL,
	})

	// Test fetching events - should return error
	response, err := client.GetEventsByIDs([]int{1})

	require.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "failed to fetch events")
	assert.Contains(t, err.Error(), "500")
}

func TestGetEventsValidationError(t *testing.T) {
	// Create a mock server that returns invalid data (missing required field)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `[
			{
				"slug": "test-event",
				"title": "Test Event"
			}
		]`

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with mock server
	client := NewClient(&ClientConfig{
		BaseURL: server.URL,
	})

	// Test fetching events - should return validation error
	response, err := client.GetEventsByIDs([]int{1})

	require.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "validation failed")
}

func TestNestedTypes(t *testing.T) {
	// Create a mock server that returns events with nested types like ImageOptimization
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `[
			{
				"id": "1",
				"slug": "test-event",
				"title": "Test Event",
				"description": "Test Description",
				"startDate": "2025-01-01T00:00:00Z",
				"endDate": "2025-01-02T00:00:00Z",
				"active": true,
				"closed": false,
				"archived": false,
				"imageOptimized": {
					"id": "img-1",
					"imageUrlSource": "https://example.com/image.png",
					"imageUrlOptimized": "https://example.com/image-optimized.png",
					"imageSizeKbSource": 500.5,
					"imageSizeKbOptimized": 150.2,
					"imageOptimizedComplete": true,
					"imageOptimizedLastUpdated": "2025-01-01T00:00:00Z",
					"relID": 1,
					"field": "image",
					"relname": "events"
				},
				"iconOptimized": {
					"id": "icon-1",
					"imageUrlSource": "https://example.com/icon.png",
					"imageUrlOptimized": "https://example.com/icon-optimized.png",
					"imageSizeKbSource": 100.0,
					"imageSizeKbOptimized": 25.0,
					"imageOptimizedComplete": true
				},
				"eventCreators": [
					{
						"id": "creator-1",
						"creatorName": "John Doe",
						"creatorHandle": "@johndoe",
						"creatorUrl": "https://twitter.com/johndoe",
						"creatorImage": "https://example.com/creator.png",
						"createdAt": "2025-01-01T00:00:00Z",
						"updatedAt": "2025-01-01T00:00:00Z"
					}
				],
				"markets": [
					{
						"id": "market-1",
						"question": "Test?",
						"conditionId": "condition-1",
						"slug": "test",
						"active": true,
						"closed": false,
						"archived": false,
						"marketType": "binary",
						"outcomes": "[\"Yes\", \"No\"]",
						"outcomePrices": "[\"0.5\", \"0.5\"]",
						"volume": "1000",
						"liquidity": "500",
						"createdAt": "2025-01-01T00:00:00Z",
						"updatedAt": "2025-01-01T00:00:00Z",
						"imageOptimized": {
							"id": "market-img-1",
							"imageUrlSource": "https://example.com/market.png",
							"imageUrlOptimized": "https://example.com/market-optimized.png",
							"imageSizeKbSource": 200.0,
							"imageSizeKbOptimized": 50.0
						}
					}
				]
			}
		]`

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client with mock server
	client := NewClient(&ClientConfig{
		BaseURL: server.URL,
	})

	// Test fetching events
	response, err := client.GetEventsByIDs([]int{1})

	require.NoError(t, err)
	require.NotNil(t, response)
	require.Len(t, response.Events, 1)

	event := response.Events[0]

	// Verify ImageOptimization for event
	require.NotNil(t, event.ImageOptimized)
	assert.Equal(t, "img-1", event.ImageOptimized.ID)
	assert.Equal(t, "https://example.com/image.png", event.ImageOptimized.ImageURLSource)
	assert.Equal(t, "https://example.com/image-optimized.png", event.ImageOptimized.ImageURLOptimized)
	assert.Equal(t, 500.5, event.ImageOptimized.ImageSizeKbSource)
	assert.Equal(t, 150.2, event.ImageOptimized.ImageSizeKbOptimized)
	assert.True(t, event.ImageOptimized.ImageOptimizedComplete)
	assert.Equal(t, 1, event.ImageOptimized.RelID)
	assert.Equal(t, "image", event.ImageOptimized.Field)
	assert.Equal(t, "events", event.ImageOptimized.Relname)

	// Verify IconOptimized for event
	require.NotNil(t, event.IconOptimized)
	assert.Equal(t, "icon-1", event.IconOptimized.ID)
	assert.Equal(t, 100.0, event.IconOptimized.ImageSizeKbSource)

	// Verify EventCreators
	require.Len(t, event.EventCreators, 1)
	assert.Equal(t, "creator-1", event.EventCreators[0].ID)
	assert.Equal(t, "John Doe", event.EventCreators[0].CreatorName)
	assert.Equal(t, "@johndoe", event.EventCreators[0].CreatorHandle)
	assert.Equal(t, "https://twitter.com/johndoe", event.EventCreators[0].CreatorURL)

	// Verify market ImageOptimization
	require.Len(t, event.Markets, 1)
	market := event.Markets[0]
	require.NotNil(t, market.ImageOptimized)
	assert.Equal(t, "market-img-1", market.ImageOptimized.ID)
	assert.Equal(t, 200.0, market.ImageOptimized.ImageSizeKbSource)
	assert.Equal(t, 50.0, market.ImageOptimized.ImageSizeKbOptimized)
}
