package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// NcaaBaseball creates a slide for baseball odds.
func NCAABaseball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "NCAA Baseball", GetHeader(models.Baseball_ncaa), CreateH2HTable("NCAA Baseball", tableData)
}
