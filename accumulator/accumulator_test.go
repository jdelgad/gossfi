package accumulator

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pushWords(w []string, wdCh chan string) {
	for _, value := range w {
		wdCh <- value
	}
	close(wdCh)
}

func TestAccumulate(t *testing.T) {
	sendWords := []string{"one", "two", "three", "two"}
	wMap := make(map[string]int)
	wMap["one"] = 1
	wMap["two"] = 2
	wMap["three"] = 1

	wdCh := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go Accumulate(wdCh, &wg)

	go pushWords(sendWords, wdCh)

	wg.Wait()

	assert.Equal(t, wMap, words)
}

func TestOrderWordsByCount(t *testing.T) {
	words = make(map[string]int)
	words["hi"] = 1
	words["boom"] = 2
	words["yes"] = 5

	wc := make(WordCount, 3)
	wc[0] = WordCountPair{"yes", 5}
	wc[1] = WordCountPair{"boom", 2}
	wc[2] = WordCountPair{"hi", 1}

	wordCount := orderWordsByCount()
	assert.Equal(t, wordCount, wc)
}

func TestOrderWordsByCountEmpty(t *testing.T) {
	words = make(map[string]int)

	wc := make(WordCount, 0)

	wordCount := orderWordsByCount()
	assert.Equal(t, wordCount, wc)
}

func TestOrderWordsByCountEquality(t *testing.T) {
	words = make(map[string]int)
	words["hi"] = 1
	words["boom"] = 2
	words["ok"] = 2
	words["yes"] = 5

	wc := make(WordCount, 4)
	wc[0] = WordCountPair{"yes", 5}
	wc[1] = WordCountPair{"boom", 2}
	wc[2] = WordCountPair{"ok", 2}
	wc[3] = WordCountPair{"hi", 1}

	wordCount := orderWordsByCount()
	assert.Equal(t, wordCount, wc)
}
