package locator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTxtFiles(t *testing.T) {
	var files []string
	fCh := make(chan string)
	go FindTextFiles("../unit_test/files", ".txt", fCh)
	for f := range fCh {
		files = append(files, f)
	}
	assert.Equal(t, len(files), 5)
}
