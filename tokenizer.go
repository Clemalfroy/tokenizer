package tokenizer

import (
	"strings"

	unidecode "github.com/mozillazg/go-unidecode"
)

func Tokenize(str string) []string {
	str = unidecode.Unidecode(str)
	runes := []rune(str)
	lenRunes := len(runes)
	for i, r := range runes {
		switch r {
		case 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90:
			runes[i] = r + 32
		case ')', '(', '}', '{', ']', '[', '"', '`', '>', '<', '\t', '\n', '-', '/', '\r':
			runes[i] = ' '
		case '\'':
			runes[i] = ' '
			charBeforeIsletter := i > 0 && runes[i-1] >= 'a' && runes[i-1] <= 'z'
			letterStartsNewWord := i-1 == 0 || (i > 1 && runes[i-2] == ' ')
			if charBeforeIsletter && letterStartsNewWord {
				runes[i-1] = ' '
			}
		case '.', ',', ';', ':', '?', '!', '@', '#', '$', '%', '&', '=', '+', '|', '^', '_', '~':
			if (lenRunes > i+1 && runes[i+1] == ' ') || lenRunes == i+1 {
				runes[i] = ' '
			}
		}
	}
	return strings.Fields(string(runes))
}
