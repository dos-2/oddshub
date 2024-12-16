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

// Boxing creates a slide for boxing odds.
func Boxing(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	headerString := "Commencement Date|Ranking|Players|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Boxing), games)

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Boxing", GetHeader(models.Boxing), CreateH2HTable(pages, "Boxing", headerString, games)
}
