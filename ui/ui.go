/*
A presentation of the UI package, implemented with tview.

# Navigation
  - Ctrl-N: Jump to next slide
  - Ctrl-P: Jump to previous slide
*/
package ui

import (
	"fmt"
	"log"
	"oddshub/endpoints"
	"oddshub/models"
	"oddshub/slides"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Slide func(nextSlide func()) (title string, content tview.Primitive)

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
	info := createInfoTextView(pages)
	// Create navigation functions
	previousSlide, nextSlide := createNavigationFunctions(info, presentationSlides)
	// Create pages for all slides
	setupSlides(events, pages, info, presentationSlides, nextSlide)
	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)
	setupNavigationShortcuts(app, previousSlide, nextSlide)
	// Start the application.
	return app.SetRoot(layout, true).EnableMouse(true).EnablePaste(true).Run()
}

func createInfoTextView(pages *tview.Pages) *tview.TextView {
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			if len(added) == 0 {
				return
			}
			pages.SwitchToPage(added[0])
		})
	return info
}

func createNavigationFunctions(info *tview.TextView, slides []slides.Slide) (func(), func()) {
	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(slide)).ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).ScrollToHighlight()
	}
	return previousSlide, nextSlide
}

func setupSlides(events map[string][]models.Event, pages *tview.Pages, info *tview.TextView, slides []slides.Slide, nextSlide func()) {
	for index, slide := range slides {
		eventList, exists := events[slide.Name]
		if !exists || index == 0 {
			eventList = []models.Event{}
		}

		title, _, primitive := slide.Content(eventList, nextSlide)
		var content tview.Primitive
		content = tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(primitive, 0, 1, true)
		if index == 0 {
			content = primitive
		}
		pages.AddPage(strconv.Itoa(index), content, true, index == 0)
		fmt.Fprintf(info, `["%d"][#00FFFF]%s[white][""]  `, index, title)
	}
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
	events := endpoints.GetSportsOdds(activeSports)
	if events == nil {
		events = make(map[string][]models.Event)
	}
	return events
}
func loadActiveSports() (map[string]models.ActiveSport, []models.ActiveSport) {
	activeSportsMap, activeSportsList, err := endpoints.FetchActiveSports()
	if err != nil {
		activeSportsMap = make(map[string]models.ActiveSport)
		activeSportsList = []models.ActiveSport{}
	}
	return activeSportsMap, activeSportsList
}
