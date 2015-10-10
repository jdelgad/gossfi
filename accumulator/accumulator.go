package accumulator

import (
	"fmt"
	"sort"
)

type WordCountPair struct {
	Word  string
	Count int
}

type WordCount []WordCountPair

func (wc WordCount) Len() int {
	return len(wc)
}

func (wc WordCount) Less(i, j int) bool {
	return wc[i].Count < wc[j].Count
}

func (wc WordCount) Swap(i, j int) {
	wc[i], wc[j] = wc[j], wc[i]
}

var words map[string]int

func Accumulate(wdCh chan string, closeCh chan bool) {
	words = make(map[string]int)
	for {
		select {
		case word := <-wdCh:
			fmt.Printf("word = %s\n", word)
			words[word]++
		case <-closeCh:
			return
		}
	}
}

func OrderWordsByCount() WordCount {
	w := make(WordCount, len(words))
	i := 0
	for word, count := range words {
		w[i] = WordCountPair{word, count}
		i++
	}
	sort.Sort(sort.Reverse(w))
	return w
}
