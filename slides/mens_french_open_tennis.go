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

// FrenchOpenTennis creates a slide for tennis odds.
func MensFrenchOpenTennis(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	headerString := "Commencement Date|Location|Players|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Tennis_atp_french_open), games)

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Mens French Open", GetHeader(models.Tennis_atp_french_open), CreateH2HTable(pages, "Mens French Open Tennis", headerString, games)
}
