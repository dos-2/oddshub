/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func parseColorTag(text string) (string, tcell.Color, tcell.Color) {
	// Default colors
	textColor := tcell.ColorWhite
	bgColor := tcell.ColorBlack

	// Check for color tags
	if strings.HasPrefix(text, "[") && strings.Contains(text, "]") && strings.Contains(text, ":") {
		endIndex := strings.Index(text, "]")
		tag := text[1:endIndex]
		text = text[endIndex+1:]

		colors := strings.Split(tag, ":")
		if len(colors) == 2 {
			textColor = tcell.GetColor(colors[0])
			bgColor = tcell.GetColor(colors[1])
		}
	}

	return text, textColor, bgColor
}

func CreateH2HTable(sportName string, tableData string) *tview.Table {
	// Set up event listeners for mouse events
	table := tview.NewTable().
		SetFixed(1, 1).
		SetBorders(false).
		SetSelectable(true, false)

	// Set up the table cells
	rows := strings.Split(tableData, "\n")
	for row := 0; row < len(rows)-1; row++ {
		line := rows[row]
		for column, cell := range strings.Split(line, "|") {
			cellText, textColor, bgColor := parseColorTag(cell)
			align := tview.AlignCenter
			selectable := true

			if row%3 == 0 { // Make every third row unselectable
				selectable = false
			}

			if row == 0 {
				bgColor = tcell.NewRGBColor(57, 255, 20)
				textColor = tcell.ColorBlack
				align = tview.AlignCenter
				selectable = false
			}

			if column == 0 && row == 0 {
				bgColor = tcell.NewRGBColor(0, 255, 255)
			}

			tableCell := tview.NewTableCell(cellText).
				SetTextColor(textColor).
				SetAlign(align).
				SetSelectable(selectable).
				SetExpansion(1).
				SetBackgroundColor(bgColor).
				SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
					Background(tcell.NewRGBColor(0, 255, 255)))

			if cellText == "Money –" {
				tableCell.
					SetClickedFunc(func() bool {
						cellSelected := table.GetCell(row, column)
						if cellSelected.Text == "Money –" {
							cellSelected.
								SetText("Money ▲").
								SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
									Background(tcell.NewRGBColor(255, 0, 0)))
						} else if cellSelected.Text == "Money ▲" {
							cellSelected.
								SetText("Money ▼").
								SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
									Background(tcell.NewRGBColor(0, 255, 0)))
						} else {
							cellSelected.
								SetText("Money –").
								SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
									Background(tcell.NewRGBColor(0, 255, 255)))
						}
						return true
					})
			}

			table.SetCell(row, column, tableCell)
		}
	}

	table.SetBorder(true).SetTitle(sportName)
	return table
}
