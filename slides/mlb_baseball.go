package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// MlbBaseball creates a slide for baseball odds.
func MLBBaseball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "MLB", GetHeader(models.Baseball_mlb), CreateH2HTable("MLB Baseball", tableData)
}
