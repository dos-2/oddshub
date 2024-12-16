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

func NCAAFootball(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Americanfootball_ncaaf), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NCAA Football", GetHeader(models.Americanfootball_ncaaf), CreateH2HTable(pages, "NCAA Football", headerString, games)
}
