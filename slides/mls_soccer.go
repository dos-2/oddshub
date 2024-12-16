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

// MLSSoccer creates a slide for soccer odds.
func MLSSoccer(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Soccer_usa_mls), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "MLS", GetHeader(models.Soccer_usa_mls), CreateH2HTable(pages, "MLS Soccer", headerString, games)
}
