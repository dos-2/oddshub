package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// WomensFrenchOpenTennis creates a slide for tennis odds.
func WomensFrenchOpenTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "Womens French Open Tennis", GetHeader(models.Tennis_wta_french_open), CreateTable("Womens French Open Tennis", tableData)
}
