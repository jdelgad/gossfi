package locator

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestTxtFiles(t *testing.T) {
	var wg sync.WaitGroup
	var files []string
	fCh := make(chan string)

	wg.Add(1)
	go FindTextFiles("../unit_test/files", ".txt", fCh, &wg)
	for f := range fCh {
		files = append(files, f)
	}
	wg.Wait()
	assert.Equal(t, len(files), 5)
}
