package reader

import (
	"bufio"
	"os"
	"sync"
)

func GetWords(fCh, wdCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case f, ok := <-fCh:
			if ok {
				file, err := os.Open(f)
				if err != nil {
					return
				}

				scanner := bufio.NewScanner(file)
				scanner.Split(bufio.ScanWords)
				for scanner.Scan() {
					wdCh <- scanner.Text()
				}

				file.Close()
			} else {
				return
			}
		}
	}
}
