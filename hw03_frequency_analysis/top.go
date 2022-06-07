package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var r = regexp.MustCompile("[А-Яа-яA-Za-z0-9'*?()$.,!-:]+")

type wordsData struct {
	word  string
	count int
}

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}

	countOfMostPopularWords := 10
	words := r.FindAllString(text, -1)

	wordsCount := make(map[string]int)
	list := make([]wordsData, 0, len(wordsCount))
	result := make([]string, 0, countOfMostPopularWords)

	for _, m := range words {
		wordsCount[m]++
	}

	for word, count := range wordsCount {
		list = append(list, wordsData{word, count})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].count > list[j].count {
			return true
		}
		if list[i].count < list[j].count {
			return false
		}
		return list[i].word < list[j].word
	})

	if len(list) < countOfMostPopularWords {
		countOfMostPopularWords = len(list)
	}
	for i := 0; i < countOfMostPopularWords; i++ {
		result = append(result, list[i].word)
	}

	return result
}
