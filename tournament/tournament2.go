package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	win  = iota
	draw = iota
	loss = iota
	none = -1
)

// Tally function calculates the tournament results and generates a formatted results table
func Tally(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	teams := make(map[string]*team)
	getTeam := func(name string) *team {
		if team, ok := teams[name]; ok {
			return team
		}
		team := newTeam(name)
		teams[name] = &team
		return &team
	}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		nameA, nameB, result, err := newMatch(strings.Split(line, ";"))
		if err != nil {
			return err
		}
		if result != none {
			teamA, teamB := getTeam(nameA), getTeam(nameB)
			switch result {
			case win:
				teamA.update(win)
				teamB.update(loss)
			case draw:
				teamA.update(draw)
				teamB.update(draw)
			case loss:
				teamA.update(loss)
				teamB.update(win)
			}
		}
	}
	var list []*team
	for _, t := range teams {
		list = append(list, t)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].points > list[j].points ||
			(list[i].points == list[j].points && list[i].name < list[j].name)
	})
	format := func(name string, matches, wins, draws, losses, points interface{}) {
		fmt.Fprintf(out, "%-30v |%3v |%3v |%3v |%3v |%3v\n", name, matches, wins, draws, losses, points)
	}
	format("Team", "MP", "W", "D", "L", "P")
	for _, t := range list {
		format(t.name, t.matches, t.wins, t.draws, t.losses, t.points)
	}
	return nil
}

type team struct {
	name    string
	matches int
	wins    int
	draws   int
	losses  int
	points  int
}

func newTeam(name string) team {
	return team{name, 0, 0, 0, 0, 0}
}

func (t *team) update(result int) {
	switch result {
	case win:
		t.wins++
		t.points += 3
	case draw:
		t.draws++
		t.points++
	case loss:
		t.losses++
	default:
		return
	}
	t.matches++
}

type match struct {
	teamA  string
	teamB  string
	result int
}

var resultMap = map[string]int{"win": win, "draw": draw, "loss": loss}

func newMatch(src []string) (string, string, int, error) {
	if src == nil {
		return "", "", none, nil
	}
	err := errors.New("Match in wrong format")
	if len(src) != 3 {
		return "", "", none, err
	}
	teamA, teamB := src[0], src[1]
	result, ok := resultMap[src[2]]
	if teamA == teamB || !ok {
		return "", "", none, err
	}
	return teamA, teamB, result, nil
}
