package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (a *Cow) Eat() {
	fmt.Println("grass")
}
func (a *Cow) Move() {
	fmt.Println("walk")
}
func (a *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (a *Bird) Eat() {
	fmt.Println("worms")
}
func (a *Bird) Move() {
	fmt.Println("fly")
}
func (a *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (a *Snake) Eat() {
	fmt.Println("mice")
}
func (a *Snake) Move() {
	fmt.Println("slither")
}
func (a *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animal := make(map[string]Animal)
	for true {
		fmt.Print(">")
		scanner.Scan()
		prompt := scanner.Text()
		prompts := strings.Fields(prompt)

		switch prompts[0] {
		case "newanimal":
			switch prompts[2] {
			case "cow":
				animal[prompts[1]] = new(Cow)
			case "bird":
				animal[prompts[1]] = new(Bird)
			case "snake":
				animal[prompts[1]] = new(Snake)
			default:
				fmt.Println("Not a valid animal.")
			}
			fmt.Println("Created it!")
		case "query":
			a, found := animal[prompts[1]]
			if found {
				switch prompts[2] {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				default:
					fmt.Println("Not a valid action.")
				}
			} else {
				fmt.Println("Animal not found!")
			}
		default:
			fmt.Println("Not a valid prompt.")
		}
	}
}
