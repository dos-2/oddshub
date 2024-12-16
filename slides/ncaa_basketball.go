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

// NcaaBasketball creates a slide for basketball odds.
func NCAABasketball(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Basketball_ncaa), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "NCAA Basketball", GetHeader(models.Basketball_ncaa), CreateH2HTable(pages, "NCAA Basketball", headerString, games)
}
