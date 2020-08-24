package main

// import "fmt"

type Game struct {
	score int
}

func (g Game) roll(p int) Game {

	g.score += p

	return g
}
