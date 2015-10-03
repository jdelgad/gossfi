package reader

import (
	"bufio"
	"os"
	"fmt"
)

func GetWords(fCh, wdCh chan string, closeCh chan bool) {
	for {
		select {
		case f := <- fCh:
			fmt.Printf("Parsing %s", f)
			file, err := os.Open(f)
			if err != nil {
				return
			}

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				fmt.Printf("Putting %s on wdCh", scanner.Text())
				wdCh <- scanner.Text()
			}

			file.Close()
		case <- closeCh:
			close(wdCh)
			return
		}
	}
}
