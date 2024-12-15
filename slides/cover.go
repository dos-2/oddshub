/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"fmt"
	"strings"

	"github.com/dos-2/oddshub/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const logo = `
                 /$$       /$$           /$$                 /$$      
                | $$      | $$          | $$                | $$      
  /$$$$$$   /$$$$$$$  /$$$$$$$  /$$$$$$$| $$$$$$$  /$$   /$$| $$$$$$$ 
 /$$__  $$ /$$__  $$ /$$__  $$ /$$_____/| $$__  $$| $$  | $$| $$__  $$
| $$  \ $$| $$  | $$| $$  | $$|  $$$$$$ | $$  \ $$| $$  | $$| $$  \ $$
| $$  | $$| $$  | $$| $$  | $$ \____  $$| $$  | $$| $$  | $$| $$  | $$
|  $$$$$$/|  $$$$$$$|  $$$$$$$ /$$$$$$$/| $$  | $$|  $$$$$$/| $$$$$$$/
 \______/  \_______/ \_______/|_______/ |__/  |__/ \______/ |_______/ 

`

const (
	subtitle   = `oddshub - A terminal UI for sports betting programmers`
	navigation = `[#FF8C00]Ctrl-N[#00FFFF]: Next slide    [#FF8C00]Ctrl-P[#00FFFF]: Previous slide    [#FF8C00]Ctrl-C[#00FFFF]: Exit`
	mouse      = `[#00FFFF](or use your mouse)`
)

// Cover returns the cover page.
func Cover(games []models.Event) (title string, header string, content tview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.NewRGBColor(57, 255, 20))
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mouse, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "oddshub", "", flex
}
