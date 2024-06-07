package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

func PGAGolf(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Outrights||" + "\n"
  if len(games) > 0 {
		tableData += FormatTournamentEvent(games[0]) 
  }
	return "PGA", GetHeader(models.Golf_pga_tournament_winner), CreateRoundRobinTable(string(models.Golf_pga_tournament_winner) , tableData)
}
