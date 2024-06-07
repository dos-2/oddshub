package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

func EPLSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "EPL Soccer", GetHeader(models.Soccer_epl), CreateH2HTable(string(models.Soccer_epl) , tableData)
}
