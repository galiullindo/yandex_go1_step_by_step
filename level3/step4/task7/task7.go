package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	swap := make(map[string]string)
	for key, value := range m {
		swap[value] = key
	}
	return swap
}
