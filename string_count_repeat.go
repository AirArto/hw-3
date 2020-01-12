package stringcount

import (
	"sort"
	"unicode"
)

type Word struct {
	text  string
	count int
}

// Do for finding top 10 repeated substrings
func Do(input *string) (output []string, err error) {
	inputStr := *input
	inputChars := []rune(inputStr)
	if len(inputStr) != 0 {
		output = parceString(&inputChars)
	}

	return
}

func parceString(inputChars *[]rune) (output []string) {
	wordSlice := []Word{}
	wordsStruct := map[string]int{}
	currentWordRunes := []rune{}
	isLastLetter := false
	for i, tempRune := range *inputChars {
		if unicode.IsLetter(tempRune) {
			currentWordRunes = append(currentWordRunes, unicode.ToLower(tempRune))
			isLastLetter = true
		}
		if (!unicode.IsLetter(tempRune) && isLastLetter) || (i == len(*inputChars)-1 && unicode.IsLetter(tempRune)) {
			wordsStruct[string(currentWordRunes)]++
			currentWordRunes = []rune{}
			isLastLetter = false
		}
	}
	for i, j := range wordsStruct {
		tempWord := Word{text: i, count: j}
		wordSlice = append(wordSlice, tempWord)
	}
	sort.Slice(wordSlice, func(i, j int) bool {
		return wordSlice[i].count > wordSlice[j].count
	})
	for _, j := range wordSlice[0:10] {
		output = append(output, j.text)
	}
	return
}
