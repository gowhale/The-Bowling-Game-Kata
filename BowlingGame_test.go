package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var g Game

func (g *Game) rollMany(n int, pins int) {
	for i := 0; i < n; i++ {

		g.roll(pins)

	}
}

func TestGutterGame(t *testing.T) {

	g.rollMany(20, 0)

	assert.EqualValues(t, 0, g.score)

}

func TestAllOnes(t *testing.T) {

	g.rollMany(20, 1)

	assert.EqualValues(t, 20, g.score)

}
