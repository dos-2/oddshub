package sport 

type Sport string

const(
  Americanfootball_ncaaf Sport = "americanfootball_ncaaf" 
  Americanfootball_nfl Sport = "americanfootball_nfl"
  Basketball_nba Sport = "basketball_nba"
  Golf_masters_tournament_winner Sport = "Golf_masters_tournament_winner "
  Icehockey_nhl Sport = "icehockey_nhl"
  Mma_mixed_martial_arts Sport = "mma_mixed_martial_arts"
  Soccer_usa_mls Sport = "soccer_usa_mls"
  Tennis_atp_french_open Sport = "tennis_atp_french_open"
  Baseball_mlb Sport = "baseball_mlb"
  Baseball_ncaa Sport = "baseball_ncaa"
  Basketball_ncaa Sport = "basketball_ncaa"
)


func  AllSports () []Sport {
  return []Sport{ Americanfootball_nfl, Americanfootball_ncaaf, Basketball_nba, Basketball_ncaa, Baseball_mlb, Baseball_ncaa, Golf_masters_tournament_winner, Icehockey_nhl, Mma_mixed_martial_arts, Soccer_usa_mls, Tennis_atp_french_open }
}
