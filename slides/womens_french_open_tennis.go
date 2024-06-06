package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// WomensFrenchOpenTennis creates a slide for tennis odds.
func WomensFrenchOpenTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Players|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "Womens French Open", GetHeader(models.Tennis_wta_french_open), CreateH2HTable("Womens French Open Tennis", tableData)
}
