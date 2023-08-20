package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print(">")
		scanner.Scan()
		prompt := scanner.Text()
		prompts := strings.Fields(prompt)

		var animal Animal

		switch prompts[0] {
		case "cow":
			animal.food = "grass"
			animal.locomotion = "walk"
			animal.noise = "moo"
		case "bird":
			animal.food = "worms"
			animal.locomotion = "fly"
			animal.noise = "peep"
		case "snake":
			animal.food = "mice"
			animal.locomotion = "slither"
			animal.noise = "hsss"
		default:
			fmt.Println("Not a valid animal.")
		}

		switch prompts[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Not a valid action.")
		}
	}
}
