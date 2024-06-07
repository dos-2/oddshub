package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

func UEFASoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "UEFA", GetHeader(models.Soccer_uefa_europa_league), CreateH2HTable(string(models.Soccer_uefa_europa_league) , tableData)
}
