package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

// MLSSoccer creates a slide for soccer odds.
func MLSSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "MLS", GetHeader(models.Soccer_usa_mls), CreateH2HTable("MLS Soccer", tableData.String())
}
