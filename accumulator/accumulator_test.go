package accumulator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func pushWords(w []string, wdCh chan string) {
	for _, value := range w {
		fmt.Printf("sending word = %s\n", value)
		wdCh <- value
	}
}

func TestAccumulate(t *testing.T) {
	sendWords := []string{"one", "two", "three", "two"}
	wMap := make(map[string]int)
	wMap["one"] = 1
	wMap["two"] = 2
	wMap["three"] = 1

	wdCh := make(chan string)
	closeCh := make(chan bool)

	go Accumulate(wdCh, closeCh)
	pushWords(sendWords, wdCh)

	closeCh <- true
	close(closeCh)
	close(wdCh)

	assert.Equal(t, wMap, words)
}
