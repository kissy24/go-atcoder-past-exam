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

func Follow(follow [][]bool, a int, b int) {
	follow[a-1][b-1] = true
}

func FollowFollow(follow [][]bool, a int, n int) {
	toFollow := make([]int, 0)
	for j := 0; j < n; j++ {
		if follow[a-1][j] {
			for k := 0; k < n; k++ {
				if follow[j][k] && k != a-1 {
					toFollow = append(toFollow, k)
				}
			}
		}
	}
	for _, j := range toFollow {
		Follow(follow, a, j+1)
	}
}

func FollowBackAll(follow [][]bool, a int, n int) {
	for j := 0; j < n; j++ {
		if follow[j][a-1] {
			Follow(follow, a, j+1)
		}
	}
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
	n, _, logs := parse()

	followMatrix := make([][]bool, n)
	for i := range followMatrix {
		followMatrix[i] = make([]bool, n)
	}

	for _, log := range logs {
		op := log[0]
		a := log[1]

		switch op {
		case 1:
			b := log[2]
			Follow(followMatrix, a, b)
		case 2:
			FollowBackAll(followMatrix, a, n)
		case 3:
			FollowFollow(followMatrix, a, n)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if followMatrix[i][j] {
				fmt.Print("Y")
			} else {
				fmt.Print("N")
			}
		}
		fmt.Println()
	}
}
