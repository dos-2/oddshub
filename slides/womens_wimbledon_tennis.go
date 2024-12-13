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

func WomensWimbledonTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money –|Total\n")
	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "Womens Wimbledon", GetHeader(models.Tennis_wta_wimbledon), CreateH2HTable(string(models.Tennis_wta_wimbledon), tableData.String())
}
