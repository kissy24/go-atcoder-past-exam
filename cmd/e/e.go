package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type States struct {
}

func Follow() {

}

func FollowFollow() {

}

func FollowBackAll() {

}

var sc = bufio.NewScanner(os.Stdin)

func parse() (int, int, [][]int) {
	sc.Scan()
	c := strings.Split(sc.Text(), " ")
	usr_c, err := strconv.Atoi(c[0])
	if err != nil {
		fmt.Println(err)
	}
	log_c, err := strconv.Atoi(c[1])
	if err != nil {
		fmt.Println(err)
	}
	logs := make([][]int, log_c)
	for i := 0; i < log_c; i++ {
		sc.Scan()
		log := ToInts(strings.Split(sc.Text(), " "))
		logs[i] = log
	}
	return usr_c, log_c, logs

}

func ToInts(s []string) []int {
	var i []int
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		i = append(i, n)
	}
	return i
}

func main() {
	_, _, _ := parse()
}
