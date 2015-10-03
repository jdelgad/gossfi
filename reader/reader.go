package reader

import (
	"bufio"
	"os"
)

func GetWords(f string) ([]string, error) {
	var words []string
	file, err := os.Open(f)
	if err != nil {
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
