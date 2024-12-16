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

func BrazilCampeonato(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	var builder strings.Builder
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Soccer_brazil_campeonato), games)

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Brazil Campeonato", GetHeader(models.Soccer_brazil_campeonato), CreateH2HTable(pages, string(models.Soccer_brazil_campeonato), headerString, games)
}
