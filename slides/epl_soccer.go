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

func EPLSoccer(games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "EPL Soccer", GetHeader(models.Soccer_epl), CreateH2HTable(string(models.Soccer_epl), builder.String())
}
