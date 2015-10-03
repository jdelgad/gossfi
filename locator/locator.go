package locator
import (
	"path/filepath"
	"os"
)

var glob string
var txtFiles []string

func isTextFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if !fi.IsDir() && filepath.Ext(fp) == glob {
		txtFiles = append(txtFiles, fp)
		return nil
	}

	return nil
}

func FindTextFiles(d, g string) []string {
	glob = g
	filepath.Walk(d, isTextFile)
	return txtFiles
}
