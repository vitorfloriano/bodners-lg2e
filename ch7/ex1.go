package main

import (
	
)

type Team struct {
	teamName    string
	playerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstTeam, secondTeam string, ftScore, stScore int) {
	firstTeamExists := false
	for _, team := range l.Teams {
		if team.teamName == firstTeam {
			firstTeamExists = true
			break
		}
	}
	secondTeamExists := false
	for _, team := range l.Teams {
		if team.teamName == secondTeam {
			secondTeamExists = true
			break
		}
	}
	if !firstTeamExists || !secondTeamExists {
		fmt.Println("One or both teams are not in the league.")
		return
	}
	if ftScore > stScore {
		l.Wins[firstTeam]++
	}
	else if stScore > ftScore {
		l.Wins[secondTeam]++
	} else {
		// do nothing
	}
	}
}
	

}


func (l League) Ranking() []Team.teamName {
	
}


