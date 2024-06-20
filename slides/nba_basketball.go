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

// NBABasketball creates a slide for basketball odds.
func NBABasketball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "NBA", GetHeader(models.Basketball_nba), CreateH2HTable("NBA Basketball", tableData.String())
}
