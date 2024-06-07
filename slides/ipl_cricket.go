package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// FootballSlide creates a slide for football odds.
func IPLCricket(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "IPL", GetHeader(models.Cricket_ipl), CreateH2HTable(string(models.Cricket_ipl) , tableData)
}
