package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

func UEFASoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")
	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "UEFA", GetHeader(models.Soccer_uefa_europa_league), CreateH2HTable(string(models.Soccer_uefa_europa_league), tableData.String())
}
