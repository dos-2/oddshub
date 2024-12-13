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

// LaLigaSoccer creates a slide for La Liga soccer odds.
func LaLigaSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money –|Total\n")

	eventsSorted := sortEvents(games, "money", "desc")

	for _, game := range eventsSorted {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "La Liga", GetHeader(models.Soccer_spain_la_liga), CreateH2HTable(string(models.Soccer_spain_la_liga), builder.String())
}
