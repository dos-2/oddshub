package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

func BrazilCampeonato(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "Brazil Campeonato", GetHeader(models.Soccer_brazil_campeonato), CreateH2HTable(string(models.Soccer_brazil_campeonato) , tableData)
}
