package API

import(
  "encoding/json"
  "net/http"
  "fmt"
  "io"
  "time"
  "os"
  "github.com/joho/godotenv"
  "log"
  "oddshub/sport"
)

type Event struct {
  ID             string     `json:"id"`
  SportKey       string     `json:"sport_key"`
  SportTitle     string     `json:"sport_title"`
  CommenceTime   time.Time  `json:"commence_time"`
  HomeTeam       string     `json:"home_team"`
  AwayTeam       string     `json:"away_team"`
  Bookmakers     []Bookmaker `json:"bookmakers"`
}

// Define a struct for the bookmaker
type Bookmaker struct {
  Key         string    `json:"key"`
  Title       string    `json:"title"`
  LastUpdate  time.Time `json:"last_update"`
  Markets     []Market  `json:"markets"`
}

// Define a struct for the market
type Market struct {
  Key         string    `json:"key"`
  LastUpdate  time.Time `json:"last_update"`
  Outcomes    []Outcome `json:"outcomes"`
}

// Define a struct for the outcome
type Outcome struct {
  Name    string  `json:"name"`
  Price   float64 `json:"price"`
  Point   float64 `json:"point,omitempty"`
}

func getAmericanFootball_ncaa() []Event {
  return nil
}


func getAmericanFootball_nfl() []Event {

  return nil
}

func getBaseball_ncaa() []Event {

  return nil
}

func getBaseball_mlb() []Event {

  return nil
}

func getBasketball_ncaa() []Event {

  return nil
}

func getBasketball_nba() []Event {

  return nil
}

func getGolf_masters() []Event {

  return nil
}

func geticehocky_nhl() []Event {

  return nil
}

func getMma() []Event {

  return nil
}

func getSoccer_uefa() []Event {

  return nil
}

func getSoccer_mls() []Event {

  return nil
}

func getTennis_french_open() []Event {

  return nil
}

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

func GetAllUpcomingEvents() []Event {
  var apiKey string = getAPIKEY()
  var allEvents []Event

  for _, sport := range sport.AllSports() {

    req, err := http.Get(fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds/?apiKey=%s&regions=us&markets=h2h,spreads,totals&oddsFormat=american", sport, apiKey))      

    if err != nil {
      fmt.Print(err.Error())
    }

    defer req.Body.Close()
    response, err := io.ReadAll(req.Body)

    if err != nil {
      fmt.Print(err.Error())
    }

    var events []Event
    msg := json.Unmarshal(response, &events)

    if msg  != nil {
      fmt.Println(err)
    }

    for i, event := range events {
      // Initialize a slice to store filtered bookmakers
      var filteredBookmakers []Bookmaker
      // Iterate over each bookmaker
      for _, bookmaker := range event.Bookmakers {
        if bookmaker.Title == "FanDuel" {
          // If it is, append it to the filteredBookmakers slice
          filteredBookmakers = append(filteredBookmakers, bookmaker)
        }
      }

      if len(filteredBookmakers) == 0 && len(event.Bookmakers) != 0 {
        filteredBookmakers = append(filteredBookmakers, event.Bookmakers[0])
      }
      events[i].Bookmakers = filteredBookmakers
    }
    allEvents = append(allEvents, events...)
  }

  return allEvents
}

