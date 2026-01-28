package main

import (
	"testing"

	reference "github.com/galiullindo/yandex_go1_step_by_step/level3/step5/task1/v2/internal"
)

func TestSolution(t *testing.T) {
	var (
		test string
		want string
		got  string
	)

	test = "one two two three three three four four four four five five five five five"
	want = reference.AnalyzeText(test)
	got = AnalyzeText(test)
	if got != want {
		t.Fatalf("expected:\n%s\ngot\n%s\n", want, got)
	}

	test = "Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!"
	want = reference.AnalyzeText(test)
	got = AnalyzeText(test)
	if got != want {
		t.Fatalf("expected:\n%s\ngot\n%s\n", want, got)
	}

	test = "Я так люблю море. Я на море. Я так люблю. Море! Я море!!! ЛЮБЛЮ МОРЕ"
	want = reference.AnalyzeText(test)
	got = AnalyzeText(test)
	if got != want {
		t.Fatalf("expected:\n%s\ngot\n%s\n", want, got)
	}
}
