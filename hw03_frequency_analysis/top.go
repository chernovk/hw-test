package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(text string) []string {
	dict := make(map[string]int)

	for _, word := range strings.Fields(text) {
		dict[word]++
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

func Top10Extra(text string) []string {
	dict := make(map[string]int)

	re := regexp.MustCompile(`^\p{P}*|\p{P}*$`)

	for _, word := range strings.Fields(text) {
		word = re.ReplaceAllString(word, "")
		if word != "" {
			word = strings.ToLower(word)
			dict[word]++
		}
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
