refl
====

syntactic sugar for reflection in go

I found it hard to remember the way reflection in go works, so I made some helpers.

No special error handling, panicking where possible (as in the reflect package), and for the rest you get the errors from reflect.

Use it at your own risk.

examples
--------

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

	type A1 struct {
		Inner B1
	}

	type A2 struct {
		Inner B2
	}

	type B1 struct{}
	type B2 struct{}

	func SetInner(a interface{}, b interface{}) {
		refl.PolySetField(a, "Inner", b)
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
		p(g.Points) // 800

		a1, a2 := &A1{}, &A2{}
		SetInner(a1, B1{})
		SetInner(a2, B2{})
		fmt.Println(refl.Inspect(a1)) // &main.A1{Inner:main.B1{}} ()
		fmt.Println(refl.Inspect(a2)) // &main.A2{Inner:main.B2{}} ()
	}
