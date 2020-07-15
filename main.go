package main

import (
	"fmt"
	"math"
)

func main() {

	var cat1 float64
	fmt.Println("Введите длину первого катета")
	fmt.Scanln(&cat1)

	var cat2 float64
	fmt.Println("Введите длину второго катета")
	fmt.Scanln(&cat2)

	fmt.Println("Площадь прямоугольного треугольника будет равна:", cat1*cat2/2)
	fmt.Println("Периметр прямоугольного треугольника будет равен:", (cat1+cat2)*2)

	var c = math.Pow(cat1, 2)
	fmt.Println("Квадрат первого катета будет равен:", c)

	var d = math.Pow(cat2, 2)
	fmt.Println("Квадрат второго катета будет равен:", d)

	sq := math.Sqrt(c + d)
	fmt.Println("Величина гепотенузы будет равна:", sq)
}

/*	const dollarRate float32 = 70.94
	var rub float32
	fmt.Println("Введите сумму в рублях")
	fmt.Scanln(&rub)

	var total = rub / dollarRate
	fmt.Println("Сумма в долларах:", total)*/
