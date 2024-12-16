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

func NRLRugby(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Rugbyleague_nrl), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NRL", GetHeader(models.Rugby_irl), CreateH2HTable(pages, string(models.Rugby_irl), headerString, games)
}
