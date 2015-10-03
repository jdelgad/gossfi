package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

var readInWords []string

func readWordChannel(wdCh chan string) {
	for word := range wdCh {
		fmt.Printf("looping in wdCh")
		readInWords = append(readInWords, word)
	}
}

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

	file := "../unit_test/files/new.txt"
	fCh := make(chan string)
	wdCh := make(chan string)
	closeCh := make(chan bool)

	go readWordChannel(wdCh)
	go GetWords(fCh, wdCh, closeCh)

	fCh <- file
	closeCh <- true


	assert.EqualValues(t, words, readInWords)
}
