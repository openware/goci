package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	scenarios := []struct {
		name   string
		input  []string
		output []string
	}{
		{"unique values", []string{"example/values.yaml", "example/versions.yaml"}, []string{"example"}},
		{"non-unique values", []string{"example1/test.yaml", "example2/test.yaml"}, []string{"example1", "example2"}},
		{"unique and non-unique values", []string{"example1", "example1", "example2"}, []string{"example1", "example2"}},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, filterUnique(test.input, 1))
		})
	}
}
