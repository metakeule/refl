package main

import (
	"fmt"
	"github.com/metakeule/refl"
)

type Rating struct {
	NumParticipants int
	Points          int
}

type Game struct {
	Rating
	Name string
}

func (ø *Game) Play(players int) {
	p("playing " + ø.Name + " with " + fmt.Sprintf("%d", players) + " players")
}

func p(i interface{}) {
	fmt.Println(i)
}

func main() {
	g := Game{Rating{2, 300}, "Game of life"}
	p(refl.Type(g))              // => "Game"
	p(refl.Kind(g))              // => "struct"
	refl.Call(&g, "Play", 2)     // => "playing Game of life with 2 players"
	p(refl.GetField(&g, "Name")) // "Game of life"
	refl.SetField(&g.Rating, "Points", 400)
	p(g.Points) // 400
	refl.Set(&g.Rating, Rating{5, 800})
	p(g.Points) // 600
}