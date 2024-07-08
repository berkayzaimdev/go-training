package main

import (
	"fmt"
	"math"
)

func adder() func(int) int { // closure fonksiyonu class gibi düşünebiliriz, her değişkene göre yenisini üretebiliriz. sum değeri de değişkene özgü olur
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 { // parametre olarak fonksiyon kabul eden fonksiyon
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 { // fonksiyonlar, main içerisinde bir type olarak tanımlanabilir.
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot)) // parametre olarak argüman kabul edecektir eğer parametresi buna uygunsa
	fmt.Println(compute(math.Pow))
	fmt.Printf("\n\n\n")


	pos, neg := adder(), adder() // adder closure'dan ayrı ayrı fonksiyonlar türettik
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
