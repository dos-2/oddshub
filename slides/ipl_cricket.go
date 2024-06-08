package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

// IPLCricket creates a slide for IPL cricket odds.
func IPLCricket(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "IPL", GetHeader(models.Cricket_ipl), CreateH2HTable(string(models.Cricket_ipl), builder.String())
}
