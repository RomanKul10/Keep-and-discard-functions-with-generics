// test_main.go
package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterSliceString(t *testing.T) {
	slice := []string{"cow", "apple", "dog", "ananas", "anana"}
	filter := NewFilteredSlice(slice, func(v string) bool {
		return strings.HasPrefix(v, "a")
	})

	keepString := filter.Keep()
	discardString := filter.Discard()

	keep := []string{"apple", "ananas", "anana"}
	discard := []string{"cow", "dog"}

	assert.Equal(t, keep, keepString)
	assert.Equal(t, discard, discardString)
}

func TestFilterSliceInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 8}
	filter := NewFilteredSlice(slice, func(v int) bool {
		return v%2 == 1
	})

	keepInt := filter.Keep()
	discardInt := filter.Discard()

	keep := []int{1, 3}
	discard := []int{2, 4, 8}

	assert.Equal(t, keep, keepInt)
	assert.Equal(t, discard, discardInt)
}
