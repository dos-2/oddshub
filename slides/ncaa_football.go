package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// FootballSlide creates a slide for football odds.
func NCAAFootball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) // check whats going on here
	}
	return "NCAA Football", GetHeader(models.Americanfootball_ncaaf), CreateH2HTable("NCAA Football", tableData)
}
