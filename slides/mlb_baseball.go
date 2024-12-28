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

// MLBBaseball creates a slide for baseball odds.
func MLBBaseball(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Baseball_mlb), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "MLB", GetHeader(models.Baseball_mlb), CreateH2HTable(pages, "MLB Baseball", headerString, games)
}
