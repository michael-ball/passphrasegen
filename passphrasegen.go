package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var numWords = flag.Int("w", 4, "Number of words in passphrase")

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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	flag.Parse()

	args := flag.Args()

	wordsFile := "/usr/share/dict/words"

	if len(args) > 0 {
		wordsFile = args[0]
	}

	words, err := readLines(wordsFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	phraseWords := make([]string, *numWords)

	for index := 0; index < *numWords; index++ {
		randInt := rand.Intn((len(words) - 1))

		phraseWords[index] = words[randInt]
	}

	fmt.Println(strings.Join(phraseWords, " "))
}
