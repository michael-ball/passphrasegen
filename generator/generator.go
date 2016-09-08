package generator

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// RandomWords picks numberOfWords random words from the file located at
// wordsFile
func RandomWords(wordsFile string, numberOfWords int) ([]string, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	words, err := readLines(wordsFile)

	if err != nil {
		return nil, err
	}

	phraseWords := make([]string, numberOfWords)

	for index := 0; index < numberOfWords; index++ {
		randInt := rand.Intn((len(words) - 1))

		phraseWords[index] = words[randInt]
	}

	return phraseWords, nil
}
