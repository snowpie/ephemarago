package main

import (
	"assert/assert"
	"fmt"
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

func TestStachSecret(t *testing.T) {
	fmt.Println(Secrets)
	key1 := stashsecret("first secret")
	key2 := stashsecret("second secret")
	fmt.Println(Secrets)

	assert.NotEqual("first key should not equal second secret", t, key1, key2)

	assert.Equal("first secret retrieved correctly", t, getsecret(key1), "first secret")
	assert.Equal("second secret retrieved correctly", t, getsecret(key2), "second secret")
	
	assert.Equal("first secret stored in backend", t, Secrets[key1](), "first secret")
	assert.Equal("second secret stored in backend", t, Secrets[key2](), "second secret")
	
	
}