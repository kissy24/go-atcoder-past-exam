package main

import (
	"fmt"
)

func inputSales() []int {
	var period int
	// fmt can be pre-typed
	fmt.Scan(&period)
	sales := make([]int, period)
	for i := 0; i < period; i++ {
		fmt.Scan(&sales[i])
	}
	return sales
}

func GenerateDiffs(sales []int) []int {
	var diffs []int
	for i := 1; i < len(sales); i++ {
		diff := sales[i] - sales[i-1]
		diffs = append(diffs, diff)
	}
	return diffs
}

func Format(salesDiff int) string {
	if salesDiff > 0 {
		return fmt.Sprintf("up %d", salesDiff)
	} else if salesDiff < 0 {
		// Sign reversal is easier to read than math.Abs.
		return fmt.Sprintf("down %d", -salesDiff)
	} else {
		return "stay"
	}
}

func main() {
	sales := inputSales()
	salesDiffs := GenerateDiffs(sales)
	for _, sd := range salesDiffs {
		fmt.Println(Format(sd))
	}
}
