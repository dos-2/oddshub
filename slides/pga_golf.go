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

func PGAGolf(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var tableData strings.Builder
	tableData.WriteString("Commencement Date|Teams|Players|Bookmaker|Outrights||\n")
	if len(games) > 0 {
		tableData.WriteString(FormatTournamentEvent(games[0]))
	}
	return "PGA", GetHeader(models.Golf_pga_tournament_winner), CreateRoundRobinTable(string(models.Golf_pga_tournament_winner), tableData.String())
}
