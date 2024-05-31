package ui

import (
  "fmt"
  "strconv"
  "oddshub/slides"

  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
)

type Slide func(nextSlide func()) (title string, content tview.Primitive)

var app = tview.NewApplication()

func UI() error {
  presentationSlides := slides.GetSlides()
  // The presentation presentationSlides.
  //	presentationSlides := []Slide{
  //	slides.Cover,
  //slides.NFL_football,
  //	  tables.NCAA_football,
  //    tables.NBA_basketball,
  //    tables.NCAA_basketball,
  //    tables.MLB_baseball,
  //    tables.NCAA_baseball,
  //    tables.MMA,
  //    tables.NHL_hockey,
  //    tables.Masters_golf,
  //    tables.French_open_tennis,
  //}

  pages := tview.NewPages()

  // The bottom row has some info on where we are.
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

  // Create the pages for all slides.
  previousSlide := func() {
    slide, _ := strconv.Atoi(info.GetHighlights()[0])
    slide = (slide - 1 + len(presentationSlides)) % len(presentationSlides)
    info.Highlight(strconv.Itoa(slide)).
      ScrollToHighlight()
  }
  nextSlide := func() {
    slide, _ := strconv.Atoi(info.GetHighlights()[0])
    slide = (slide + 1) % len(presentationSlides)
    info.Highlight(strconv.Itoa(slide)).
      ScrollToHighlight()
  }
  for index, slide := range presentationSlides {
    title, primitive := slide(nextSlide)
    pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
    fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
  }
  info.Highlight("0")

  // Create the main layout.
  layout := tview.NewFlex().
    SetDirection(tview.FlexRow).
    AddItem(pages, 0, 1, true).
    AddItem(info, 1, 1, false)

  // Shortcuts to navigate the slides.
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

  // Start the application.
  return app.SetRoot(layout, true).EnableMouse(true).EnablePaste(true).Run()
}
