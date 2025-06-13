package main

import (
	"assert/assert"
	"testing"
)

func TestAdd( t *testing.T) {
	assert.Equal("sum of 1 and 2 is 3",t, Add(1, 2), 3)

}

func TestNotAdd( t *testing.T) {
	assert.NotEqual("sum of 1 and 2 is not 4",t, Add(1, 2), 4)
}