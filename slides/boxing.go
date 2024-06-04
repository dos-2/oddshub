package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// Boxing creates a slide for boxing odds.
func Boxing(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "Boxing", GetHeader(models.Boxing), CreateTable("Boxing", tableData)
}
