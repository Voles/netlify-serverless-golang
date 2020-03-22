package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetItems(t *testing.T) {
	actual := getItems()
	expected := []item{
		{Title: "First Item"},
		{Title: "Second Item"},
		{Title: "Third Item"},
	}

	assert.Equal(t, expected, actual)
}
