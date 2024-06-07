package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// MastersGolf creates a slide for golf odds.
func MastersGolf(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Teams|Players|Bookmaker|Outrights||" + "\n"
  if len(games) > 0 {
		tableData += FormatTournamentEvent(games[0]) 
  }
	return "Masters Tournament", GetHeader(models.Golf_masters_tournament_winner), CreateRoundRobinTable(string(models.Golf_masters_tournament_winner), tableData)
}
