package locator
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestTxtFiles(t *testing.T) {
	f := FindTextFiles("../unit_test/files", ".txt")
	fmt.Printf("%s\n", f)
	assert.Equal(t, len(f), 5)
}
