package API

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"oddshub/models"
	"oddshub/sports"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func getAPIKEY() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API key is missing")
	}

	return apiKey
}

func fetchEventsMap(sport sports.Sport, apiKey string, wg *sync.WaitGroup, mu *sync.Mutex, eventMap map[string][]models.Event) {
	defer wg.Done()

	var url string
	if sport == sports.Golf_masters_tournament_winner {
		url = fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds/?apiKey=%s&regions=us&markets=&oddsFormat=american&commenceTimeFrom=2024-06-04T00:00:00Z&commenceTimeTo=2024-09-29T00:00:00Z", sport, apiKey)
	} else {
		url = fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds/?apiKey=%s&regions=us&markets=h2h,spreads,totals&oddsFormat=american&commenceTimeFrom=2024-06-04T00:00:00Z&commenceTimeTo=2024-09-29T00:00:00Z", sport, apiKey)
	}

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

	var events []models.Event
	if err := json.Unmarshal(response, &events); err != nil {
		log.Println("Error unmarshalling response:", err)
		return
	}

	for i, event := range events {
		var filteredBookmakers []models.Bookmaker
		for _, bookmaker := range event.Bookmakers {
			if bookmaker.Title == "DraftKings" {
				filteredBookmakers = append(filteredBookmakers, bookmaker)
			}
		}

		if len(filteredBookmakers) == 0 && len(event.Bookmakers) != 0 {
			filteredBookmakers = append(filteredBookmakers, event.Bookmakers[0])
		}
		events[i].Bookmakers = filteredBookmakers
	}

	mu.Lock()
	eventMap[string(sport)] = events
	mu.Unlock()
}

func GetAllUpcomingEventsMap() map[string][]models.Event {
	apiKey := getAPIKEY()
	eventMap := make(map[string][]models.Event)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, sport := range sports.AllSports() {
		wg.Add(1)
		go fetchEventsMap(sport, apiKey, &wg, &mu, eventMap)
	}

	wg.Wait()
	return eventMap
}
