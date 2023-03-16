package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	YES = "Y"
	NO  = "N"
)

type SNS struct {
	follow_state []string
}

func (sns *SNS) Follow(a int, b int) {
	sns.follow_state[a-1] = sns.follow_state[a-1][:b-1] + YES + sns.follow_state[a-1][b:]
}

func (sns *SNS) FollowFollow(a int, n int) {
	toFollow := make([]int, 0)
	for j := 0; j < n; j++ {
		if sns.follow_state[a-1][j] == YES[0] {
			for k := 0; k < n; k++ {
				if sns.follow_state[j][k] == YES[0] && k != a-1 {
					toFollow = append(toFollow, k)
				}
			}
		}
	}
	for _, j := range toFollow {
		sns.Follow(a, j+1)
	}
}

func (sns *SNS) FollowBackAll(a int, n int) {
	for j := 0; j < n; j++ {
		if sns.follow_state[j][a-1] == YES[0] {
			sns.Follow(a, j+1)
		}
	}
}

func InitSNS(n int) *SNS {
	follow_state := make([]string, n)
	for i := range follow_state {
		follow_state[i] = strings.Repeat(NO, n)
	}
	return &SNS{follow_state: follow_state}
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

	sns := InitSNS(n)

	for _, log := range logs {
		op := log[0]
		a := log[1]

		switch op {
		case 1:
			b := log[2]
			sns.Follow(a, b)
		case 2:
			sns.FollowBackAll(a, n)
		case 3:
			sns.FollowFollow(a, n)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(sns.follow_state[i])
	}
}
