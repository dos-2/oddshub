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

func EPLSoccer(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Soccer_epl), games)

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "EPL Soccer", GetHeader(models.Soccer_epl), CreateH2HTable(pages, string(models.Soccer_epl), headerString, games)
}
