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
		"testing",
		"testing",
		"testing",
	}

	readInWords, err := GetWords("../unit_test/files/new.txt")
	assert.EqualValues(t, words, readInWords)
	assert.Nil(t, err)
}
