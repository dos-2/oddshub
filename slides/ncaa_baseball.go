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

// NcaaBaseball creates a slide for baseball odds.
func NCAABaseball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money –|Total\n")

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "NCAA Baseball", GetHeader(models.Baseball_ncaa), CreateH2HTable("NCAA Baseball", tableData.String())
}
