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

// IPLCricket creates a slide for IPL cricket odds.
func IPLCricket(games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "IPL", GetHeader(models.Cricket_ipl), CreateH2HTable(string(models.Cricket_ipl), builder.String())
}
