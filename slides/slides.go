package slides

import (
	"github.com/rivo/tview"
)

// Slide is a function which returns the slide's main primitive and its title.
// It receives a "nextSlide" function which can be called to advance the
// presentation to the next slide.
type Slide func(nextSlide func()) (title string, content tview.Primitive)


func GetSlides() []Slide {
    return []Slide{
        Cover,
        NFL_football,
		//	NCAA_football,
		//  NBA_basketball,
		//  NCAA_basketball,
		//  MLB_baseball,
		//  NCAA_baseball,
		//  MMA,
		//  NHL_hockey,
		//  Masters_golf,
		//  French_open_tennis,
    }
}
