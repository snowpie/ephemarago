package main

import (
	"assert/assert"
	
	"testing"
)
func TestSecret(t *testing.T) {
	secrets := make(map[string]func() string)
	secrets["first"] = secret("first secret")
	secrets["second"] =	secret("second secret")
	assert.Equal("first secret", t, secrets["first"](), "first secret")
	assert.Equal("second secret", t, secrets["second"](), "second secret")
	assert.NotEqual("first secret should not equal second secret", t, secrets["first"](), secrets["second"]())
	assert.NotEqual("second secret should not equal first secret", t, secrets["second"](), secrets["first"]())
}