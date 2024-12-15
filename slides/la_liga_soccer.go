/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"strings"

	"github.com/dos-2/oddshub/models"
	"github.com/dos-2/oddshub/sports"

	"github.com/rivo/tview"
)

// LaLigaSoccer creates a slide for La Liga soccer odds.
func LaLigaSoccer(games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n")
	models.AddLoadedEvent(string(sports.Soccer_spain_la_liga), games)

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "La Liga", GetHeader(models.Soccer_spain_la_liga), CreateH2HTable(string(models.Soccer_spain_la_liga), builder.String())
}
