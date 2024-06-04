package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// FootballSlide creates a slide for football odds.
func NHLHockey(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) // check whats going on here
	}
	return "NHL Hockey", GetHeader(models.Icehockey_nhl), CreateTable("NHL Hockey", tableData)
}
