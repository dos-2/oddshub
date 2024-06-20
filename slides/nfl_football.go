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

// FootballSlide creates a slide for football odds.
func NFLfootball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")
	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NFL", GetHeader(models.Americanfootball_nfl), CreateH2HTable("NFL Football", tableData.String())
}
