package generator

import "testing"

func TestRandomWords(t *testing.T) {
	words, err := RandomWords("/usr/share/dict/words", 4)

	if err != nil {
		t.Errorf("%s", err)
	}

	if len(words) != 4 {
		t.Errorf("RandomWords does not return the correct number of words")
	}
}
