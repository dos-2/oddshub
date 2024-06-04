package colors

import (
	"oddshub/models"
	"oddshub/sports"
)

var ColorsMap = map[string]map[string]models.TeamColors{
	string(sports.Americanfootball_ncaaf):         NCAATeamColorMap,
	string(sports.Americanfootball_nfl):           NFLTeamColorMap,
	string(sports.Basketball_nba):                 NBATeamColorMap,
	string(sports.Golf_masters_tournament_winner): {},
	string(sports.Icehockey_nhl):                  NHLTeamColorMap,
	string(sports.Mma_mixed_martial_arts):         {},
	string(sports.Soccer_usa_mls):                 MLSTeamColorMap,
	string(sports.Tennis_atp_french_open):         {},
	string(sports.Tennis_wta_french_open):         {},
	string(sports.Baseball_mlb):                   MLBTeamColorMap,
	string(sports.Baseball_ncaa):                  NCAATeamColorMap,
	//string(sports.Basketball_ncaa): NcaaTeamColorMap,
	string(sports.Boxing): {},
}
