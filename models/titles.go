package models

type Title string

const (
	Americanfootball_ncaaf  Title = "NCAA Football"
	Americanfootball_nfl    Title = "NFL"
	Basketball_nba          Title = "NBA Basketball"
	Golf_masters_tournament Title = "The Masters"
	Icehockey_nhl           Title = "NHL Hockey"
	Mma_mixed_martial_arts  Title = "MMA"
	Soccer_usa_mls          Title = "MLS Soccer"
	Tennis_atp_french_open  Title = "Mens French Open"
	Tennis_wta_french_open  Title = "Womens French Open"
	Baseball_mlb            Title = "MLB Baseball"
	Baseball_ncaa           Title = "NCAA Baseball"
	Basketball_ncaa         Title = "NCAA Basketball"
	Boxing                  Title = "Boxing"
)

func AllTitles() []Title {
	return []Title{Americanfootball_nfl, Americanfootball_ncaaf, Basketball_nba, Basketball_ncaa, Baseball_mlb, Baseball_ncaa, Boxing, Golf_masters_tournament, Icehockey_nhl, Mma_mixed_martial_arts, Soccer_usa_mls, Tennis_atp_french_open, Tennis_wta_french_open}
}
