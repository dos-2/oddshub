/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"fmt"
	"strings"
	"time"

	"github.com/dos-2/oddshub/models"
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

func GetBackground(cellText string) tcell.Color {
	if strings.Contains(cellText, "▲") {
		return tcell.NewRGBColor(255, 0, 0)
	} else if strings.Contains(cellText, "▼") {
		return tcell.NewRGBColor(0, 255, 0)
	} else if strings.Contains(cellText, "–") {
		return tcell.NewRGBColor(255, 255, 0)
	}
	return tcell.NewRGBColor(255, 255, 255)
}

func GetClickFunction(pages *tview.Pages,
	table *tview.Table,
	row int,
	column int,
	cellText string,
	cellField string,
	tableHeader string,
	sportName string) func() bool {
	return func() bool {
		debug := models.GetDebug()
		cellSelected := table.GetCell(row, column)
		currentPage := models.GetCurrentPage()
		sortedEvents := models.GetLoadedEvents(currentPage)
		newTableHeader := ""

		switch cellText {
		case "Money":
			if strings.Contains(tableHeader, "Money –") {
				newTableHeader = strings.Replace(tableHeader, "Money –", "Money ▲", 1)
			} else if strings.Contains(tableHeader, "Money ▲") {
				newTableHeader = strings.Replace(tableHeader, "Money ▲", "Money ▼", 1)
			} else if strings.Contains(tableHeader, "Money ▼") {
				newTableHeader = strings.Replace(tableHeader, "Money ▼", "Money –", 1)
			}
		case "Spread":
			if strings.Contains(tableHeader, "Spread –") {
				newTableHeader = strings.Replace(tableHeader, "Spread –", "Spread ▲", 1)
			} else if strings.Contains(tableHeader, "Spread ▲") {
				newTableHeader = strings.Replace(tableHeader, "Spread ▲", "Spread ▼", 1)
			} else if strings.Contains(tableHeader, "Spread ▼") {
				newTableHeader = strings.Replace(tableHeader, "Spread ▼", "Spread –", 1)
			}
		case "Total":
			if strings.Contains(tableHeader, "Total –") {
				newTableHeader = strings.Replace(tableHeader, "Total –", "Total ▲", 1)
			} else if strings.Contains(tableHeader, "Total ▲") {
				newTableHeader = strings.Replace(tableHeader, "Total ▲", "Total ▼", 1)
			} else if strings.Contains(tableHeader, "Total ▼") {
				newTableHeader = strings.Replace(tableHeader, "Total ▼", "Total –", 1)
			}
		}

		if cellSelected.Text == cellText+" –" {
			if debug {
				fmt.Printf("[%s] Sort %s by asc", time.Now(), cellText)
				fmt.Println()
				fmt.Printf("[%s] Length of events loaded to be sorted: %d", time.Now(), len(sortedEvents))
				fmt.Println()
			}
			sortedEvents = sortEvents(sortedEvents, cellField, "asc")
		} else if cellSelected.Text == cellText+" ▲" {
			if debug {
				fmt.Printf("[%s] Sort %s by desc", time.Now(), cellText)
				fmt.Println()
				fmt.Printf("[%s] Length of events loaded to be sorted: %d", time.Now(), len(sortedEvents))
				fmt.Println()
			}
			sortedEvents = sortEvents(sortedEvents, cellField, "desc")
		} else {
			if debug {
				fmt.Printf("[%s] Sort %s by none", time.Now(), cellText)
				fmt.Println()
				fmt.Printf("[%s] Length of events loaded to be sorted: %d", time.Now(), len(sortedEvents))
				fmt.Println()
			}
		}
		table := CreateH2HTable(pages, sportName, newTableHeader, sortedEvents)
		content := tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(table, 0, 1, true)
		pages.AddAndSwitchToPage(models.GetCurrentPageIndex(), content, true)
		return true
	}
}

func CreateH2HTable(pages *tview.Pages, sportName string, tableHeader string, games []models.Event) *tview.Table {
	var builder strings.Builder
	builder.WriteString(tableHeader)
	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	tableData := builder.String()

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

			if strings.Contains(cellText, "Money") {
				tableCell.
					SetClickedFunc(GetClickFunction(pages, table, row, column, "Money", "money", tableHeader, sportName)).
					SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
						Background(GetBackground(cellText)))
			} else if strings.Contains(cellText, "Spread") {
				tableCell.
					SetClickedFunc(GetClickFunction(pages, table, row, column, "Spread", "spread", tableHeader, sportName)).
					SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
						Background((GetBackground(cellText))))
			} else if strings.Contains(cellText, "Total") {
				tableCell.
					SetClickedFunc(GetClickFunction(pages, table, row, column, "Total", "total", tableHeader, sportName)).
					SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
						Background((GetBackground(cellText))))
			}

			table.SetCell(row, column, tableCell)
		}
	}

	table.SetBorder(true).SetTitle(sportName)
	return table
}
