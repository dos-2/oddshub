package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

func NHLHockey(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")
	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NHL", GetHeader(models.Icehockey_nhl), CreateH2HTable("NHL Hockey", tableData.String())
}
