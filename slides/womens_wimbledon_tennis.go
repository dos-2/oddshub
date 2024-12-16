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

func WomensWimbledonTennis(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Tennis_wta_wimbledon), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "Womens Wimbledon", GetHeader(models.Tennis_wta_wimbledon), CreateH2HTable(pages, string(models.Tennis_wta_wimbledon), headerString, games)
}
