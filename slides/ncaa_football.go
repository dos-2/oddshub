package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

func NCAAFootball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "NCAA Football", GetHeader(models.Americanfootball_ncaaf), CreateH2HTable("NCAA Football", tableData)
}
