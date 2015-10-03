package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTxtFile(t *testing.T) {
	words := []string{
		"just",
		"a",
		"file",
		"I",
		"added",
		"to",
		"for",
		"testing",
	}

	readInWords, err := Read("../unit_test/files/new.txt")
	assert.EqualValues(t, words, readInWords)
	assert.Nil(t, err)
}
