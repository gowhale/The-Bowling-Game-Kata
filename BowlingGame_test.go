package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGutterGame(t *testing.T) {
	assert.EqualValues(t, PlayGame(), 0)
}
