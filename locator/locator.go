package locator

import (
	"os"
	"path/filepath"
)

var ext string
var inputFileCh chan string

func isTextFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if !fi.IsDir() && filepath.Ext(fp) == ext {
		inputFileCh <- fp
	}

	return nil
}

func FindTextFiles(d, e string, fileCh chan string) {
	ext = e
	inputFileCh = fileCh
	filepath.Walk(d, isTextFile)
	close(inputFileCh)
}
