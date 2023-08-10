package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordFields := strings.Fields(s)
	count := map[string]int{}
	for i := range wordFields {
		word := wordFields[i]
		v, ok := count[word]
		if ok {
			count[word] = v + 1
		} else {
			count[word] = 1
		}
	}
	fmt.Println(count)
	return count
}

func main() {
	wc.Test(WordCount)
}
