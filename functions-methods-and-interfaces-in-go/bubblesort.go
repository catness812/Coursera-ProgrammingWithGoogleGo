package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(sli []int, j int) {
	sli[j], sli[j+1] = sli[j+1], sli[j]
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a sequence of up to 10 integers separated by a single space (e.g. '1 2 3')")
	scanner.Scan()
	str_nrs := scanner.Text()
	splits := strings.Split(str_nrs, " ")
	nrs := []int{}

	for _, split := range splits {
		nr, err := strconv.Atoi(split)
		if err == nil {
			nrs = append(nrs, nr)
		}
	}

	if len(nrs) > 10 {
		fmt.Println("You entered more than 10 integers.")
	} else {
		BubbleSort(nrs)
		fmt.Println("Sorted sequence:", nrs)
	}
}
