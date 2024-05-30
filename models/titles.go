package models

type Title string

const(
  Americanfootball_ncaaf Title = "NCAA Football" 
  Americanfootball_nfl Title = "NFL Football"
  Basketball_nba Title = "NBA Basketball"
  Golf_masters_tournament_winner Title = "The Masters"
  Icehockey_nhl Title = "NHL Hockey"
  Mma_mixed_martial_arts Title = "MMA"
  Soccer_usa_mls Title = "MLS Soccer"
  Tennis_atp_french_open Title = "French Open"
  Baseball_mlb Title = "MLB Baseball"
  Baseball_ncaa Title = "NCAA Baseball"
  Basketball_ncaa Title = "NCAA Basketball"
)


func  AllTitles () []Title {
  return []Title{ Americanfootball_nfl, Americanfootball_ncaaf, Basketball_nba, Basketball_ncaa, Baseball_mlb, Baseball_ncaa, Golf_masters_tournament_winner, Icehockey_nhl, Mma_mixed_martial_arts, Soccer_usa_mls, Tennis_atp_french_open }
}
