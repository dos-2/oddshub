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

// FootballSlide creates a slide for football odds.
func NFLfootball(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Americanfootball_nfl), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}
	return "NFL", GetHeader(models.Americanfootball_nfl), CreateH2HTable(pages, "NFL Football", headerString, games)
}
