// Package polymarket_gamma provides a Go client for interacting with the Polymarket Gamma structure API.
//
// This package:
//   - Does not error when new fields are added
//   - Will validate known fields
//   - May become out of date (please make a PR!)
//   - Supports querying by event, which is Polymarket's recommended method for market & event discovery
//
// # Installation
//
//	go get github.com/CalderWhite/polymarket-gamma-go
//
// # Usage
//
// Fetching events by ID:
//
//	client := polymarket_gamma.NewClient(nil)
//	response, err := client.GetEventsByIDs([]int{2890, 2891, 2892})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, event := range response.Events {
//	    fmt.Printf("Event: %s\n", event.Title)
//	    for _, market := range event.Markets {
//	        fmt.Printf("  Market: %s\n", market.Question)
//	    }
//	}
//
// Pagination (for event/market discovery):
//
//	// Fetch the latest 10 events (by ID)
//	response, err := client.GetEventsByPage(0, 10, false)
//
//	// Fetch events in order
//	response, err := client.GetEventsByPage(0, 10, true)
package polymarket_gamma

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/go-playground/validator/v10"
)

const (
	defaultBaseURL = "https://gamma-api.polymarket.com"
	defaultTimeout = 30 * time.Second
)

type ClientConfig struct {
	BaseURL string
	Timeout time.Duration
	// Custom transport (optional)
	Transport http.RoundTripper
	// Custom HTTP client (optional)
	HTTPClient *http.Client
}

// Polymarket Gamma API client
type Client struct {
	baseURL    string
	httpClient *http.Client
	validator  *validator.Validate
}

func NewClient(config *ClientConfig) *Client {
	if config == nil {
		config = &ClientConfig{}
	}

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	timeout := config.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}

	var httpClient *http.Client
	if config.HTTPClient != nil {
		httpClient = config.HTTPClient
	} else {
		httpClient = &http.Client{
			Timeout: timeout,
		}

		if config.Transport != nil {
			httpClient.Transport = config.Transport
		}
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		validator:  validator.New(),
	}
}

// GetEventsByIDs fetches events by their IDs from the Polymarket Gamma API
func (c *Client) GetEventsByIDs(ids []int) (*GetEventsResponse, error) {
	queryParams := url.Values{}

	// Add multiple id parameters (API expects integers)
	for _, id := range ids {
		queryParams.Add("id", strconv.Itoa(id))
	}

	return c.getEvents(queryParams)
}

// GetEventsByPage fetches events with pagination from the Polymarket Gamma API
func (c *Client) GetEventsByPage(offset, limit int, ascending bool) (*GetEventsResponse, error) {
	queryParams := url.Values{}
	queryParams.Set("offset", strconv.Itoa(offset))
	queryParams.Set("limit", strconv.Itoa(limit))
	queryParams.Set("ascending", strconv.FormatBool(ascending))
	queryParams.Set("order", "id")

	return c.getEvents(queryParams)
}

func (c *Client) GetActiveEventsByPage(offset, limit int, ascending bool) (*GetEventsResponse, error) {
	queryParams := url.Values{}
	queryParams.Set("offset", strconv.Itoa(offset))
	queryParams.Set("limit", strconv.Itoa(limit))
	queryParams.Set("ascending", strconv.FormatBool(ascending))
	queryParams.Set("order", "id")
	queryParams.Set("closed", "false") // polymarket doesn't seem to use the `active` column

	return c.getEvents(queryParams)
}

// getEvents is the private implementation that fetches events from the Polymarket Gamma API
func (c *Client) getEvents(queryParams url.Values) (*GetEventsResponse, error) {

	// Build URL
	apiURL := fmt.Sprintf("%s/events", c.baseURL)
	if len(queryParams) > 0 {
		apiURL = fmt.Sprintf("%s?%s", apiURL, queryParams.Encode())
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Accept gzip encoding to reduce bandwidth
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch events: %d %s - %s", resp.StatusCode, resp.Status, string(body))
	}

	// Handle gzip decompression if needed
	var reader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var events []Event
	if err := sonic.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	for i, event := range events {
		// Validate event (skipMissingProperties and whitelist:false equivalent)
		if err := c.validator.Struct(event); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				return nil, fmt.Errorf("validation failed for event %d: %v", i, validationErrs)
			}
			return nil, fmt.Errorf("validation failed for event %d: %w", i, err)
		}

		// Validate markets
		for j, market := range event.Markets {
			if err := c.validator.Struct(market); err != nil {
				if validationErrs, ok := err.(validator.ValidationErrors); ok {
					return nil, fmt.Errorf("validation failed for market %d in event %d: %v", j, i, validationErrs)
				}
				return nil, fmt.Errorf("validation failed for market %d in event %d: %w", j, i, err)
			}
		}
	}

	return &GetEventsResponse{
		Events: events,
	}, nil
}
