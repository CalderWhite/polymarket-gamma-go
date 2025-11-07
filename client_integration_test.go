package polymarket_gamma

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestGetRealEventsByPagination tests fetching events via pagination from the real API
func TestGetRealEventsByPagination(t *testing.T) {
	g := NewWithT(t)
	client := NewClient(nil)

	response, err := client.GetEventsByPage(0, 10)

	g.Expect(err).To(BeNil())
	g.Expect(response).ToNot(BeNil())
	g.Expect(response.Events).ToNot(BeEmpty())
}

// TestGetExistingEvents tests fetching events 2890 and 2891 with full field validation
func TestGetExistingEvents(t *testing.T) {
	g := NewWithT(t)
	client := NewClient(nil)

	response, err := client.GetEventsByIDs([]int{2890, 2891})

	g.Expect(err).To(BeNil())
	g.Expect(response).ToNot(BeNil())
	g.Expect(response.Events).To(HaveLen(2))

	// Find events by ID
	var event2890, event2891 *Event
	for i := range response.Events {
		if response.Events[i].ID == "2890" {
			event2890 = &response.Events[i]
		} else if response.Events[i].ID == "2891" {
			event2891 = &response.Events[i]
		}
	}

	g.Expect(event2890).ToNot(BeNil(), "Event 2890 should be present")
	g.Expect(event2891).ToNot(BeNil(), "Event 2891 should be present")

	// Validate Event 2890 - NBA Mavericks vs Grizzlies
	expectedDate2890, _ := time.Parse(time.RFC3339, "2021-12-04T00:00:00Z")

	g.Expect(event2890.ID).To(Equal("2890"))
	g.Expect(event2890.Ticker).To(Equal("nba-will-the-mavericks-beat-the-grizzlies-by-more-than-5pt5-points-in-their-december-4-matchup"))
	g.Expect(event2890.Slug).To(Equal("nba-will-the-mavericks-beat-the-grizzlies-by-more-than-5pt5-points-in-their-december-4-matchup"))
	g.Expect(event2890.Title).To(Equal("NBA: Will the Mavericks beat the Grizzlies by more than 5.5 points in their December 4 matchup?"))
	g.Expect(event2890.Description).To(ContainSubstring("Dallas Mavericks"))
	g.Expect(event2890.Description).To(ContainSubstring("Memphis Grizzlies"))
	g.Expect(event2890.ResolutionSource).To(Equal("https://www.nba.com/games"))
	g.Expect(event2890.StartDate).To(Equal(expectedDate2890))
	g.Expect(event2890.CreationDate).To(Equal(expectedDate2890))
	g.Expect(event2890.EndDate).To(Equal(expectedDate2890))
	g.Expect(event2890.Active).To(BeTrue())
	g.Expect(event2890.Closed).To(BeTrue())
	g.Expect(event2890.Archived).To(BeFalse())
	g.Expect(event2890.New).To(BeFalse())
	g.Expect(event2890.Featured).To(BeFalse())
	g.Expect(event2890.Restricted).To(BeFalse())
	g.Expect(event2890.Category).To(Equal("Sports"))
	g.Expect(event2890.Volume).To(BeNumerically("==", 1335.05))
	g.Expect(event2890.OpenInterest).To(BeNumerically("==", 0))
	g.Expect(event2890.CommentCount).To(BeNumerically("==", 8125))
	g.Expect(event2890.Markets).To(HaveLen(1))
	g.Expect(event2890.Markets[0].ID).To(Equal("239826"))
	g.Expect(event2890.Markets[0].Question).To(Equal("NBA: Will the Mavericks beat the Grizzlies by more than 5.5 points in their December 4 matchup?"))
	g.Expect(event2890.Markets[0].Category).To(Equal("Sports"))
	g.Expect(event2890.Markets[0].Active).To(BeTrue())
	g.Expect(event2890.Markets[0].Closed).To(BeTrue())
	g.Expect(event2890.Series).To(HaveLen(1))
	g.Expect(event2890.Series[0].ID).To(Equal("2"))
	g.Expect(event2890.Series[0].Ticker).To(Equal("nba"))
	g.Expect(event2890.Series[0].Title).To(Equal("NBA"))
	g.Expect(event2890.Tags).To(HaveLen(1))
	g.Expect(event2890.Tags[0].Label).To(Equal("All"))

	// Validate Event 2891 - NFL Falcons vs Panthers
	expectedDate2891, _ := time.Parse(time.RFC3339, "2021-10-31T00:00:00Z")

	g.Expect(event2891.ID).To(Equal("2891"))
	g.Expect(event2891.Ticker).To(Equal("nfl-will-the-falcons-beat-the-panthers-by-more-than-3pt5-points-in-their-october-31st-matchup"))
	g.Expect(event2891.Slug).To(Equal("nfl-will-the-falcons-beat-the-panthers-by-more-than-3pt5-points-in-their-october-31st-matchup"))
	g.Expect(event2891.Title).To(Equal("NFL: Will the Falcons beat the Panthers by more than 3.5 points in their October 31st matchup?"))
	g.Expect(event2891.Description).To(ContainSubstring("Atlanta Falcons"))
	g.Expect(event2891.Description).To(ContainSubstring("Carolina Panthers"))
	g.Expect(event2891.ResolutionSource).To(Equal("https://www.nfl.com/scores/"))
	g.Expect(event2891.StartDate).To(Equal(expectedDate2891))
	g.Expect(event2891.CreationDate).To(Equal(expectedDate2891))
	g.Expect(event2891.EndDate).To(Equal(expectedDate2891))
	g.Expect(event2891.Active).To(BeTrue())
	g.Expect(event2891.Closed).To(BeTrue())
	g.Expect(event2891.Archived).To(BeFalse())
	g.Expect(event2891.New).To(BeFalse())
	g.Expect(event2891.Featured).To(BeFalse())
	g.Expect(event2891.Restricted).To(BeFalse())
	g.Expect(event2891.Category).To(Equal("Sports"))
	g.Expect(event2891.Volume).To(BeNumerically("==", 5332.42))
	g.Expect(event2891.OpenInterest).To(BeNumerically("==", 0))
	g.Expect(event2891.CommentCount).To(BeNumerically("==", 3787))
	g.Expect(event2891.Markets).To(HaveLen(1))
	g.Expect(event2891.Markets[0].ID).To(Equal("239167"))
	g.Expect(event2891.Markets[0].Question).To(Equal("NFL: Will the Falcons beat the Panthers by more than 3.5 points in their October 31st matchup?"))
	g.Expect(event2891.Markets[0].Category).To(Equal("Sports"))
	g.Expect(event2891.Markets[0].Active).To(BeTrue())
	g.Expect(event2891.Markets[0].Closed).To(BeTrue())
	g.Expect(event2891.Series).To(HaveLen(1))
	g.Expect(event2891.Series[0].ID).To(Equal("1"))
	g.Expect(event2891.Series[0].Ticker).To(Equal("nfl"))
	g.Expect(event2891.Series[0].Title).To(Equal("NFL"))
	g.Expect(event2891.Tags).To(HaveLen(1))
	g.Expect(event2891.Tags[0].Label).To(Equal("All"))
}
