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

func NCAAFootball(games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n")
	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NCAA Football", GetHeader(models.Americanfootball_ncaaf), CreateH2HTable("NCAA Football", tableData.String())
}
