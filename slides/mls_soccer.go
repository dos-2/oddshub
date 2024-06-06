package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// MLSSoccer creates a slide for soccer odds.
func MLSSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "MLS", GetHeader(models.Soccer_usa_mls), CreateH2HTable("MLS Soccer", tableData)
}
