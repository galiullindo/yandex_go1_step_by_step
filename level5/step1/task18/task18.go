package main

func DeleteVowels(s string) string {
	new := make([]rune, 0)
	for _, letter := range []rune(s) {
		switch letter {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
		default:
			new = append(new, letter)
		}
	}
	return string(new)
}
