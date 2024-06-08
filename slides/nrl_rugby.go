package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

func NRLRugby(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")
	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NRL", GetHeader(models.Rugby_irl), CreateH2HTable(string(models.Rugby_irl), tableData.String())
}
