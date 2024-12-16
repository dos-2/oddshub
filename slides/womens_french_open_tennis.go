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

// WomensFrenchOpenTennis creates a slide for tennis odds.
func WomensFrenchOpenTennis(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Players|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Tennis_wta_french_open), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "Womens French Open", GetHeader(models.Tennis_wta_french_open), CreateH2HTable(pages, "Womens French Open Tennis", headerString, games)
}
