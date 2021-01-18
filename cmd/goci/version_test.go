package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTagByFile(t *testing.T) {
	tag := "1.0.0"
	ioutil.WriteFile("./.tags", []byte(tag), 0644)
	val, err := getTag()
	os.Remove("./.tags")
	require.NoError(t, err)
	assert.ObjectsAreEqual(tag, val)
}
