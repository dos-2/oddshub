package slides

import (
	"oddshub/models"
  "strings"
	"github.com/rivo/tview"
)

// FrenchOpenTennis creates a slide for tennis odds.
func MensFrenchOpenTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Players|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Mens French Open", GetHeader(models.Tennis_atp_french_open), CreateH2HTable("Mens French Open Tennis", builder.String())
}
