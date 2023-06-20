package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var file string
	fmt.Print("Input file name: ")
	fmt.Scan(&file)
	if f, err := os.Open(file); err == nil {
		scanner := bufio.NewScanner(f)
		res := []Name{}
		for scanner.Scan() {
			line := scanner.Text()
			name := strings.Split(line, " ")
			nameStruct := Name{name[0], name[1]}
			res = append(res, nameStruct)
		}
		for _, line := range res {
			fmt.Print(line.fname + " ")
			fmt.Println(line.lname)
		}
	} else {
		fmt.Println(err)
	}
}
