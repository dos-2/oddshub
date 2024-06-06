package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// FrenchOpenTennis creates a slide for tennis odds.
func MensFrenchOpenTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Players|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "Mens French Open Tennis", GetHeader(models.Tennis_atp_french_open), CreateH2HTable("Mens French Open Tennis", tableData)
}
