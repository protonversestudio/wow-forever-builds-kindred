package main

// Kindred: rank all nine WoW Forever classes against four answers.

import (
	"encoding/json"
	"fmt"
	"os"
)

type Answers struct {
	Solo   int `json:"solo"`   // 0..2: solo, duo, group
	Speed  int `json:"speed"`  // 0..2: slow, steady, fast
	Risk   int `json:"risk"`   // 0..2: safe, mixed, risky
	Active int `json:"active"` // 0..2: relaxed, normal, intense
}

type Class struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Solo  []int   `json:"solo"`
	Speed []int   `json:"speed"`
	Risk  []int   `json:"risk"`
	Base  float64 `json:"base"`
}

type Ranked struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Score  float64 `json:"score"`
	Why    string  `json:"why"`
}

var classes = []Class{
	{"warrior", "Warrior", "#C79C6E", []int{2, 1, 0}, []int{0, 1, 2}, []int{0, 1, 2}, 6.2},
	{"paladin", "Paladin", "#F58CBA", []int{2, 2, 1}, []int{0, 1, 1}, []int{0, 0, 1}, 6.4},
	{"hunter", "Hunter", "#ABD473", []int{2, 2, 1}, []int{2, 2, 2}, []int{1, 1, 2}, 7.6},
	{"rogue", "Rogue", "#FFF569", []int{2, 0, 0}, []int{2, 2, 1}, []int{2, 2, 2}, 6.0},
	{"priest", "Priest", "#FFFFFF", []int{1, 1, 2}, []int{0, 1, 1}, []int{0, 0, 0}, 6.8},
	{"shaman", "Shaman", "#0070DE", []int{1, 2, 2}, []int{1, 1, 2}, []int{1, 1, 1}, 7.0},
	{"mage", "Mage", "#69CCF0", []int{2, 1, 2}, []int{2, 2, 2}, []int{2, 2, 1}, 7.4},
	{"warlock", "Warlock", "#9482C9", []int{2, 2, 2}, []int{1, 2, 2}, []int{2, 2, 1}, 7.2},
	{"druid", "Druid", "#FF7D0A", []int{2, 2, 2}, []int{1, 2, 2}, []int{1, 1, 2}, 7.0},
}

func rank(a Answers) []Ranked {
	out := make([]Ranked, 0, len(classes))
	for _, c := range classes {
		s := c.Base + float64(c.Solo[a.Solo]+c.Speed[a.Speed]+c.Risk[a.Risk]) * 0.9
		why := whyFor(c, a)
		out = append(out, Ranked{c.ID, c.Name, s, why})
	}
	// insertion sort, stable, no imports needed
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j].Score > out[j-1].Score; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}

func whyFor(c Class, a Answers) string {
	switch {
	case c.Solo[a.Solo] == 2 && a.Solo < 2:
		return "holds its own without a group"
	case c.Speed[a.Speed] == 2 && a.Speed == 2:
		return "clears quests at your pace and faster"
	case c.Risk[a.Risk] >= 1 && a.Risk == 2:
		return "rewards the risks you like taking"
	default:
		return "a steady fit for the way you play"
	}
}

func main() {
	var a Answers
	if len(os.Args) > 1 {
		if err := json.Unmarshal([]byte(os.Args[1]), &a); err != nil {
			fmt.Fprintln(os.Stderr, "answers JSON expected")
			os.Exit(1)
		}
	} else {
		a = Answers{1, 1, 1, 1}
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(rank(a))
}
