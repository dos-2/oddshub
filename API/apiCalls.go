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

  func GetAllUpcomingEvents() string {
      err := godotenv.Load()
      if err != nil {
        fmt.Println("Error loading .env file")
      }

      apiKey := os.Getenv("API_KEY")
      if apiKey == "" {
          fmt.Println("API key not found in environment variables")
          return ""
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
     var event []Event 
      msg := json.Unmarshal(response, &event)
      if msg  != nil {
        fmt.Println(err)
      } 
  //    for _, event := range event {
	//	    fmt.Printf("ID: %s\n", event.ID)
	//	    fmt.Printf("Sport: %s\n", event.SportTitle)
	//	    fmt.Printf("Commence Time: %s\n", event.CommenceTime)
	//	    fmt.Printf("Home Team: %s\n", event.HomeTeam)
	//	    fmt.Printf("Away Team: %s\n", event.AwayTeam)
	//	    fmt.Println("Bookmakers:")
	//	    for _, bookmaker := range event.Bookmakers {
	//		    fmt.Printf("  - Title: %s\n", bookmaker.Title)
	//		    fmt.Println("    Markets:")
	//		    for _, market := range bookmaker.Markets {
	//			    fmt.Printf("      - Key: %s\n", market.Key)
	//			    fmt.Printf("        Last Update: %s\n", market.LastUpdate)
	//			    fmt.Println("        Outcomes:")
	//			    for _, outcome := range market.Outcomes {
	//				    fmt.Printf("          - Name: %s, Price: %f, Point: %f\n", outcome.Name, outcome.Price, outcome.Point)
	//			}
	//		}
	//	}
	//	fmt.Println()
	//}

fmt.Printf("ID: %s\n", event[0].ID)
fmt.Printf("Sport: %s\n", event[0].SportTitle)
fmt.Printf("Commence Time: %s\n", event[0].CommenceTime)
fmt.Printf("Home Team: %s\n", event[0].HomeTeam)
fmt.Printf("Away Team: %s\n", event[0].AwayTeam)
fmt.Println("Bookmakers:")
for _, bookmaker := range event[0].Bookmakers {
    fmt.Printf("  - Title: %s\n", bookmaker.Title)
    fmt.Println("    Markets:")
    for _, market := range bookmaker.Markets {
        fmt.Printf("      - Key: %s\n", market.Key)
        fmt.Printf("        Last Update: %s\n", market.LastUpdate)
        fmt.Println("        Outcomes:")
        for _, outcome := range market.Outcomes {
            fmt.Printf("          - Name: %s, Price: %f, Point: %f\n", outcome.Name, outcome.Price, outcome.Point)
        }
    }
}
fmt.Println()
      return string(response)
}
