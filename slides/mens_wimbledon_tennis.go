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

func MensWimbledonTennis(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Tennis_atp_wimbledon), games)

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Wimbledon", GetHeader(models.Tennis_atp_wimbledon), CreateH2HTable(pages, string(models.Tennis_atp_wimbledon), headerString, games)
}
