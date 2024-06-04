package slides

import (
	"oddshub/models"
  "oddshub/sports"
	"github.com/rivo/tview"
)

// Slide is a struct representing a slide in the presentation.
type Slide struct {
	Name    string // Name of the slide
	Content func(games []models.Event, nextSlide func()) (title string, header string, content tview.Primitive)
}

// GetSlides returns a slice of slides for the presentation.
func GetSlides() []Slide {
	return []Slide{
		{Name: "Cover", Content: Cover},
		{Name: string(sports.Americanfootball_nfl), Content: NFL_football},
		// Add other slides here
	}
}
