package accumulator
import "fmt"

var words map[string]int

func Accumulate(wdCh chan string, closeCh chan bool) {
	words = make(map[string]int)
	for {
		select {
		case word := <- wdCh:
			if word != "" {
				fmt.Printf("word = %s\n", word)
				words[word]++
			}
		case <- closeCh:
			return
		}
	}
}
