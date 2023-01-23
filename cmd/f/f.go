package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func SortByDoubleCamel(s string) string {
	words := regexp.MustCompile(`[A-Z][a-z]*[A-Z]`).FindAllString(s, -1)
	sort.Slice(words, func(i, j int) bool {
		return strings.ToLower(words[i]) < strings.ToLower(words[j])
	})
	return strings.Join(words, "")
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	words := SortByDoubleCamel(s)
	fmt.Println(words)
}
