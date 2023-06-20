package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Please enter your address: ")
	scanner.Scan()
	address := scanner.Text()

	m := map[string]string{"name": name, "address": address}
	if json, err := json.Marshal(m); err == nil {
		fmt.Println(string(json))
	}
}
