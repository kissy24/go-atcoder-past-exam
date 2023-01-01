package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Doubling a numeric string.

Args:
	s(string): Various string

Return:
	(int): Doubled number
	(error): String cannot be converted to int.
*/
func ToDoubledInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	return val * 2, err
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	doubledNum, err := ToDoubledInt(s)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(doubledNum)
	}
}
