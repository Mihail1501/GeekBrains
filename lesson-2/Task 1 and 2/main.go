package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("Введенное число четное")
	} else {
		fmt.Println("Введенное число нечетное")
	}

	fmt.Println("Теперь проверьте, делится ли число на 3.\nВведите число")
	fmt.Scanln(&number)

	if number%3 == 0 {
		fmt.Println("Введенное число делится на 3")
	} else {
		fmt.Println("Введенное число не делится на 3")
	}

}
