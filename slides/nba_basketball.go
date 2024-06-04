package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// NbaBasketball creates a slide for basketball odds.
func NBABasketball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
	return "NBA Basketball", GetHeader(models.Basketball_nba), CreateTable("NBA Basketball", tableData)
}
