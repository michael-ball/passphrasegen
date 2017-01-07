package main

import (
	"flag"
	"fmt"
	"strings"

	generator "github.com/michael-ball/passphrasegen/generator"
)

var numWords = flag.Int("w", 4, "Number of words in passphrase")

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()

	wordsFile := "/usr/share/dict/words"

	if len(args) > 0 {
		wordsFile = args[0]
	}

	phraseWords, err := generator.RandomWords(wordsFile, *numWords)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(strings.Join(phraseWords, " "))
}
