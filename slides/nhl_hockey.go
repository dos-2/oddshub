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

func NHLHockey(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Icehockey_nhl), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NHL", GetHeader(models.Icehockey_nhl), CreateH2HTable(pages, "NHL Hockey", headerString, games)
}
