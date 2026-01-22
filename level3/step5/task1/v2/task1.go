package main

import (
	"fmt"
	"sort"
	"strings"
)

var memo = make(map[string]string)

func splitText(s string, sep string) []string {
	split := func(r rune) bool {
		return strings.Contains(sep, string(r))
	}
	return strings.FieldsFunc(s, split)
}

func getWords(parts []string) []string {
	words := make([]string, len(parts))
	for i, part := range parts {
		if part != "" {
			words[i] = part
		}
	}
	return words
}

func getWordMap(words []string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	return wordMap
}

func getTopWords(wordMap map[string]int, n int) []string {
	topWords := make([]string, n)

	values := make([]int, 0, len(wordMap))
	for _, value := range wordMap {
		values = append(values, value)
	}

	sort.Ints(values)
	end := len(values) - 1
	for i := end; i > end-n; i-- {
		for key, value := range wordMap {
			if values[i] == value {
				topWords[end-i] = key
				break
			}
		}
	}

	return topWords
}

func AnalyzeText(text string) string {
	lower := strings.ToLower(text)

	if message, ok := memo[lower]; ok {
		return message
	}

	parts := splitText(lower, " .,!?")
	words := getWords(parts)
	wordMap := getWordMap(words)
	topWords := getTopWords(wordMap, 5)

	numberOfWords := len(words)
	numberOfUniqueWords := len(wordMap)

	message := ""
	message += fmt.Sprintf("Количество слов: %d\n", numberOfWords)
	message += fmt.Sprintf("Количество уникальных слов: %d\n", numberOfUniqueWords)
	message += fmt.Sprintf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", topWords[0], wordMap[topWords[0]])
	message += "Топ-5 самых часто встречающихся слов:\n"

	for _, topWord := range topWords {
		message += fmt.Sprintf("\"%s\": %d раз\n", topWord, wordMap[topWord])
	}

	memo[lower] = message
	return message
}
