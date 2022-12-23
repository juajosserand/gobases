package main

import "fmt"

func main() {
	var w string = "word"

	lenWord(w)
	spellOutWord(w)
}

func lenWord(w string) {
	fmt.Println(len(w))
}

func spellOutWord(w string) {
	for _, c := range w {
		fmt.Println(string(c))
	}
}
