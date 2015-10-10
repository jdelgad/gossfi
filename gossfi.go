package main

import (
	"github.com/jdelgad/gossfi/accumulator"
	"github.com/jdelgad/gossfi/locator"
	"github.com/jdelgad/gossfi/reader"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fileCh := make(chan string)
	wordCh := make(chan string)

	wg.Add(3)
	go accumulator.Accumulate(wordCh, &wg)
	go reader.GetWords(fileCh, wordCh, &wg)
	go locator.FindTextFiles("unit_test", ".txt", fileCh, &wg)

	wg.Wait()

	accumulator.PrintTopWords(10)
}
