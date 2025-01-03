/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"strings"

	"github.com/dos-2/oddshub/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateRoundRobinTable(sportName string, tableHeader string, games []models.Event) *tview.Table {
	var builder strings.Builder

	builder.WriteString(tableHeader)

	if len(games) > 0 {
		builder.WriteString(FormatTournamentEvent(games[0]))
	}

	// Set up event listeners for mouse events
	table := tview.NewTable().
		SetFixed(1, 1).
		SetBorders(false).
		SetSelectable(true, false)

		// Set up the table cells

	tableData := builder.String()

	rows := strings.Split(tableData, "\n")
	for row := 0; row < len(rows)-1; row++ {
		line := rows[row]
		for column, cell := range strings.Split(line, "|") {
			cellText, textColor, bgColor := parseColorTag(cell)
			align := tview.AlignCenter
			selectable := true

			if row == 0 {
				bgColor = tcell.NewRGBColor(57, 255, 20)
				textColor = tcell.ColorBlack
				align = tview.AlignCenter
				selectable = false
			}

			if column == 0 && row == 0 {
				bgColor = tcell.NewRGBColor(0, 255, 255)
			}
			if row != 0 && row%2 == 0 {
				bgColor = tcell.GetColor("#000000")
			}

			tableCell := tview.NewTableCell(cellText).
				SetTextColor(textColor).
				SetAlign(align).
				SetSelectable(selectable).
				SetExpansion(1).
				SetBackgroundColor(bgColor).
				SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
					Background(tcell.NewRGBColor(0, 255, 255)))

			table.SetCell(row, column, tableCell)
		}
	}

	table.SetBorder(true).SetTitle(sportName)
	return table
}
