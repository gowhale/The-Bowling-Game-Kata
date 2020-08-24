package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func (g *Game) rollMany(n int, pins int) {
	for i := 0; i < n; i++ {
		g.roll(pins)
	}
}

func TestGutterGame(t *testing.T) {

	var g Game
	g.rollMany(20, 0)
	assert.EqualValues(t, 0, g.score)
}

func TestAllOnes(t *testing.T) {

	var g Game
	g.rollMany(20, 1)
	assert.EqualValues(t, 20, g.score)
}

// func TestOneSpare(t *testing.T) {
// 	var g Game

// 	g.roll(5)
// 	g.roll(5)
// 	g.roll(3)
// 	g.rollMany(17, 0)
// 	assert.EqualValues(t, 16, g.score)
// }
