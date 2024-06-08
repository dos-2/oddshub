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
		{Name: string(sports.Americanfootball_nfl), Content: NFLfootball},
		{Name: string(sports.Basketball_nba), Content: NBABasketball},
		{Name: string(sports.Baseball_mlb), Content: MLBBaseball},
		{Name: string(sports.Icehockey_nhl), Content: NHLHockey},
		{Name: string(sports.Soccer_usa_mls), Content: MLSSoccer},
		{Name: string(sports.Americanfootball_ncaaf), Content: NCAAFootball},
		{Name: string(sports.Basketball_ncaa), Content: NCAABasketball},
		{Name: string(sports.Baseball_ncaa), Content: NCAABaseball},
		{Name: string(sports.Boxing), Content: Boxing},
		{Name: string(sports.Mma_mixed_martial_arts), Content: Mma},
		{Name: string(sports.Tennis_atp_french_open), Content: MensFrenchOpenTennis},
		{Name: string(sports.Tennis_wta_french_open), Content: WomensFrenchOpenTennis},
		{Name: string(sports.Golf_masters_tournament_winner), Content: MastersGolf},
    {Name: string(sports.Golf_pga_championship_winner), Content: PGAGolf},
    {Name: string(sports.Soccer_brazil_campeonato), Content: BrazilCampeonato},
    {Name: string(sports.Rugbyleague_nrl), Content: NRLRugby},
	}
}
