/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package models

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

var debug = false

func SetDebug(option bool) {
	debug = option
}

func GetDebug() bool {
	return debug
}
