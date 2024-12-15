/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"strings"

	"github.com/dos-2/oddshub/models"

	"github.com/rivo/tview"
)

// MastersGolf creates a slide for golf odds.
func MastersGolf(games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Teams|Players|Bookmaker|Outrights||\n")

	if len(games) > 0 {
		builder.WriteString(FormatTournamentEvent(games[0]))
	}

	return "Masters", GetHeader(models.Golf_masters_tournament_winner), CreateRoundRobinTable(string(models.Golf_masters_tournament_winner), builder.String())
}
