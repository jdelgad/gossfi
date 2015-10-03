package locator
import (
	"testing"
	"github.com/stretchr/testify/assert"
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
