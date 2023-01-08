package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	var inputs []string = strings.Split(sc.Text(), " ")
	sort.Strings(inputs)
	strings.Join(inputs, " ")
	fmt.Println(inputs)
}
