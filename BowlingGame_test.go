package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGutterGame(t *testing.T) {

	var g Game

	for i := 0; i < 20; i++ {

		g = g.roll(0)

	}

	assert.EqualValues(t, 0, g.score)

}

func TestAllOnes(t *testing.T) {

	var g Game

	for i := 0; i < 20; i++ {

		g = g.roll(1)

	}

	assert.EqualValues(t, 20, g.score)

}
