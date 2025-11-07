package polymarket_gamma

import (
	"time"
)

// ImageOptimization represents optimized image metadata
type ImageOptimization struct {
	ID                       string  `json:"id"`
	ImageURLSource           string  `json:"imageUrlSource"`
	ImageURLOptimized        string  `json:"imageUrlOptimized"`
	ImageSizeKbSource        float64 `json:"imageSizeKbSource"`
	ImageSizeKbOptimized     float64 `json:"imageSizeKbOptimized"`
	ImageOptimizedComplete   bool    `json:"imageOptimizedComplete"`
	ImageOptimizedLastUpdate string  `json:"imageOptimizedLastUpdated"`
	RelID                    int     `json:"relID"`
	Field                    string  `json:"field"`
	Relname                  string  `json:"relname"`
}

// Tag represents a tag associated with an event or market
type Tag struct {
	ID          string    `json:"id"`
	Label       string    `json:"label"`
	Slug        string    `json:"slug"`
	ForceShow   bool      `json:"forceShow"`
	PublishedAt string    `json:"publishedAt"`
	CreatedBy   int       `json:"createdBy"`
	UpdatedBy   int       `json:"updatedBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ForceHide   bool      `json:"forceHide"`
	IsCarousel  bool      `json:"isCarousel"`
}

// Category represents a category for events or markets
type Category struct {
	ID             string    `json:"id"`
	Label          string    `json:"label"`
	ParentCategory string    `json:"parentCategory"`
	Slug           string    `json:"slug"`
	PublishedAt    string    `json:"publishedAt"`
	CreatedBy      string    `json:"createdBy"`
	UpdatedBy      string    `json:"updatedBy"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// Series represents a series that an event belongs to
type Series struct {
	ID              string     `json:"id"`
	Ticker          string     `json:"ticker"`
	Slug            string     `json:"slug"`
	Title           string     `json:"title"`
	Subtitle        string     `json:"subtitle"`
	SeriesType      string     `json:"seriesType"`
	Recurrence      string     `json:"recurrence"`
	Description     string     `json:"description"`
	Image           string     `json:"image"`
	Icon            string     `json:"icon"`
	Layout          string     `json:"layout"`
	Active          bool       `json:"active"`
	Closed          bool       `json:"closed"`
	Archived        bool       `json:"archived"`
	New             bool       `json:"new"`
	Featured        bool       `json:"featured"`
	Restricted      bool       `json:"restricted"`
	PublishedAt     string     `json:"publishedAt"`
	CreatedBy       string     `json:"createdBy"`
	UpdatedBy       string     `json:"updatedBy"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	CommentsEnabled bool       `json:"commentsEnabled"`
	Competitive     string     `json:"competitive"`
	Volume24hr      float64    `json:"volume24hr"`
	Volume          float64    `json:"volume"`
	Liquidity       float64    `json:"liquidity"`
	StartDate       time.Time  `json:"startDate"`
	CommentCount    int        `json:"commentCount"`
	Categories      []Category `json:"categories"`
	Tags            []Tag      `json:"tags"`
}

// EventCreator represents a creator of an event
type EventCreator struct {
	ID            string    `json:"id"`
	CreatorName   string    `json:"creatorName"`
	CreatorHandle string    `json:"creatorHandle"`
	CreatorURL    string    `json:"creatorUrl"`
	CreatorImage  string    `json:"creatorImage"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// Event represents a Polymarket event from the Gamma API
type Event struct {
	ID                     string             `json:"id" validate:"required"`
	Ticker                 string             `json:"ticker"`
	Slug                   string             `json:"slug"`
	Title                  string             `json:"title"`
	Subtitle               string             `json:"subtitle"`
	Description            string             `json:"description"`
	ResolutionSource       string             `json:"resolutionSource"`
	StartDate              time.Time          `json:"startDate"`
	CreationDate           time.Time          `json:"creationDate"`
	EndDate                time.Time          `json:"endDate"`
	ImageURL               string             `json:"image"`
	Icon                   string             `json:"icon"`
	Active                 bool               `json:"active"`
	Closed                 bool               `json:"closed"`
	Archived               bool               `json:"archived"`
	New                    bool               `json:"new"`
	Featured               bool               `json:"featured"`
	Restricted             bool               `json:"restricted"`
	Liquidity              float64            `json:"liquidity"`
	Volume                 float64            `json:"volume"`
	OpenInterest           float64            `json:"openInterest"`
	SortBy                 string             `json:"sortBy"`
	Category               string             `json:"category"`
	Subcategory            string             `json:"subcategory"`
	PublishedAt            string             `json:"published_at"`
	CreatedBy              string             `json:"createdBy"`
	UpdatedBy              string             `json:"updatedBy"`
	CreatedAt              time.Time          `json:"createdAt"`
	UpdatedAt              time.Time          `json:"updatedAt"`
	CommentsEnabled        bool               `json:"commentsEnabled"`
	Competitive            float64            `json:"competitive"`
	Volume24hr             float64            `json:"volume24hr"`
	Volume1wk              float64            `json:"volume1wk"`
	Volume1mo              float64            `json:"volume1mo"`
	Volume1yr              float64            `json:"volume1yr"`
	FeaturedImage          string             `json:"featuredImage"`
	EnableOrderBook        bool               `json:"enableOrderBook"`
	LiquidityAmm           float64            `json:"liquidityAmm"`
	LiquidityClob          float64            `json:"liquidityClob"`
	NegRisk                bool               `json:"negRisk"`
	NegRiskMarketID        string             `json:"negRiskMarketID"`
	CommentCount           int                `json:"commentCount"`
	ImageOptimized         *ImageOptimization `json:"imageOptimized"`
	IconOptimized          *ImageOptimization `json:"iconOptimized"`
	FeaturedImageOptimized *ImageOptimization `json:"featuredImageOptimized"`
	SubEvents              []string           `json:"subEvents"`
	Markets                []Market           `json:"markets"`
	Series                 []Series           `json:"series"`
	Categories             []Category         `json:"categories"`
	Tags                   []Tag              `json:"tags"`
	Cyom                   bool               `json:"cyom"`
	ClosedTime             time.Time          `json:"closedTime"`
	ShowAllOutcomes        bool               `json:"showAllOutcomes"`
	ShowMarketImages       bool               `json:"showMarketImages"`
	EnableNegRisk          bool               `json:"enableNegRisk"`
	SeriesSlug             string             `json:"seriesSlug"`
	Live                   bool               `json:"live"`
	Ended                  bool               `json:"ended"`
	EventCreators          []EventCreator     `json:"eventCreators"`
	// Allow extra fields by not using strict parsing
}

// Market represents a Polymarket market from the Gamma API
type Market struct {
	ID                 string             `json:"id" validate:"required"`
	Question           string             `json:"question"`
	ConditionID        string             `json:"conditionId"`
	Slug               string             `json:"slug"`
	ResolutionSource   string             `json:"resolutionSource"`
	EndDate            time.Time          `json:"endDate"`
	StartDate          time.Time          `json:"startDate"`
	Description        string             `json:"description"`
	Active             bool               `json:"active"`
	Closed             bool               `json:"closed"`
	Archived           bool               `json:"archived"`
	MarketType         string             `json:"marketType"`
	RewardsMinSize     float64            `json:"rewardsMinSize"`
	RewardsMaxSpread   float64            `json:"rewardsMaxSpread"`
	Outcomes           string             `json:"outcomes"`
	OutcomePrices      string             `json:"outcomePrices"`
	Volume             string             `json:"volume"`
	Liquidity          string             `json:"liquidity"`
	Category           string             `json:"category"`
	CreatedAt          time.Time          `json:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt"`
	CreatedBy          int                `json:"createdBy"`
	UpdatedBy          int                `json:"updatedBy"`
	MarketMakerAddress string             `json:"marketMakerAddress"`
	New                bool               `json:"new"`
	Featured           bool               `json:"featured"`
	Restricted         bool               `json:"restricted"`
	VolumeNum          float64            `json:"volumeNum"`
	LiquidityNum       float64            `json:"liquidityNum"`
	Volume24hr         float64            `json:"volume24hr"`
	Volume1wk          float64            `json:"volume1wk"`
	Volume1mo          float64            `json:"volume1mo"`
	Volume1yr          float64            `json:"volume1yr"`
	EnableOrderBook    bool               `json:"enableOrderBook"`
	ClobTokenIds       string             `json:"clobTokenIds"`
	Competitive        float64            `json:"competitive"`
	Spread             float64            `json:"spread"`
	LastTradePrice     float64            `json:"lastTradePrice"`
	BestBid            float64            `json:"bestBid"`
	BestAsk            float64            `json:"bestAsk"`
	ImageOptimized     *ImageOptimization `json:"imageOptimized"`
	IconOptimized      *ImageOptimization `json:"iconOptimized"`
	Categories         []Category         `json:"categories"`
	Tags               []Tag              `json:"tags"`
	CommentsEnabled    bool               `json:"commentsEnabled"`
	// Allow extra fields by not using strict parsing
}

// GetEventsResponse represents the response from the events endpoint
type GetEventsResponse struct {
	Events []Event `json:"events"`
}
