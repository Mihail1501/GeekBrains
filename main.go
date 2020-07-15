package main

import (
	"fmt"
)

func main() {

	const dollarRate float32 = 70.94
	var rub float32
	fmt.Println("Введите сумму в рублях")
	fmt.Scanln(&rub)

	var total = rub / dollarRate
	fmt.Println("Сумма в долларах:", total)

}
