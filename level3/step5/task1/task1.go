package main

import (
	"fmt"
	"sort"
	"strings"
)

var memo = make(map[string]string)

func splitText(text string) []string {
	words := make([]string, 0)
	builder := strings.Builder{}

	for _, letter := range []rune(text) {
		switch letter {
		case ' ', '.', ',', '!', '?':
			word := builder.String()
			if word != "" {
				words = append(words, word)
				builder.Reset()
			}

		default:
			builder.WriteRune(letter)
		}
	}

	word := builder.String()
	if word != "" {
		words = append(words, word)
	}

	return words
}

func getTopWords(wordMap map[string]int, n int) []string {
	topWords := make([]string, 0, len(wordMap))

	values := make([]int, 0, len(wordMap))
	for _, value := range wordMap {
		values = append(values, value)
	}

	sort.Ints(values)

	invertedWordMap := make(map[int]string)
	for key, value := range wordMap {
		invertedWordMap[value] = key
	}

	length := len(values)
	for i := length - 1; i >= length-n; i-- {
		topWords = append(topWords, invertedWordMap[values[i]])
	}

	return topWords
}

func AnalyzeText(text string) {
	text = strings.ToLower(text)

	message, ok := memo[text]
	if !ok {
		words := splitText(text)
		countWords := len(words)

		wordMap := make(map[string]int)
		for _, word := range words {
			wordMap[word]++
		}

		countUniqueWords := len(wordMap)
		topWords := getTopWords(wordMap, 5)

		message += fmt.Sprintf("Количество слов: %d\n", countWords)
		message += fmt.Sprintf("Количество уникальных слов: %d\n", countUniqueWords)
		message += fmt.Sprintf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", topWords[0], wordMap[topWords[0]])
		message += "Топ-5 самых часто встречающихся слов:\n"

		for _, word := range topWords {
			message += fmt.Sprintf("\"%s\": %d раз\n", word, wordMap[word])
		}

		memo[text] = message
	}

	fmt.Print(message)
}
