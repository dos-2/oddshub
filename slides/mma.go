package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// Mma creates a slide for MMA odds.
func Mma(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "MMA", GetHeader(models.Mma_mixed_martial_arts), CreateTable("MMA", tableData)
}
