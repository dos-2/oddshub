package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// FootballSlide creates a slide for football odds.
func LaLigaSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "La Liga", GetHeader(models.Soccer_spain_la_liga), CreateH2HTable(string(models.Soccer_spain_la_liga) , tableData)
}
