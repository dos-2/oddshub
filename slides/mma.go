/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"strings"

	"github.com/dos-2/oddshub/models"

	"github.com/rivo/tview"
)

// Mma creates a slide for MMA odds.
func Mma(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Ranking|Players|Bookmaker|Spread –|Money –|Total –\n")

	for _, game := range games {
		tableData.WriteString(FormatTeamEvent(game))
	}

	return "MMA", GetHeader(models.Mma_mixed_martial_arts), CreateH2HTable("MMA", tableData.String())
}
