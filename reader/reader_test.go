package reader

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var readInWords []string

func readWordChannel(wdCh chan string) {
	for word := range wdCh {
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

	var wg sync.WaitGroup
	wg.Add(1)

	go readWordChannel(wdCh)
	go GetWords(fCh, wdCh, &wg)

	fCh <- file
	close(fCh)

	wg.Wait()
	close(wdCh)

	assert.EqualValues(t, words, readInWords)
}
