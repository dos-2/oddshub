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

// Boxing creates a slide for boxing odds.
func Boxing(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Ranking|Players|Bookmaker|Spread –|Money –|Total –\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Boxing", GetHeader(models.Boxing), CreateH2HTable("Boxing", builder.String())
}
