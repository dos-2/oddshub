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

func fetchEventsMap(test bool, sport sports.Sport, apiKey string, wg *sync.WaitGroup, mu *sync.Mutex, eventMap map[string][]models.Event) {
  defer wg.Done()

  var events []models.Event
  if test {
    jsonData, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", sport))
    if err != nil {
      log.Println("Error reading test data:", err)
      return
    }

    if err := json.Unmarshal(jsonData, &events); err != nil {
      log.Println("Error unmarshalling test data:", err)
      return
    }
  } else {
    var url string
    if sport == sports.Golf_masters_tournament_winner || sport == sports.Golf_pga_championship_winner {
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

    if err := json.Unmarshal(response, &events); err != nil {
      // Unmarshal into ErrorCode struct to check for error response
      var errorCode models.ErrorCode
      if err := json.Unmarshal(response, &errorCode); err != nil {
        log.Printf("Error unmarshalling response for sport %s: %v\nResponse Data: %s", sport, err, string(response))
      } else {
        log.Printf("API Error for sport %s: %s\nDetails: %s", sport, errorCode.Message, errorCode.DetailsURL)
      }
      return
    }
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

// FetchActiveSports retrieves the list of active sports. If test is true, it reads from a test file.
func FetchActiveSports(test bool) (map[string]models.ActiveSport, error) {
  var activeSportsList []models.ActiveSport
  activeSportsMap := make(map[string]models.ActiveSport)

  if test {
    jsonData, err := os.ReadFile("testdata/active_sports.json")
    if err != nil {
      return nil, fmt.Errorf("error reading test data: %w", err)
    }

    if err := json.Unmarshal(jsonData, &activeSportsList); err != nil {
      return nil, fmt.Errorf("error unmarshalling test data: %w", err)
    }
  } else {
    apiKey := getAPIKEY() // Assumes you have a function to retrieve the API key
    url := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/?apiKey=%s", apiKey)

    resp, err := http.Get(url)
    if err != nil {
      return nil, fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    response, err := io.ReadAll(resp.Body)
    if err != nil {
      return nil, fmt.Errorf("error reading response: %w", err)
    }

if err := json.Unmarshal(response, &activeSportsMap); err != nil {
			// Unmarshal into ErrorCode struct to check for error response
			var errorCode models.ErrorCode
			if err := json.Unmarshal(response, &errorCode); err != nil {
				return nil, fmt.Errorf("Error unmarshalling response for active sports: %v\nResponse Data: %s", err, string(response))
			} else {
				return nil, fmt.Errorf("API Error for active sports: %s\nDetails: %s",errorCode.Message, errorCode.DetailsURL)
			}
		}
  }
  for _, sport := range activeSportsList {
    activeSportsMap[string(sport.Key)] = sport
  }

  return activeSportsMap, nil
}

func GetAllUpcomingEventsMap(test bool) map[string][]models.Event {
  apiKey := getAPIKEY()
  eventMap := make(map[string][]models.Event)
  var wg sync.WaitGroup
  var mu sync.Mutex

  for _, sport := range sports.AllSports() {
    wg.Add(1)
    go fetchEventsMap(test, sport, apiKey, &wg, &mu, eventMap)
  }

  wg.Wait()
  return eventMap
}
