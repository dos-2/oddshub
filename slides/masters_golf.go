package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// MastersGolf creates a slide for golf odds.
func MastersGolf(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Players|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "Masters Golf", GetHeader(models.Golf_masters_tournament), CreateRoundRobinTable("Masters Golf", tableData)
}
