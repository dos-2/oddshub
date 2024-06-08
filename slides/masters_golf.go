package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

// MastersGolf creates a slide for golf odds.
func MastersGolf(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Teams|Players|Bookmaker|Outrights||\n")

	if len(games) > 0 {
		builder.WriteString(FormatTournamentEvent(games[0]))
	}

	return "Masters Tournament", GetHeader(models.Golf_masters_tournament_winner), CreateRoundRobinTable(string(models.Golf_masters_tournament_winner), builder.String())
}
