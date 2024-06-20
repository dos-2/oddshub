/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

// WomensFrenchOpenTennis creates a slide for tennis odds.
func WomensFrenchOpenTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Players|Bookmaker|Spread|Money|Total\n")
	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "Womens French Open", GetHeader(models.Tennis_wta_french_open), CreateH2HTable("Womens French Open Tennis", tableData.String())
}
