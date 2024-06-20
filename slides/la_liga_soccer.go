/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

// LaLigaSoccer creates a slide for La Liga soccer odds.
func LaLigaSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "La Liga", GetHeader(models.Soccer_spain_la_liga), CreateH2HTable(string(models.Soccer_spain_la_liga), builder.String())
}
