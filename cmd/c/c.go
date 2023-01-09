package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type sequence []int

func ToSeq(strSlice []string) (seq sequence, err error) {
	for _, s := range strSlice {
		var v int
		v, err = strconv.Atoi(s)
		if err != nil {
			return
		}
		seq = append(seq, v)
	}
	return
}

func main() {
	sc.Scan()
	var inputs []string = strings.Split(sc.Text(), " ")
	seq, err := ToSeq(inputs)
	if err != nil {
		fmt.Println(err)
	}
	sort.Ints(seq)
	fmt.Println(seq[3])
}
