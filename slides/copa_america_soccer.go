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

func CopaAmericaSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Copa America", GetHeader(models.Soccer_conmebol_copa_america), CreateH2HTable(string(models.Soccer_conmebol_copa_america), builder.String())
}
