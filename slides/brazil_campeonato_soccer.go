package slides

import (
	"oddshub/models"
  "strings"
	"github.com/rivo/tview"
)

func BrazilCampeonato(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Brazil Campeonato", GetHeader(models.Soccer_brazil_campeonato), CreateH2HTable(string(models.Soccer_brazil_campeonato), builder.String())
}
