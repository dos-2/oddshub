package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// MastersGolf creates a slide for golf odds.
func MastersGolf(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "Masters Golf", GetHeader(models.Golf_masters_tournament), CreateTable("Masters Golf", tableData)
}
