package main

import (
	"fmt"
	"sort"
	"os"
	"io"	
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
	} else if stScore > ftScore {
		l.Wins[secondTeam]++
	} else {
		// do nothing
	}
}

func (l *League) Ranking() []string {
	teamNames := []string{}
	for team := range l.Wins {
		teamNames = append(teamNames, team)
	}
	sort.Slice(teamNames, func(i, j int) bool {
		return l.Wins[teamNames[i]] > l.Wins[teamNames[j]]
	})
	return teamNames		
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	rankings := r.Ranking()
	for _, team := range rankings {
		io.WriteString(w, team+"\n")
	}
}

func main() {
	league := League{
		Teams: []Team{
			{teamName: "Raptors", playerNames: []string{"Alice", "Bob"}},
			{teamName: "Lakers", playerNames: []string{"Charlie", "Dana"}},
		},
		Wins: map[string]int{
			"Raptors": 5,
			"Lakers": 3,
		},
	}

	league.MatchResult("Raptors", "Lakers", 102, 98)
	
	fmt.Println("Rankings: ")
	RankPrinter(&league, os.Stdout)
}	
