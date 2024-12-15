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

// NcaaBasketball creates a slide for basketball odds.
func NCAABasketball(games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n")

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "NCAA Basketball", GetHeader(models.Basketball_ncaa), CreateH2HTable("NCAA Basketball", tableData.String())
}
