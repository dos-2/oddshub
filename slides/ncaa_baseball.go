package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// NcaaBaseball creates a slide for baseball odds.
func NCAABaseball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "NCAA Baseball", GetHeader(models.Baseball_ncaa), CreateTable("NCAA Baseball", tableData)
}
