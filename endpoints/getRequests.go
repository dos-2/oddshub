/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"oddshub/models"
	"oddshub/sports"
	"sort"
	"sync"
)

type Events []models.Event

// // Implementing sort.Interface for Events based on CommenceTime.
func (e Events) Len() int           { return len(e) }
func (e Events) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Events) Less(i, j int) bool { return e[i].CommenceTime.Before(e[j].CommenceTime) }

func GetSportsOdds(activeSports []models.ActiveSport) map[string][]models.Event {
	eventsMap := make(map[string][]models.Event)
	sportsAvailable := sports.GetSportsMap()
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, sport := range activeSports {
		wg.Add(1)
		sportName := sports.Sport(sport.Key)
		if _, exists := sportsAvailable[string(sportName)]; !exists {
			wg.Done()
			continue
		}
		//log.Printf("Fetching events for %s", sport)
		go fetchEventsMap(sports.Sport(sport.Key), &wg, &mu, eventsMap)
	}
	wg.Wait()
	return eventsMap
}

func fetchEventsMap(sport sports.Sport, wg *sync.WaitGroup, mu *sync.Mutex, eventsMap map[string][]models.Event) {
	defer wg.Done()
	var events []models.Event
	// Fetch events
	url := fmt.Sprintf("https://oddshub-backend-wc4uw.ondigitalocean.app/sports-data/%s", sport)
	req, err := http.Get(url)
	if err != nil {
		log.Println("Error making request:", err)
		return
	}
	defer req.Body.Close()

	response, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return
	}
	if err := json.Unmarshal(response, &events); err != nil {
		log.Printf("Something is wrong with the endpoint for events")
	}
	if events == nil {
		return
	}
	// sort by commence time function
	sort.Sort(Events(events))
	mu.Lock()
	eventsMap[string(sport)] = events
	mu.Unlock()
}

func FetchActiveSports() (map[string]models.ActiveSport, []models.ActiveSport, error) {
	url := "https://oddshub-backend-wc4uw.ondigitalocean.app/sports-data/active_sports"
	req, err := http.Get(url)
	if err != nil {
		log.Println("Error making request:", err)
		return nil, nil, err
	}
	defer req.Body.Close()

	var activeSports []models.ActiveSport
	var activeSportsMap = make(map[string]models.ActiveSport)
	response, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return nil, nil, err
	}
	if err := json.Unmarshal(response, &activeSports); err != nil {
		log.Printf("Something is wrong with the endpoint for active sports")
		return nil, nil, err
	}
	for _, sport := range activeSports {
		activeSportsMap[string(sport.Key)] = sport
	}
	return activeSportsMap, activeSports, nil
}
