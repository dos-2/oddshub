package API
import(
  "encoding/json"
  "net/http"
  "fmt"
  "io/ioutil"
  "time"
  "os"
  "github.com/joho/godotenv"
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

func GetAllUpcomingEvents() []Event {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }

  apiKey := os.Getenv("API_KEY")

  if apiKey == "" {
    fmt.Println("API key not found in environment variables")
    return nil 
  }

  req, err := http.Get(fmt.Sprintf("https://api.the-odds-api.com/v4/sports/upcoming/odds/?apiKey=%s&regions=us&markets=h2h,spreads,totals&oddsFormat=american", apiKey))      

  if err != nil {
    fmt.Print(err.Error())
  }
  defer req.Body.Close()
  response, err := ioutil.ReadAll(req.Body)

  if err != nil {
    fmt.Print(err.Error())
  }

  var events []Event
  msg := json.Unmarshal(response, &events)

  if msg  != nil {
    fmt.Println(err)
  }

  //var filteredEvents []Event
  //set := map[string]bool{
  //  "americanfootball_ncaaf":              true,
  //  "americanfootball_nfl":                true,
  //  "basketball_nba":                      true,
  //  "golf_masters_tournament_winner":      true,
  //  "icehockey_nhl":                       true,
  //  "icehockey_nhl_championship_winner":   true,
  //  "mma_mixed_martial_arts":              true,
  //  "soccer_usa_mls":                      true,
  //  "tennis_atp_french_open":              true,
  //  "baseball_mlb":                        true,
  //  "baseball_ncaa":                       true,
  //  "basketball_ncaa":                     true,
  //  "golf_us_open_winner":                 true,
  //  "basketball_nba_championship_winner":  true,
  //  "basketball_nba_championship_winner":  true,
  //}
  //
  ///*for _,event := range events {
  //  fmt.Print(event.SportKey)
  //  if set[event.SportKey] {
  //      filteredEvents = append(filteredEvents, event)
  //  }
  //}
  //events = filteredEvents

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
    events[i].Bookmakers = filteredBookmakers
  }

  return events
}
