/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package models

import (
	"fmt"
	"time"
)

var loadedEvents = make(map[string][]Event)

func AddLoadedEvent(key string, value []Event) {
	tempMap := loadedEvents
	tempMap[key] = value
	loadedEvents = tempMap
}

func GetLoadedEvents(key string) []Event {
	return loadedEvents[key]
}

var currentPage = ""

func SetCurrentPage(page string) {
	currentPage = page
}

func GetCurrentPage() string {
	return currentPage
}

var currentPageIndex = ""

func SetCurrentPageIndex(index string) {
	currentPageIndex = index
}

func GetCurrentPageIndex() string {
	return currentPageIndex
}

var debug = false

func SetDebug(option bool) {
	debug = option
}

func GetDebug() bool {
	return debug
}

func LoadEvent(sport string, games []Event) {
	debug := GetDebug()

	if len(GetLoadedEvents(sport)) == 0 {
		if debug {
			fmt.Printf("[%s] Add loaded events for %s", time.Now(), sport)
			fmt.Println()
		}
		AddLoadedEvent(sport, games)
	}
}
