package main

// import "fmt"

type Game struct {
	score int
	rolls [int]
}

func (g *Game) roll(p int) {

	g.score += p

}
