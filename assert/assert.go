package assert
// https://dev.to/yawaramin/why-i-dont-use-a-third-party-assertion-library-in-go-unit-tests-1gak
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

import (
	"fmt"
	"testing"
)

func Equal[V comparable](name string, t *testing.T, got, expected V) {
    t.Helper()
    if expected != got {
        t.Errorf(`assert.Equal( t, got: %v , expected: %v)`, got, expected)
    }
    fmt.Printf(name+" assert.Equal got: %v , expected: %v\n", got, expected)

}

func NotEqual[V comparable](name string, t *testing.T, got, expected V) {
    t.Helper()
    // fmt.Println(name)
    if expected == got {
        t.Errorf(`assert.NotEqual( t, got: %v , expected: %v)`, got, expected)
    }
    fmt.Printf(name+" assert.NotEqual got: %v , expected: %v\n", got, expected)

}
