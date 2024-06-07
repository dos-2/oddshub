package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

func MensWimbledonTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "Wimbledon", GetHeader(models.Tennis_atp_wimbledon), CreateH2HTable(string(models.Tennis_atp_wimbledon) , tableData)
}
