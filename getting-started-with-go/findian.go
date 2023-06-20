package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var x string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input a string")
	scanner.Scan()
	x = scanner.Text()
	x = strings.ToLower(x)
	x = strings.ReplaceAll(x, " ", "")
	if strings.HasPrefix(x, "i") && strings.Contains(x, "a") && strings.HasSuffix(x, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
