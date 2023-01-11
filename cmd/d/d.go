package main

import "fmt"

type sequence []int

func inputSeq() sequence {
	var period int
	fmt.Scan(&period)
	seq := make([]int, period)
	for i := 0; i < period; i++ {
		fmt.Scan(&seq[i])
	}
	return seq
}

func main() {
	seq := inputSeq()
}
