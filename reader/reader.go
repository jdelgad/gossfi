package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var fs fileSystem = osFS{}

type fileSystem interface {
	Open(name string) (file, error)
}

type file interface {
	io.Closer
	io.Reader
}

type osFS struct{}

func (osFS) Open(name string) (file, error) {
	return os.Open(name)
}

func Read(f string) ([]string, error) {
	var words []string
	file, err := os.Open(f)
	if err != nil {
		fmt.Printf("could not open file")
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}
