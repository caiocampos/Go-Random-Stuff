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
)

// Tally function calculates the tournament results and generates a formatted results table
func Tally(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanRunes)
	teams := make(map[string]*team)
	getTeam := func(name string) *team {
		if team, ok := teams[name]; ok {
			return team
		}
		team := newTeam(name)
		teams[name] = &team
		return &team
	}
	for {
		line, eof := readLine(scanner)
		match, err := newMatch(line)
		if err != nil {
			return err
		}
		if match != nil {
			teamA, teamB := getTeam(match.teamA), getTeam(match.teamB)
			switch match.result {
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
		if eof {
			break
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

func readLine(scanner *bufio.Scanner) ([]string, bool) {
	var b strings.Builder
	var res []string
	popString := func() {
		if b.Len() > 0 {
			s := b.String()
			if !strings.HasPrefix(s, "#") {
				res = append(res, b.String())
			}
		}
		b.Reset()
	}
	eof := true
	for scanner.Scan() {
		s := scanner.Text()
		if s == "\n" {
			eof = false
			break
		}
		if s == ";" {
			popString()
		} else {
			b.WriteString(s)
		}
	}
	popString()
	return res, eof
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

func newMatch(src []string) (*match, error) {
	if src == nil {
		return nil, nil
	}
	err := errors.New("Match in wrong format")
	if len(src) != 3 {
		return nil, err
	}
	teamA, teamB, result := src[0], src[1], getResult(src[2])
	if teamA == teamB || result == -1 {
		return nil, err
	}
	return &match{teamA, teamB, result}, nil
}

func getResult(info string) int {
	switch info {
	case "win":
		return win
	case "draw":
		return draw
	case "loss":
		return loss
	default:
		return -1
	}
}
