package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}

	r := regexp.MustCompile("[А-Яа-яA-Za-z0-9'*?()$.,!-:]+")
	words := r.FindAllString(text, -1)

	wordsCount := make(map[string]int)

	for _, m := range words {
		wordsCount[m]++
	}

	const countOfMostPopularWords = 10

	list := make([]string, 0, countOfMostPopularWords)

	for w := range wordsCount {
		list = append(list, w)
	}

	sort.Slice(list, func(i, j int) bool {
		return wordsCount[list[i]] > wordsCount[list[j]]
	})

	if len(list) >= countOfMostPopularWords {
		return list[:countOfMostPopularWords]
	}

	return nil
}
