package main

import "fmt"

type sequence []int
type counter map[int]int

const DUP = 2
const MISS = 0

func CountNums(seq sequence) counter {
	c := make(map[int]int)
	// To define missing values
	for i := 1; i < len(seq)+1; i++ {
		c[i] = 0
	}
	for _, num := range seq {
		c[num]++
	}
	return c
}

func getMissNum(c counter) (k int, ok bool) {
	for k, v := range c {
		if v == MISS {
			return k, true
		}
	}
	return

}

func getDupNum(c counter) (k int, ok bool) {
	for k, v := range c {
		if v == DUP {
			return k, true
		}
	}
	return
}

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
	c := CountNums(seq)
	dupNum, _ := getDupNum(c)
	missNum, ok := getMissNum(c)
	if !ok {
		fmt.Println("Collect")
	} else {
		fmt.Printf("%d %d\n", dupNum, missNum)
	}
}
