package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// Boxing creates a slide for boxing odds.
func Boxing(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Ranking|Players|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "Boxing", GetHeader(models.Boxing), CreateH2HTable("Boxing", tableData)
}
