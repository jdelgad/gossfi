package main

import (
	"sync"

	"github.com/jdelgad/gossfi/accumulator"
	"github.com/jdelgad/gossfi/locator"
	"github.com/jdelgad/gossfi/reader"
)

func main() {
	var wg sync.WaitGroup
	fileCh := make(chan string)
	wordCh := make(chan string)

	wg.Add(2)
	go accumulator.Accumulate(wordCh, &wg)
	go locator.FindTextFiles("unit_test", ".txt", fileCh, &wg)

	var rwg sync.WaitGroup
	rwg.Add(3)
	go reader.GetWords(fileCh, wordCh, &rwg)
	go reader.GetWords(fileCh, wordCh, &rwg)
	go reader.GetWords(fileCh, wordCh, &rwg)
	rwg.Wait()
	close(wordCh)

	wg.Wait()

	accumulator.PrintTopWords(10)
}
