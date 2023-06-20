package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	var x string
	var i int = -1
	for true {
		fmt.Println("Enter integer or 'X' to stop")
		fmt.Scan(&x)
		if x == "X" {
			break
		} else {
			if nr, err := strconv.Atoi(x); err == nil {
				if i == len(sli) {
					sli = append(sli, nr)
				} else {
					sli = append(sli[:i+1], nr)
					i++
				}
				sort.Ints(sli)
				fmt.Println(sli)
			}
		}
	}
}
