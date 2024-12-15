/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
/*
# Navigation
  - Ctrl-N: Jump to next slide
  - Ctrl-P: Jump to previous slide
*/

package ui

import (
	"fmt"
	"log"

	"strconv"
	"time"

	"github.com/dos-2/oddshub/endpoints"
	"github.com/dos-2/oddshub/models"
	"github.com/dos-2/oddshub/slides"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func RunApp() error {
	// Load active sports
	activeSportsMap, activeSports := loadActiveSports()
	// Load slides
	presentationSlides, err := slides.GetActiveSlides(activeSportsMap)
	if err != nil {
		log.Print("Something is wrong with getting active slides.")
	}

	// Load sports events
	events := loadEvents(activeSports)

	pages := tview.NewPages()
	// Info text view
	info := createInfoTextView(pages, presentationSlides)
	// Create navigation functions
	previousSlide, nextSlide := createNavigationFunctions(info, presentationSlides)
	// Create pages for all slides
	setupSlides(events, pages, info, presentationSlides)
	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)
	setupNavigationShortcuts(app, previousSlide, nextSlide)
	// Start the application.
	return app.SetRoot(layout, true).EnableMouse(true).EnablePaste(true).Run()
}

func createInfoTextView(pages *tview.Pages, slides []slides.Slide) *tview.TextView {
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			if len(added) == 0 {
				return
			}

			pageIndex, _ := strconv.Atoi(added[0])
			pageName := slides[pageIndex].Name

			// Log
			if models.GetDebug() {
				fmt.Println("Jump To Slide")
				fmt.Printf(`[%s] %s/%d: %s`, time.Now(), added[0], pages.GetPageCount()-1, pageName)
				fmt.Println()
			}

			models.SetCurrentPage(pageName)
			pages.SwitchToPage(added[0])
		})
	return info
}

func createNavigationFunctions(info *tview.TextView, slides []slides.Slide) (func(), func()) {
	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)

		// Log
		if models.GetDebug() {
			fmt.Println("Previous Slide")
			fmt.Printf(`[%s] %d/%d: %s`, time.Now(), slide, len(slides), slides[slide].Name)
			fmt.Println()
		}

		info.Highlight(strconv.Itoa(slide)).ScrollToHighlight()
		models.SetCurrentPage(slides[slide].Name)
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)

		// Log
		if models.GetDebug() {
			fmt.Println("Next Slide")
			fmt.Printf(`[%s] %d/%d: %s`, time.Now(), slide, len(slides), slides[slide].Name)
			fmt.Println()
		}

		info.Highlight(strconv.Itoa(slide)).ScrollToHighlight()
		models.SetCurrentPage(slides[slide].Name)
	}
	return previousSlide, nextSlide
}

func setupSlides(events map[string][]models.Event, pages *tview.Pages, info *tview.TextView, slides []slides.Slide) {
	fmt.Println("Setup Slides")
	for index, slide := range slides {
		eventList, exists := events[slide.Name]
		if !exists || index == 0 {
			eventList = []models.Event{}
		}

		title, _, primitive := slide.Content(eventList)
		var content tview.Primitive
		content = tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(primitive, 0, 1, true)
		if index == 0 {
			content = primitive
		}
		pages.AddPage(strconv.Itoa(index), content, true, index == 0)

		// Log
		if models.GetDebug() {
			fmt.Printf(`[%s] %d: %s - %s`, time.Now().String(), index, title, slide.Name)
			fmt.Println()
			fmt.Fprintf(info, `["%d"][#00FFFF]%s[white][""]  `, index, title)
		}
	}
	fmt.Println("End Setup")
	fmt.Println()
	info.Highlight("0")
}

func setupNavigationShortcuts(app *tview.Application, previousSlide func(), nextSlide func()) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		}
		return event
	})
}

func loadEvents(activeSports []models.ActiveSport) map[string][]models.Event {
	eventsChan := make(chan map[string][]models.Event)
	go func() {
		eventsChan <- endpoints.GetSportsOdds(activeSports)
	}()

	select {
	case events := <-eventsChan:
		return events
	case <-time.After(5 * time.Second): // Timeout example
		return make(map[string][]models.Event)
	}
}

func loadActiveSports() (map[string]models.ActiveSport, []models.ActiveSport) {
	activeSportsMap, activeSportsList, err := endpoints.FetchActiveSports()
	if err != nil {
		activeSportsMap = make(map[string]models.ActiveSport)
		activeSportsList = []models.ActiveSport{}
	}
	return activeSportsMap, activeSportsList
}
