package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	dict := make(map[string]int)

	for _, word := range strings.Fields(text) {
		dict[word] = dict[word] + 1

	}

	words := make([]string, 0, len(dict))
	for word := range dict {
		words = append(words, word)

	}

	sort.Slice(words, func(i, j int) bool {
		if dict[words[i]] == dict[words[j]] {
			return words[i] < words[j]
		}
		return dict[words[i]] > dict[words[j]]
	})

	if len(words) > 10 {
		words = words[:10]
	}
	return words
}
