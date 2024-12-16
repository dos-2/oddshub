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

// Mma creates a slide for MMA odds.
func Mma(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var tableData strings.Builder
	headerString := "Commencement Date|Ranking|Players|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Mma_mixed_martial_arts), games)

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "MMA", GetHeader(models.Mma_mixed_martial_arts), CreateH2HTable(pages, "MMA", headerString, games)
}
