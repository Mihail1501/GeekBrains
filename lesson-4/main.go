package main

import "fmt"

type sort interface {
	Swap(i int)
}

type phones []int

func (p phones) Swap() {
	for i, phone := range p {
		fmt.Printf("\t %v) %v \n", i, phone)
	}
}

func main() {
	adressBook := make(map[string]phones)

	adressBook["Миша"] = phones{78293467382}
	adressBook["Никита"] = phones{89167253764, 89635437382}

	for name, ph := range adressBook {
		ph.Swap()
		fmt.Println(name)
	}
}
