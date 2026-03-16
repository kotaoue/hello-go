package main

func MakeRuneMap(i int) map[rune]int {
	return make(map[rune]int, i)
}
func Make2RuneMap(i int) map[[2]rune]int {
	return make(map[[2]rune]int, i)
}

func MakeStringMap(i int) map[string]int {
	return make(map[string]int, i)
}
